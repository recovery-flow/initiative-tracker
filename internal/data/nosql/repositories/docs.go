package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Docs interface {
	Add(ctx context.Context, doc models.Docs) (*models.Docs, error)
	Select(ctx context.Context) ([]models.Docs, error)
	Get(ctx context.Context) (*models.Docs, error)

	Filter(filters map[string]interface{}) Docs

	UpdateOne(ctx context.Context, fields map[string]any) (*models.Docs, error)

	DeleteOne(ctx context.Context) error
}

type docs struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (d *docs) Add(ctx context.Context, doc models.Docs) (*models.Docs, error) {
	if d.filters == nil || len(d.filters) == 0 {
		return nil, fmt.Errorf("no filters set for doc creation")
	}

	doc.ID = primitive.NewObjectID()
	doc.CreatedAt = time.Now().UTC()

	logrus.Infof("Creating doc: %v", doc)

	update := bson.M{
		"$push": bson.M{
			"doc": doc,
		},
		"$set": bson.M{
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := d.collection.UpdateOne(ctx, d.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to create doc: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no initiatives found with the given filters")
	}
	return &doc, nil
}

func (d *docs) Select(ctx context.Context) ([]models.Docs, error) {
	iniID, ok := d.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("point ID filter is missing (filters['_id'])")
	}

	var point models.Point
	err := d.collection.FindOne(ctx, bson.M{"_id": iniID}).Decode(&point)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("point not found")
		}
		return nil, fmt.Errorf("failed to find point: %w", err)
	}

	var results []models.Docs

	docsNameVal, docsNameHas := d.filters["docs.name"]
	docsDescVal, docsDescHas := d.filters["docs.desc"]
	docsLinkVal, docsLinkHas := d.filters["docs.link"]

	for _, doc := range point.Docs {
		if docsNameHas && doc.Name != docsNameVal {
			continue
		}
		if docsDescHas && doc.Desc != docsDescVal {
			continue
		}
		if docsLinkHas && doc.Link != docsLinkVal {
			continue
		}
		results = append(results, doc)
	}

	return results, nil
}

func (d *docs) Get(ctx context.Context) (*models.Docs, error) {
	pointID, ok := d.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("point pointID filter is missing (filters['_id'])")
	}

	filter := bson.M{
		"_id":     pointID,
		"docs.id": d.filters["docs.id"],
	}
	projection := bson.M{
		"docs": bson.M{
			"$elemMatch": bson.M{"id": d.filters["docs.id"]},
		},
	}

	opts := options.FindOne().SetProjection(projection)
	var point models.Point
	err := d.collection.FindOne(ctx, filter, opts).Decode(&point)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("point not found")
		}
		return nil, fmt.Errorf("failed to find point: %w", err)
	}

	if len(point.Docs) == 0 {
		return nil, fmt.Errorf("docs not found")
	}

	return &point.Docs[0], nil
}

func (d *docs) Filter(filters map[string]any) Docs {
	if d.filters == nil || d.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (point ID)")
		return d
	}

	if id, ok := filters["id"]; ok && id != nil {
		d.filters["docs.id"] = id
	}

	if name, ok := filters["name"]; ok && name != nil {
		d.filters["docs.name"] = name
	}
	return d
}

func (d *docs) UpdateOne(ctx context.Context, fields map[string]any) (*models.Docs, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	pointID, ok := d.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("point ID filter is missing (filters['_id'])")
	}

	validFields := map[string]bool{
		"name": true,
		"desc": true,
		"link": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] && value != nil {
			updateFields["docs.$."+key] = value
		}
	}

	updateFields["docs.$[docs].updated_at"] = time.Now().UTC()

	update := bson.M{"$set": updateFields}

	var subFilters []bson.M
	for key, val := range d.filters {
		if strings.HasPrefix(key, "docs.") {
			field := strings.TrimPrefix(key, "docs.")
			subFilters = append(subFilters, bson.M{"docs." + field: val})
		}
	}

	if len(subFilters) == 0 {
		return nil, fmt.Errorf("no docs filter found (e.filters['docs.*'])")
	}

	arrayFilter := bson.M{"$and": subFilters}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	result, err := d.collection.UpdateOne(
		ctx,
		bson.M{"_id": pointID},
		update,
		arrayFilters,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update docs: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no docs found with the given criteria")
	}

	return d.Get(ctx)
}

func (d *docs) DeleteOne(ctx context.Context) error {
	pointID, ok := d.filters["_id"]
	if !ok {
		return fmt.Errorf("point ID filter is missing (filters['_id'])")
	}

	var subFilters []bson.M
	for key, val := range d.filters {
		if strings.HasPrefix(key, "docs.") {
			field := strings.TrimPrefix(key, "docs.")
			subFilters = append(subFilters, bson.M{"docs." + field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no docs filter found (e.filters['docs.*'])")
	}

	condition := bson.M{"$and": subFilters}

	filter := bson.M{
		"_id":  pointID,
		"docs": bson.M{"$elemMatch": condition},
	}
	projection := bson.M{
		"docs": bson.M{"$elemMatch": condition},
	}
	findOpts := options.FindOne().SetProjection(projection)

	var point models.Point
	err := d.collection.FindOne(ctx, filter, findOpts).Decode(&point)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no matching docs found")
		}
		return fmt.Errorf("failed to find matching docs: %w", err)
	}

	if len(point.Docs) == 0 {
		return fmt.Errorf("no matching docs found")
	}

	firstMatched := point.Docs[0]

	pullFilter := bson.M{"_id": pointID}
	pullUpdate := bson.M{
		"$pull": bson.M{
			"docs": bson.M{"id": firstMatched.ID},
		},
	}

	res, err := d.collection.UpdateOne(ctx, pullFilter, pullUpdate)
	if err != nil {
		return fmt.Errorf("failed to delete the matched docs: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed to delete docs: it may no longer match or was deleted by someone else")
	}

	return nil
}
