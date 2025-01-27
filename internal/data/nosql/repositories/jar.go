package repositories

import (
	"context"
	"fmt"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Jar interface {
	Update(ctx context.Context, jar models.Jar) (*models.Jar, error)
}

type jar struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (j *jar) Update(ctx context.Context, jarData models.Jar) (*models.Jar, error) {
	// 1. Проверяем, что есть фильтры – иначе не знаем, какую запись обновлять
	if len(j.filters) == 0 {
		return nil, fmt.Errorf("no filters provided")
	}

	// 2. Формируем $set для обновления поля "bank_info"
	update := bson.M{
		"$set": bson.M{
			"bank_info": jarData,
		},
	}

	// 3. Выполняем UpdateOne (по j.filters)
	result, err := j.collection.UpdateOne(ctx, j.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update jar: %w", err)
	}

	// Если ModifiedCount == 0, это может означать, что документ не найден
	// (или новые данные те же самые, и Mongo посчитала, что обновлять нечего).
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no document matches the given filters")
	}
	// Optional: если нужно гарантированно узнать, действительно ли что-то изменилось
	// смотрите на result.ModifiedCount. Если == 0, то документ найден, но не изменён.

	// 4. Считываем обратно обновлённый документ (чтобы вернуть актуальные данные jar)
	var updatedPoint models.Point
	if err := j.collection.FindOne(ctx, j.filters).Decode(&updatedPoint); err != nil {
		return nil, fmt.Errorf("failed to fetch updated jar: %w", err)
	}

	// 5. Возвращаем обновлённую структуру Jar (поле bank_info)
	//    Если в updatedPoint нет поля Jar (nil), значит где-то не совпали названия/типы.
	return updatedPoint.Jar, nil
}
