package repositories

import (
	"context"
	"fmt"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (i *initiatives) FilterOrgs(filters map[string]any) Initiatives {
	/*
	   Ожидаем, что в filters могут быть ключи:
	     - "organization_id" (bson/primitive.ObjectID)
	     - "status" (StatusOrg)
	     - "since" (primitive.DateTime или time.Time)
	   Если они есть, формируем subFilter. Затем i.filters["organizations"] = {"$elemMatch": subFilter}.
	*/
	subFilter := bson.M{}

	if orgID, ok := filters["organization_id"]; ok && orgID != nil {
		// orgID должен быть primitive.ObjectID, иначе игнорируем или логируем
		subFilter["organization_id"] = orgID
	}

	if statusVal, ok := filters["status"]; ok && statusVal != nil {
		// statusVal должно быть StatusOrg (строка)
		subFilter["status"] = statusVal
	}

	if sinceVal, ok := filters["since"]; ok && sinceVal != nil {
		// sinceVal может быть primitive.DateTime или time.Time
		// Если это строка, можно попытаться распарсить...
		// Для простоты предполагаем, что это уже подходящий тип
		subFilter["since"] = sinceVal
	}

	// Если subFilter не пустой, делаем elemMatch
	if len(subFilter) > 0 {
		i.filters["organizations"] = bson.M{"$elemMatch": subFilter}
	}

	return i
}

func (i *initiatives) UpdateOrgMember(
	ctx context.Context,
	orgID primitive.ObjectID,
	fields map[string]any,
) (*models.Initiative, error) {

	// Проверяем, что есть фильтр для самой инициативы (часто _id)
	if i.filters == nil || i.filters["_id"] == nil {
		return nil, fmt.Errorf("initiative filter is missing (need at least '_id')")
	}

	// Разрешённые поля для обновления внутри OrgMember
	validFields := map[string]bool{
		"status": true,
		"since":  true,
		// Возможно, будут другие поля, если расширите OrgMember
	}

	// Собираем updateFields
	updateFields := bson.M{}
	for key, val := range fields {
		if !validFields[key] {
			continue
		}
		// Мы меняем organizations.$[org].<key>
		updateFields["organizations.$[org]."+key] = val
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no valid fields to update in org member")
	}

	update := bson.M{
		"$set": updateFields,
	}

	// Готовим arrayFilters: обновим только элемент, где organization_id == orgID
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"org.organization_id": orgID},
		},
	})

	// Основной фильтр по инициативе (например, {"_id": <someInitiativeID>})
	filter := bson.M{}
	for k, v := range i.filters {
		filter[k] = v
	}

	// Выполняем UpdateOne
	res, err := i.collection.UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		return nil, fmt.Errorf("failed to update org member: %w", err)
	}

	if res.ModifiedCount == 0 {
		// Значит не нашёлся документ, где "organizations.organization_id" == orgID
		// или поля не изменились
		return nil, fmt.Errorf("no matching org member found or nothing changed")
	}

	// Возвращаем обновлённый документ (Инициативу)
	var updated models.Initiative
	// filter всё тот же, так как _id не меняется
	err = i.collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated initiative: %w", err)
	}

	return &updated, nil
}

func (i *initiatives) AddOrgMember(ctx context.Context, member models.OrgMember) (*models.Initiative, error) {
	// Проверяем, что есть обязательный фильтр (например, _id инициативы)
	if i.filters == nil || i.filters["_id"] == nil {
		return nil, fmt.Errorf("initiative filter is missing (need at least '_id')")
	}

	// Формируем запрос на добавление
	update := bson.M{
		"$push": bson.M{
			"organizations": member,
		},
	}

	// Выполняем UpdateOne
	res, err := i.collection.UpdateOne(ctx, i.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to add org member: %w", err)
	}
	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("no initiative found with the given criteria")
	}

	// Возвращаем обновлённую инициативу
	var updated models.Initiative
	if err := i.collection.FindOne(ctx, i.filters).Decode(&updated); err != nil {
		return nil, fmt.Errorf("failed to fetch updated initiative: %w", err)
	}

	return &updated, nil
}
