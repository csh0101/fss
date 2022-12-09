package database_test

import (
	"context"
	"fmt"
	"fss/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	uri    = "mongodb://localhost:27017"
	dbName = "test"
	coll   = "test"
)

func TestMonger(t *testing.T) {
	monger := database.InitMonger(uri)

	defer monger.Close()

	{
		err := monger.Insert(context.TODO(), dbName, coll, primitive.M{
			"title": "this is a nice book",
			"name":  "the book of csh0101",
		})
		assert.Equal(t, nil, err)
	}

	{
		A := new(struct {
			Title string `bson:"title"`
			Name  string `bson:"name"`
		})
		err := monger.QueryOne(context.TODO(), dbName, coll, primitive.M{"title": "this is a nice book"}, A)
		assert.Equal(t, nil, err)
		var expected = &struct {
			Title string `bson:"title"`
			Name  string `bson:"name"`
		}{
			Title: "this is a nice book",
			Name:  "the book of csh0101",
		}
		assert.EqualValues(t, expected, A)
		fmt.Println(A)
	}
}
