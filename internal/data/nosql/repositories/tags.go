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

type Tags interface {
	Add(ctx context.Context, tag models.Tag) (*models.Tag, error)
	Select(ctx context.Context) ([]models.Tag, error)
	Get(ctx context.Context) (*models.Tag, error)

	Filter(filters map[string]any) Tags

	UpdateOne(ctx context.Context, fields map[string]any) error

	DeleteOne(ctx context.Context) error
}

type tags struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (t *tags) Add(ctx context.Context, tag models.Tag) (*models.Tag, error) {
	if t.filters == nil || len(t.filters) == 0 {
		return nil, fmt.Errorf("no filters set for tag creation")
	}

	logrus.Infof("Creating tag: %v", tag)

	update := bson.M{
		"$push": bson.M{
			"tags": tag,
		},
		"$set": bson.M{
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := t.collection.UpdateOne(ctx, t.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to create tag: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no tag found with the given filters")
	}

	return &tag, nil
}

func (t *tags) Select(ctx context.Context) ([]models.Tag, error) {
	iniID, ok := t.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var ini models.Initiative
	err := t.collection.FindOne(ctx, bson.M{"_id": iniID}).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("initiatives not found")
		}
		return nil, fmt.Errorf("failed to find initiatives: %w", err)
	}

	var results []models.Tag

	tagIDVal, hasTagID := t.filters["tags.tag_id"]
	TagNameVal, hasTagName := t.filters["tags.name"]
	TagWeightVal, hasTagWeight := t.filters["tags.weight"]

	for _, tag := range ini.Tags {
		if hasTagID {
			id, okCast := tagIDVal.(primitive.ObjectID)
			if !okCast || tag.ID != id {
				continue
			}
		}

		if hasTagName {
			fn, okCast := TagNameVal.(string)
			if !okCast || tag.Name != fn {
				continue
			}
		}

		if hasTagWeight {
			sn, okCast := TagWeightVal.(int)
			if !okCast || tag.Weight != sn {
				continue
			}
		}

		results = append(results, tag)
	}

	return results, nil
}

func (t *tags) Get(ctx context.Context) (*models.Tag, error) {
	iniID, ok := t.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	tagID, ok := t.filters["tags.tag_id"]
	if !ok {
		return nil, fmt.Errorf("participant user_id filter is missing (filters['tags.tag_id'])")
	}

	filter := bson.M{
		"_id":         iniID,
		"tags.tag_id": tagID,
	}
	projection := bson.M{
		"tags": bson.M{
			"$elemMatch": bson.M{"tag_id": tagID},
		},
	}

	opts := options.FindOne().SetProjection(projection)
	var ini models.Initiative
	err := t.collection.FindOne(ctx, filter, opts).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("tag not found")
		}
		return nil, fmt.Errorf("failed to find tag: %w", err)
	}

	if len(ini.Tags) == 0 {
		return nil, fmt.Errorf("tag not found")
	}

	return &ini.Tags[0], nil
}

func (t *tags) Filter(filters map[string]any) Tags {
	if t.filters == nil || t.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (initiatives ID)")
		return t
	}

	if tagID, ok := filters["tag_id"]; ok && tagID != nil {
		t.filters["tags.tag_id"] = tagID
	}

	if name, ok := filters["name"]; ok && name != nil {
		t.filters["tags.name"] = name
	}

	if weight, ok := filters["weight"]; ok && weight != nil {
		t.filters["tags.weight"] = weight
	}

	return t
}

func (t *tags) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields provided for update")
	}

	validFields := map[string]bool{
		"name":   true,
		"weight": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] && value != nil {
			updateFields["tags.$."+key] = value
		}
	}

	if len(updateFields) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	if t.filters == nil || t.filters["_id"] == nil {
		return fmt.Errorf("tag filters are empty or tag ID is not set")
	}

	filter := bson.M{
		"_id":         t.filters["_id"],
		"tags.tag_id": t.filters["tags."],
	}

	update := bson.M{"$set": updateFields}

	result, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update tag: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no tag found with the given criteria")
	}

	return nil
}

func (t *tags) DeleteOne(ctx context.Context) error {
	orgID, ok := t.filters["_id"]
	if !ok {
		return fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var subFilters []bson.M
	for key, val := range t.filters {
		if strings.HasPrefix(key, "tags.") {
			field := strings.TrimPrefix(key, "tags.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}
	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participant.*')")
	}

	participantCondition := bson.M{"$and": subFilters}

	filter := bson.M{
		"_id":  orgID,
		"tags": bson.M{"$elemMatch": participantCondition},
	}
	projection := bson.M{
		"tags": bson.M{"$elemMatch": participantCondition},
	}
	findOpts := options.FindOne().SetProjection(projection)

	var org models.Initiative
	err := t.collection.FindOne(ctx, filter, findOpts).Decode(&org)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no matching tag found")
		}
		return fmt.Errorf("failed to find matching tag: %w", err)
	}

	if len(org.Tags) == 0 {
		return fmt.Errorf("no matching tag found")
	}

	firstMatched := org.Tags[0]

	pullFilter := bson.M{"_id": firstMatched.ID}
	pullUpdate := bson.M{
		"$pull": bson.M{
			"tags": bson.M{"tag_id": firstMatched.ID},
		},
	}

	res, err := t.collection.UpdateOne(ctx, pullFilter, pullUpdate)
	if err != nil {
		return fmt.Errorf("failed to delete the matched tag: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed to delete tag: it may no longer match or was deleted by someone else")
	}

	return nil
}
