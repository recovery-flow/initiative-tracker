package repositories

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *initiatives) FilterExact(filters map[string]any) Initiatives {
	var validFilters = map[string]bool{
		"_id":            true,
		"name":           true,
		"desc":           true,
		"goal":           true,
		"verified":       true,
		"location":       true,
		"type":           true, // IniType (enum)
		"status":         true, // IniStatus (enum)
		"chat_id":        true,
		"final_cost":     true,
		"collected_cost": true,
		"likes":          true,
		"reposts":        true,
		"reports":        true,
		// Даты — обрабатываем в FilterDateBounds, поэтому здесь не указываем
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		// Строгое равенство, например {field: value}
		i.filters[field] = value
	}
	return i
}

func (i *initiatives) FilterSoft(filters map[string]any) Initiatives {
	softFields := map[string]bool{
		"name":     true,
		"desc":     true,
		"goal":     true,
		"location": true,
	}

	for field, value := range filters {
		if !softFields[field] {
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
		i.filters[field] = bson.M{
			"$regex":   strVal,
			"$options": "i",
		}
	}
	return i
}

func (i *initiatives) FilterDateBounds(dateFilters map[string]any, after bool) Initiatives {
	validDateFields := map[string]bool{
		"start_at":   true,
		"end_at":     true,
		"updated_at": true,
		"closed_at":  true,
	}

	var op string
	if after {
		op = "$gte"
	} else {
		op = "$lte"
	}

	for field, value := range dateFilters {
		if !validDateFields[field] {
			continue
		}
		if value == nil {
			continue
		}

		var t time.Time
		switch val := value.(type) {
		case time.Time:
			t = val
		case *time.Time:
			t = *val
		case string:
			parsed, err := time.Parse(time.RFC3339, val)
			if err != nil {
				logrus.Warnf("FilterDateBounds: cannot parse date '%s': %v", val, err)
				continue
			}
			t = parsed
		default:
			logrus.Warnf("FilterDateBounds: field %s is not a recognized date type: %T", field, value)
			continue
		}

		i.filters[field] = bson.M{op: t}
	}

	return i
}

func (i *initiatives) FilterCounts(countFilters map[string]any, greater bool) Initiatives {
	validCountFields := map[string]bool{
		"likes":   true,
		"reposts": true,
		"reports": true,
	}

	op := "$lte"
	if greater {
		op = "$gte"
	}

	for field, value := range countFilters {
		if !validCountFields[field] {
			continue
		}
		if value == nil {
			continue
		}

		var n int64
		switch val := value.(type) {
		case int:
			n = int64(val)
		case int64:
			n = val
		case float64:
			n = int64(val)
		case string:
			parsed, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				logrus.Warnf("FilterCounts: cannot parse int '%s': %v", val, err)
				continue
			}
			n = parsed
		default:
			logrus.Warnf("FilterCounts: field %s is not a recognized numeric type: %T", field, value)
			continue
		}
		i.filters[field] = bson.M{op: n}
	}

	return i
}
