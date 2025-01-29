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

type Participants interface {
	Create(ctx context.Context, participant models.Participant) (*models.Participant, error)
	Select(ctx context.Context) ([]models.Participant, error)
	Get(ctx context.Context) (*models.Participant, error)

	Filter(filters map[string]any) Participants
	UpdateOne(ctx context.Context, fields map[string]any) error

	DeleteMany(ctx context.Context) error
	DeleteOne(ctx context.Context) error
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

func (p *participants) Create(ctx context.Context, participant models.Participant) (*models.Participant, error) {
	if p.filters == nil || len(p.filters) == 0 {
		return nil, fmt.Errorf("no filters set for participants creation")
	}

	participant.CreatedAt = time.Now().UTC()

	logrus.Infof("Creating participants: %v", participant)

	update := bson.M{
		"$push": bson.M{
			"participants": participant,
		},
		"$set": bson.M{
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := p.collection.UpdateOne(ctx, p.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to add participants to initiatives: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no initiatives found with the given filters")
	}
	return &participant, nil
}

func (p *participants) Select(ctx context.Context) ([]models.Participant, error) {
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

	userIDVal, hasUserID := p.filters["participants.user_id"]
	firstNameVal, hasFirstName := p.filters["participants.first_name"]
	secondNameVal, hasSecondName := p.filters["participants.second_name"]

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

func (p *participants) Get(ctx context.Context) (*models.Participant, error) {
	iniID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	userID, ok := p.filters["participants.user_id"]
	if !ok {
		return nil, fmt.Errorf("participants user_id filter is missing (filters['participants.user_id'])")
	}

	filter := bson.M{
		"_id":                  iniID,
		"participants.user_id": userID,
	}
	projection := bson.M{
		"participants": bson.M{
			"$elemMatch": bson.M{"user_id": userID},
		},
	}

	opts := options.FindOne().SetProjection(projection)
	var ini models.Initiative
	err := p.collection.FindOne(ctx, filter, opts).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("participants not found")
		}
		return nil, fmt.Errorf("failed to find participants: %w", err)
	}

	if len(ini.Participants) == 0 {
		return nil, fmt.Errorf("participants not found")
	}

	return &ini.Participants[0], nil
}

func (p *participants) Filter(filters map[string]any) Participants {
	if p.filters == nil || p.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (initiatives ID)")
		return p
	}

	if userID, ok := filters["user_id"]; ok && userID != nil {
		p.filters["participants.user_id"] = userID
	}

	if firstName, ok := filters["first_name"]; ok && firstName != nil {
		p.filters["participants.first_name"] = firstName
	}

	if secondName, ok := filters["second_name"]; ok && secondName != nil {
		p.filters["participants.second_name"] = secondName
	}

	if displayName, ok := filters["display_name"]; ok && displayName != nil {
		p.filters["participants.display_name"] = displayName
	}

	if position, ok := filters["position"]; ok && position != nil {
		p.filters["participants.position"] = position
	}

	logrus.Infof("Filters: %v", p.filters)

	return p
}

func (p *participants) UpdateOne(ctx context.Context, fields map[string]any) error {
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
			updateFields["participants.$[participants]."+key] = value
		}
	}

	updateFields["participants.$[participants].updated_at"] = time.Now().UTC()

	if len(updateFields) == 1 {
		return fmt.Errorf("no valid fields to update")
	}

	update := bson.M{"$set": updateFields}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participants.") {
			field := strings.TrimPrefix(key, "participants.")
			subFilters = append(subFilters, bson.M{"participants." + field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no participants filter found (e.filters['participants.*'])")
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
		return fmt.Errorf("failed to update participants: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no participants found with the given criteria")
	}

	return nil
}

func (p *participants) DeleteMany(ctx context.Context) error {
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participants.") {
			field := strings.TrimPrefix(key, "participants.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no participants filters found (keys like 'participants.*')")
	}

	participantsCondition := bson.M{"$and": subFilters}

	update := bson.M{
		"$pull": bson.M{
			"participants": participantsCondition,
		},
	}

	filter := bson.M{"_id": orgID}
	result, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete participants: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no participants found with the given criteria")
	}

	return nil
}

func (p *participants) DeleteOne(ctx context.Context) error {
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("initiatives ID filter is missing (filters['_id'])")
	}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participants.") {
			field := strings.TrimPrefix(key, "participants.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}
	if len(subFilters) == 0 {
		return fmt.Errorf("no participants filters found (keys like 'participants.*')")
	}

	participantCondition := bson.M{"$and": subFilters}

	filter := bson.M{
		"_id":          orgID,
		"participants": bson.M{"$elemMatch": participantCondition},
	}
	projection := bson.M{
		"participants": bson.M{"$elemMatch": participantCondition},
	}
	findOpts := options.FindOne().SetProjection(projection)

	var ini models.Initiative
	err := p.collection.FindOne(ctx, filter, findOpts).Decode(&ini)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no matching participants found")
		}
		return fmt.Errorf("failed to find matching participants: %w", err)
	}

	if len(ini.Participants) == 0 {
		return fmt.Errorf("no matching participants found")
	}

	firstMatched := ini.Participants[0]

	if firstMatched.UserID == uuid.Nil {
		return fmt.Errorf("found participants but user_id is empty/invalid, cannot delete uniquely")
	}

	pullFilter := bson.M{"_id": orgID}
	pullUpdate := bson.M{
		"$pull": bson.M{
			"participants": bson.M{"user_id": firstMatched.UserID},
		},
	}

	res, err := p.collection.UpdateOne(ctx, pullFilter, pullUpdate)
	if err != nil {
		return fmt.Errorf("failed to delete the matched participants: %w", err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed to delete participants: it may no longer match or was deleted by someone else")
	}

	return nil
}
