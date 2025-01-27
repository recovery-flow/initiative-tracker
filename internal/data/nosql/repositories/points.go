package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Points interface {
	New() Points

	Insert(ctx context.Context, org models.Point) (*models.Point, error)
	DeleteOne(ctx context.Context) error
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Point, error)
	Get(ctx context.Context) (*models.Point, error)

	Filter(filters map[string]any) Points

	Jar() Jar
	Docs() Docs

	Update(ctx context.Context, fields map[string]any) (*models.Point, error)

	SortBy(field string, ascending bool) Points
	Limit(limit int64) Points
	Skip(skip int64) Points
}

type points struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewPoint(uri, dbName, collectionName string) (Points, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &points{
		client:     client,
		database:   database,
		collection: collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (p *points) New() Points {
	return &points{
		client:     p.client,
		database:   p.database,
		collection: p.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (p *points) Insert(ctx context.Context, point models.Point) (*models.Point, error) {
	point.ID = primitive.NewObjectID()

	_, err := p.collection.InsertOne(ctx, point)
	if err != nil {
		return nil, fmt.Errorf("failed to insert points: %w", err)
	}

	return &point, nil
}

func (p *points) DeleteOne(ctx context.Context) error {
	_, err := p.collection.DeleteOne(ctx, p.filters)
	if err != nil {
		return fmt.Errorf("failed to delete points: %w", err)
	}

	return nil
}

func (p *points) Count(ctx context.Context) (int64, error) {
	count, err := p.collection.CountDocuments(ctx, p.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count points: %w", err)
	}
	return count, nil
}

func (p *points) Select(ctx context.Context) ([]models.Point, error) {
	cursor, err := p.collection.Find(ctx, p.filters, options.Find().SetSort(p.sort).SetLimit(p.limit).SetSkip(p.skip))
	if err != nil {
		return nil, fmt.Errorf("failed to select points: %w", err)
	}

	var pnts []models.Point
	if err = cursor.All(ctx, &pnts); err != nil {
		return nil, fmt.Errorf("failed to decode points: %w", err)
	}

	return pnts, nil
}

func (p *points) Get(ctx context.Context) (*models.Point, error) {
	var point models.Point
	err := p.collection.FindOne(ctx, p.filters).Decode(&point)
	if err != nil {
		return nil, fmt.Errorf("failed to find points: %w", err)
	}

	return &point, nil
}

func (p *points) Filter(filters map[string]any) Points {
	var validFilters = map[string]bool{
		"_id":           true,
		"initiative_id": true,
		"parent_id":     true,
		"title":         true,
		"desc":          true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		p.filters[field] = value
	}
	return p
}

func (p *points) Jar() Jar {
	return &jar{
		client:     p.client,
		database:   p.database,
		collection: p.collection,
		filters:    p.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (p *points) Docs() Docs {
	return &docs{
		client:     p.client,
		database:   p.database,
		collection: p.collection,
		filters:    p.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (p *points) Update(ctx context.Context, fields map[string]any) (*models.Point, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	_, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	validFields := map[string]bool{
		"parent_id":            true,
		"title":                true,
		"desc":                 true,
		"local_cost":           true,
		"local_collected":      true,
		"cumulative_cost":      true,
		"cumulative_collected": true,
		"bank_info":            true,
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
		return nil, fmt.Errorf("failed to update initiatives: %w", err)
	}

	return p.Get(ctx)
}

func (p *points) SortBy(field string, ascending bool) Points {
	order := 1
	if !ascending {
		order = -1
	}
	p.sort = append(p.sort, bson.E{Key: field, Value: order})
	return p
}

func (p *points) Skip(skip int64) Points {
	p.skip = skip
	return p
}

func (p *points) Limit(limit int64) Points {
	p.limit = limit
	return p
}
