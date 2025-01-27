package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Participant interface {
	Create(ctx context.Context, participant models.Participant) (*models.Participant, error)
	Select(ctx context.Context) ([]models.Participant, error)
	Get(ctx context.Context) (*models.Participant, error)

	Filter(filters map[string]any) Participant
	UpdateOne(ctx context.Context, fields map[string]any) error

	DeleteMany(ctx context.Context) error
	DeleteOne(ctx context.Context) error
}

type participant struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (p *participant) Create(ctx context.Context, participant models.Participant) (*models.Participant, error) {
	if p.filters == nil || len(p.filters) == 0 {
		return nil, fmt.Errorf("no filters set for participant creation")
	}

	participant.CreatedAt = time.Now().UTC()

	logrus.Infof("Creating participant: %v", participant)

	update := bson.M{
		"$push": bson.M{
			"participant": participant,
		},
		"$set": bson.M{
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := p.collection.UpdateOne(ctx, p.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to add participant to initiatives: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no initiatives found with the given filters")
	}
	return &participant, nil
}

func (p *participant) Select(ctx context.Context) ([]models.Participant, error) {
	iniID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var ini models.Initiative
	err := p.collection.FindOne(ctx, bson.M{"_id": iniID}).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("initiatives not found")
		}
		return nil, fmt.Errorf("failed to find initiatives: %w", err)
	}

	var results []models.Participant

	userIDVal, hasUserID := p.filters["participant.user_id"]
	firstNameVal, hasFirstName := p.filters["participant.first_name"]
	secondNameVal, hasSecondName := p.filters["participant.second_name"]

	for _, emp := range ini.Participants {
		if hasUserID {
			uid, okCast := userIDVal.(uuid.UUID)
			if !okCast || emp.UserID != uid {
				continue
			}
		}

		if hasFirstName {
			fn, okCast := firstNameVal.(string)
			if !okCast || emp.FirstName != fn {
				continue
			}
		}

		if hasSecondName {
			sn, okCast := secondNameVal.(string)
			if !okCast || emp.SecondName != sn {
				continue
			}
		}

		results = append(results, emp)
	}

	return results, nil
}

func (p *participant) Get(ctx context.Context) (*models.Participant, error) {
	orgID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	userID, ok := p.filters["participant.user_id"]
	if !ok {
		return nil, fmt.Errorf("participant user_id filter is missing (filters['participant.user_id'])")
	}

	filter := bson.M{
		"_id":                 orgID,
		"participant.user_id": userID,
	}
	projection := bson.M{
		"participant": bson.M{
			"$elemMatch": bson.M{"user_id": userID},
		},
	}

	opts := options.FindOne().SetProjection(projection)
	var ini models.Initiative
	err := p.collection.FindOne(ctx, filter, opts).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("participant not found")
		}
		return nil, fmt.Errorf("failed to find participant: %w", err)
	}

	if len(ini.Participants) == 0 {
		return nil, fmt.Errorf("participant not found")
	}

	return &ini.Participants[0], nil
}

func (p *participant) Filter(filters map[string]any) Participant {
	if p.filters == nil || p.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (initiatives ID)")
		return p
	}

	if userID, ok := filters["user_id"]; ok && userID != nil {
		p.filters["participant.user_id"] = userID
	}

	if firstName, ok := filters["first_name"]; ok && firstName != nil {
		p.filters["participant.first_name"] = firstName
	}

	if secondName, ok := filters["second_name"]; ok && secondName != nil {
		p.filters["participant.second_name"] = secondName
	}

	if displayName, ok := filters["display_name"]; ok && displayName != nil {
		p.filters["participant.display_name"] = displayName
	}

	if position, ok := filters["position"]; ok && position != nil {
		p.filters["participant.position"] = position
	}

	return p
}

func (p *participant) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields["participant.$[participant]."+key] = value
		}
	}

	updateFields["participant.$[participant].updated_at"] = time.Now().UTC()

	if len(updateFields) == 1 {
		return fmt.Errorf("no valid fields to update")
	}

	update := bson.M{"$set": updateFields}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participant.") {
			field := strings.TrimPrefix(key, "participant.")
			subFilters = append(subFilters, bson.M{"participant." + field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filter found (e.filters['participant.*'])")
	}

	arrayFilter := bson.M{"$and": subFilters}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	result, err := p.collection.UpdateOne(
		ctx,
		bson.M{"_id": orgID},
		update,
		arrayFilters,
	)
	if err != nil {
		return fmt.Errorf("failed to update participant: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no participant found with the given criteria")
	}

	return nil
}

func (p *participant) DeleteMany(ctx context.Context) error {
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participant.") {
			field := strings.TrimPrefix(key, "participant.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participant.*')")
	}

	participantsCondition := bson.M{"$and": subFilters}

	update := bson.M{
		"$pull": bson.M{
			"participant": participantsCondition,
		},
	}

	filter := bson.M{"_id": orgID}
	result, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete participant: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no participant found with the given criteria")
	}

	return nil
}

func (p *participant) DeleteOne(ctx context.Context) error {
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participant.") {
			field := strings.TrimPrefix(key, "participant.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}
	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participant.*')")
	}

	participantCondition := bson.M{"$and": subFilters}

	filter := bson.M{
		"_id":         orgID,
		"participant": bson.M{"$elemMatch": participantCondition},
	}
	projection := bson.M{
		"participant": bson.M{"$elemMatch": participantCondition},
	}
	findOpts := options.FindOne().SetProjection(projection)

	var ini models.Initiative
	err := p.collection.FindOne(ctx, filter, findOpts).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no matching participant found")
		}
		return fmt.Errorf("failed to find matching participant: %w", err)
	}

	if len(ini.Participants) == 0 {
		return fmt.Errorf("no matching participant found")
	}

	firstMatched := ini.Participants[0]

	if firstMatched.UserID == uuid.Nil {
		return fmt.Errorf("found participant but user_id is empty/invalid, cannot delete uniquely")
	}

	pullFilter := bson.M{"_id": orgID}
	pullUpdate := bson.M{
		"$pull": bson.M{
			"participant": bson.M{"user_id": firstMatched.UserID},
		},
	}

	res, err := p.collection.UpdateOne(ctx, pullFilter, pullUpdate)
	if err != nil {
		return fmt.Errorf("failed to delete the matched participant: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed to delete participant: it may no longer match or was deleted by someone else")
	}

	return nil
}
