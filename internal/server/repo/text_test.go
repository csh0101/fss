package repo_test

import (
	"context"
	"fmt"
	"fss/internal/database"
	"fss/internal/domain"
	"fss/internal/server/repo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	uri      = "mongodb://localhost:27017"
	dbName   = "test"
	collName = "test"
)

func TestTextRepo(t *testing.T) {

	flag := true
	monger := database.InitMonger(uri)
	data, err := bson.Marshal(&domain.Text{
		CreateTime: uint64(time.Now().Unix()),
		Content:    "00:41 a nice time",
		Name:       "text file 1",
	})
	assert.Equal(t, nil, err)

	m := primitive.M{}
	{
		err := bson.Unmarshal(data, &m)
		assert.Equal(t, nil, err)
	}

	{
		if flag {
			err := monger.Insert(context.TODO(), dbName, collName, m)
			assert.Equal(t, nil, err)
		}
	}

	{
		filter := &domain.TextFilter{KeyWord: []string{"a", "nice"}}
		repo := repo.NewTextRepo(monger)
		res, err := repo.QueryTextByFilter(context.TODO(), filter)
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, len(res))
		fmt.Println(len(res))
		fmt.Println(res[len(res)-1].CreateTime)
	}

	{
		filter := &domain.TextFilter{KeyWord: []string{"y", "b"}}
		repo := repo.NewTextRepo(monger)
		res, err := repo.QueryTextByFilter(context.TODO(), filter)
		assert.Equal(t, 0, len(res))
		assert.Equal(t, nil, err)
	}

	{

		// data, err := bson.Marshal(&domain.Text{
		// 	CreateTime: 5,
		// 	Content:    "00:41 a nice time",
		// 	Name:       "text file 1",
		// })
		// assert.Equal(t, nil, err)
		// m := primitive.M{}
		// {
		// 	err := bson.Unmarshal(data, &m)
		// 	assert.Equal(t, nil, err)
		// }
		// err = monger.Insert(context.TODO(), dbName, collName, m)
		// assert.Equal(t, nil, err)
		filter := &domain.TextFilter{StartTime: 5, EndTime: 5, KeyWord: []string{"a", "x"}}
		repo := repo.NewTextRepo(monger)
		res, err := repo.QueryTextByFilter(context.TODO(), filter)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, nil, err)
		{
			filter := &domain.TextFilter{StartTime: 8, EndTime: 10, KeyWord: []string{"a", "x"}}
			res, err := repo.QueryTextByFilter(context.TODO(), filter)
			assert.Equal(t, 0, len(res))
			assert.Equal(t, nil, err)

		}
	}

}
