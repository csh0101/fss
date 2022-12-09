package repo

import (
	"context"
	"fss/internal/database"
	"fss/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ domain.ITextRepo = new(text)

type text struct {
	monger *database.Mongoer
}

func NewTextRepo(monger *database.Mongoer) domain.ITextRepo {
	return &text{
		monger: monger,
	}
}

func (t *text) QueryTextByFilter(ctx context.Context, filter *domain.TextFilter) ([]*domain.Text, error) {

	var (
		keywordsCondtion, startCondition, endCondition primitive.M
		condtions                                      []primitive.M
	)
	if len(filter.KeyWord) > 0 {
		var regexExpr string = "("
		flag := false
		for i := 0; i < len(filter.KeyWord); i++ {
			if flag {
				regexExpr += "|"
			} else {
				flag = true
			}
			regexExpr += filter.KeyWord[i]
		}
		regexExpr += ")"
		keywordsCondtion = primitive.M{
			"content": primitive.Regex{
				Pattern: regexExpr,
			},
		}
		condtions = append(condtions, keywordsCondtion)
	}

	if filter.StartTime != 0 {
		startCondition = primitive.M{
			"create_time": primitive.M{
				"$gte": filter.StartTime,
			},
		}
		condtions = append(condtions, startCondition)
	}
	if filter.EndTime != 0 {
		endCondition = primitive.M{
			"create_time": primitive.M{
				"$lte": filter.EndTime,
			},
		}
		condtions = append(condtions, endCondition)
	}

	var f primitive.M
	length := len(condtions)
	switch {
	case length == 0:
		f = primitive.M{}
	case length == 1:
		f = condtions[0]
	default:
		f = primitive.M{"$and": condtions}
	}
	cursor, err := t.monger.QueryWithCursor(ctx, f)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*domain.Text, 0)

	for cursor.Next(ctx) {
		text := &domain.Text{}
		if err := cursor.Decode(text); err != nil {
			return nil, err
		}
		res = append(res, text)
	}
	return res, nil
}
