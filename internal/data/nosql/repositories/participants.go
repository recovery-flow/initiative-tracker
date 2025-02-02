package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Participants interface {
	New() Participants

	Insert(ctx context.Context, participant models.Participant) (*models.Participant, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Participant, error)
	Get(ctx context.Context) (*models.Participant, error)

	FilterExact(filters map[string]any) Participants
	FilterSoft(filters map[string]any) Participants
	FilterDateBounds(dateFilters map[string]any, after bool) Participants

	UpdateOne(ctx context.Context, fields map[string]any) (*models.Participant, error)

	DeleteMany(ctx context.Context) error
	DeleteOne(ctx context.Context) error

	//TODO add Tag methods

	SortBy(field string, ascending bool) Participants
	Limit(limit int64) Participants
	Skip(skip int64) Participants
}

type participants struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewParticipant(uri, dbName, collectionName string) (Participants, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &participants{
		client:     client,
		database:   database,
		collection: collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (p *participants) New() Participants {
	return &participants{
		client:     p.client,
		database:   p.database,
		collection: p.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (p *participants) Insert(ctx context.Context, participant models.Participant) (*models.Participant, error) {
	participant.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	participant.ID = primitive.NewObjectID()

	res, err := p.collection.InsertOne(ctx, participant)
	if err != nil {
		return nil, fmt.Errorf("failed to add participants to initiatives: %w", err)
	}
	participant.ID = res.InsertedID.(primitive.ObjectID)

	return &participant, nil
}

func (p *participants) Select(ctx context.Context) ([]models.Participant, error) {
	findOptions := options.Find().SetSort(p.sort).SetLimit(p.limit).SetSkip(p.skip)
	cursor, err := p.collection.Find(ctx, p.filters, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to select participant: %w", err)
	}

	var participant []models.Participant
	if err = cursor.All(ctx, &participant); err != nil {
		return nil, fmt.Errorf("failed to decode initiatives: %w", err)
	}

	return participant, nil
}

func (p *participants) Get(ctx context.Context) (*models.Participant, error) {
	var participant models.Participant
	err := p.collection.FindOne(ctx, p.filters).Decode(&participant)
	if err != nil {
		return nil, fmt.Errorf("failed to find participant: %w", err)
	}

	return &participant, nil
}

func (p *participants) Count(ctx context.Context) (int64, error) {
	count, err := p.collection.CountDocuments(ctx, p.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count participants: %w", err)
	}
	return count, nil
}

func (p *participants) FilterExact(filters map[string]any) Participants {
	// Поля, разрешённые для точной фильтрации
	var validFiltersExact = map[string]bool{
		"_id":           true,
		"user_id":       true,
		"initiative_id": true,
		"first_name":    true,
		"second_name":   true,
		"third_name":    true,
		"display_name":  true,
		"verified":      true,
		"role":          true,
		"position":      true,
		"created_at":    true,
		"updated_at":    true,
	}

	for field, value := range filters {
		if !validFiltersExact[field] {
			continue
		}
		if value == nil {
			continue
		}
		// Точное соответствие
		p.filters[field] = value
	}
	return p
}

func (p *participants) FilterSoft(filters map[string]any) Participants {
	// Поля, которые допускаем для подстрочного поиска
	softFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"third_name":   true,
		"display_name": true,
		"position":     true,
	}

	for field, value := range filters {
		if !softFields[field] {
			// Если поле не разрешено для мягкого поиска, пропустим
			continue
		}
		if value == nil {
			continue
		}

		strVal, ok := value.(string)
		if !ok {
			logrus.Warnf("FilterSoft: field %s is not a string, got %T", field, value)
			continue
		}
		// Делаем фильтр вида {"field": {"$regex": strVal, "$options": "i"}}
		p.filters[field] = bson.M{
			"$regex":   strVal,
			"$options": "i",
		}
	}

	return p
}

func (p *participants) FilterDateBounds(dateFilters map[string]any, after bool) Participants {
	// Разрешаем фильтровать только по этим двум полям
	validDateFields := map[string]bool{
		"created_at": true,
		"updated_at": true,
	}

	// Для удобства определим оператор
	var op string
	if after {
		// Ищем записи "после или в" этой даты
		op = "$gte"
	} else {
		// Ищем записи "до или в" этой даты
		op = "$lte"
	}

	for field, value := range dateFilters {
		if !validDateFields[field] {
			// Игнорируем любые поля, кроме "created_at" и "updated_at"
			continue
		}
		if value == nil {
			continue
		}

		// Попробуем интерпретировать value как time.Time или строку
		var t time.Time
		switch val := value.(type) {
		case time.Time:
			t = val
		case *time.Time:
			t = *val
		case string:
			// Попробуем распарсить строку как RFC3339 (или любой другой формат)
			parsed, err := time.Parse(time.RFC3339, val)
			if err != nil {
				// Если парсинг не удался, можно проигнорировать или залогировать
				logrus.Warnf("FilterDateBounds: cannot parse date '%s': %v", val, err)
				continue
			}
			t = parsed
		default:
			// Если тип неподходящий – просто пропустим
			logrus.Warnf("FilterDateBounds: field %s is not a recognized date type: %T", field, value)
			continue
		}

		// Формируем условие вида: {"created_at": {"$gte": t}} или {"updated_at": {"$lte": t}}
		p.filters[field] = bson.M{op: t}
	}

	return p
}

func (p *participants) UpdateOne(ctx context.Context, fields map[string]any) (*models.Participant, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	_, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"third_name":   true,
		"display_name": true,
		"desc":         true,
		"verified":     true,
		"role":         true,
		"position":     true,
	}

	updateFields := bson.M{}
	for field, value := range fields {
		if !validFields[field] {
			continue
		}
		if value == nil {
			continue
		}
		updateFields[field] = value
	}

	updateFields["updated_at"] = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	_, err := p.collection.UpdateOne(ctx, p.filters, bson.M{"$set": updateFields})
	if err != nil {
		return nil, fmt.Errorf("failed to update participant: %w", err)
	}

	return p.Get(ctx)
}

func (p *participants) DeleteMany(ctx context.Context) error {
	_, err := p.collection.DeleteMany(ctx, p.filters)
	if err != nil {
		return fmt.Errorf("failed to delete participants: %w", err)
	}
	return nil
}

func (p *participants) DeleteOne(ctx context.Context) error {
	_, err := p.collection.DeleteOne(ctx, p.filters)
	if err != nil {
		return fmt.Errorf("failed to delete participant: %w", err)
	}
	return nil
}

func (p *participants) SortBy(field string, ascending bool) Participants {
	order := 1
	if !ascending {
		order = -1
	}
	p.sort = append(p.sort, bson.E{Key: field, Value: order})
	return p
}

func (p *participants) Skip(skip int64) Participants {
	p.skip = skip
	return p
}

func (p *participants) Limit(limit int64) Participants {
	p.limit = limit
	return p
}
