package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Initiatives interface {
	New() Initiatives

	Insert(ctx context.Context, org models.Initiative) (*models.Initiative, error)
	DeleteOne(ctx context.Context) error
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Initiative, error)
	Get(ctx context.Context) (*models.Initiative, error)

	Filter(filters map[string]any) Initiatives

	Tags() Tags

	UpdateOne(ctx context.Context, fields map[string]any) (*models.Initiative, error)

	SortBy(field string, ascending bool) Initiatives
	Limit(limit int64) Initiatives
	Skip(skip int64) Initiatives
}

type initiatives struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewInitiative(uri, dbName, collectionName string) (Initiatives, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &initiatives{
		client:     client,
		database:   database,
		collection: collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (i *initiatives) New() Initiatives {
	return &initiatives{
		client:     i.client,
		database:   i.database,
		collection: i.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (i *initiatives) Insert(ctx context.Context, initiative models.Initiative) (*models.Initiative, error) {
	initiative.ID = primitive.NewObjectID()
	initiative.Name = strings.TrimSpace(initiative.Name)
	if initiative.Name == "" {
		return nil, fmt.Errorf("initiatives name cannot be empty")
	}
	initiative.Desc = strings.TrimSpace(initiative.Desc)
	if initiative.Desc == "" {
		return nil, fmt.Errorf("initiatives description cannot be empty")
	}
	initiative.Goal = strings.TrimSpace(initiative.Goal)
	if initiative.Location != nil {
		*initiative.Location = strings.TrimSpace(*initiative.Location)
		if *initiative.Location == "" {
			return nil, fmt.Errorf("incorect location")
		}
	}
	iniStatus := models.StatusFromString(string(initiative.Status))
	if iniStatus == nil {
		return nil, fmt.Errorf("incorect status")
	}
	initiative.Status = *iniStatus
	initiative.ChatID = primitive.NewObjectID()
	initiative.CreatedAt = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	res, err := i.collection.InsertOne(ctx, initiative)
	if err != nil {
		return nil, fmt.Errorf("failed to insert initiatives: %w", err)
	}
	initiative.ID = res.InsertedID.(primitive.ObjectID)

	return &initiative, nil
}

func (i *initiatives) DeleteOne(ctx context.Context) error {
	_, err := i.collection.DeleteOne(ctx, i.filters)
	if err != nil {
		return fmt.Errorf("failed to delete initiatives: %w", err)
	}
	return nil
}

func (i *initiatives) Count(ctx context.Context) (int64, error) {
	count, err := i.collection.CountDocuments(ctx, i.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count initiatives: %w", err)
	}
	return count, nil
}

func (i *initiatives) Select(ctx context.Context) ([]models.Initiative, error) {
	findOptions := options.Find().SetSort(i.sort).SetLimit(i.limit).SetSkip(i.skip)
	cursor, err := i.collection.Find(ctx, i.filters, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to select initiatives: %w", err)
	}

	var initiatives []models.Initiative
	if err = cursor.All(ctx, &initiatives); err != nil {
		return nil, fmt.Errorf("failed to decode initiatives: %w", err)
	}

	return initiatives, nil
}

func (i *initiatives) Get(ctx context.Context) (*models.Initiative, error) {
	var ini models.Initiative
	err := i.collection.FindOne(ctx, i.filters).Decode(&ini)
	if err != nil {
		return nil, fmt.Errorf("failed to find initiatives: %w", err)
	}

	return &ini, nil
}

func (i *initiatives) Filter(filters map[string]any) Initiatives {
	var validFilters = map[string]bool{
		"_id":      true,
		"name":     true,
		"verified": true,
		"location": true,
		"status":   true,
		"tags":     true,
		"likes":    true,
		"reposts":  true,
		"reports":  true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		i.filters[field] = value
	}
	return i
}

func (i *initiatives) Tags() Tags {
	return &tags{
		client:     i.client,
		database:   i.database,
		collection: i.collection,
		filters:    i.filters,
		sort:       i.sort,
		limit:      i.limit,
		skip:       i.skip,
	}
}

func (i *initiatives) UpdateOne(ctx context.Context, fields map[string]any) (*models.Initiative, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	_, ok := i.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	validFields := map[string]bool{
		"name":         true,
		"desc":         true,
		"goal":         true,
		"verified":     true,
		"location":     true,
		"status":       true,
		"tags":         true,
		"projects":     true,
		"participants": true,
		"likes":        true,
		"reposts":      true,
		"reports":      true,
		"closed_at":    true,
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

	_, err := i.collection.UpdateOne(ctx, i.filters, bson.M{"$set": updateFields})
	if err != nil {
		return nil, fmt.Errorf("failed to update initiatives: %w", err)
	}

	return i.Get(ctx)
}

func (i *initiatives) SortBy(field string, ascending bool) Initiatives {
	order := 1
	if !ascending {
		order = -1
	}
	i.sort = append(i.sort, bson.E{Key: field, Value: order})
	return i
}

func (i *initiatives) Skip(skip int64) Initiatives {
	i.skip = skip
	return i
}

func (i *initiatives) Limit(limit int64) Initiatives {
	i.limit = limit
	return i
}
