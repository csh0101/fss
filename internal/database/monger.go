package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongoer struct {
	client   *mongo.Client
	dbName   string
	collName string
}

type MongerOptions func(monger *Mongoer)

func DBNameOptions(dbName string) MongerOptions {
	return func(monger *Mongoer) {
		monger.dbName = dbName
	}
}

func CollNameOptions(collName string) MongerOptions {
	return func(monger *Mongoer) {
		monger.collName = collName
	}
}

func InitMonger(uri string, mogerOptions ...MongerOptions) *Mongoer {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	m := &Mongoer{
		client:   client,
		dbName:   "test",
		collName: "test",
	}
	for _, opt := range mogerOptions {
		opt(m)
	}
	return m
}

func (mogo *Mongoer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return mogo.client.Disconnect(ctx)
}

func (mogo *Mongoer) Insert(ctx context.Context, dbName, collName string, data primitive.M) error {
	coll := mogo.client.Database(dbName).Collection(collName)
	if _, err := coll.InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}

func (mogo *Mongoer) QueryWithCursor(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	return mogo.queryWithCursor(ctx, mogo.dbName, mogo.collName, filter)
}

func (mogo *Mongoer) queryWithCursor(ctx context.Context, dbName, collName string, filter interface{}) (*mongo.Cursor, error) {
	coll := mogo.client.Database(dbName).Collection(collName)
	return coll.Find(ctx, filter)
}

func (mogo *Mongoer) QueryOne(ctx context.Context, dbName, collName string, filter interface{}, data interface{}) error {
	coll := mogo.client.Database(dbName).Collection(collName)
	res := coll.FindOne(ctx, filter)
	if res.Err() != nil {
		return res.Err()
	}
	return res.Decode(data)
}

// func (mogo *Mongoer) QueryMany(ctx context.Context, dbName, collName string, option *options.FindOptions, filter interface{}, data []interface{}) error {
// 	if data == nil {
// 		return errors.New("the input data is nil,should be a type array")
// 	}
// 	coll := mogo.client.Database(dbName).Collection(collName)
// 	cursor, err := coll.Find(ctx, filter)
// 	if err != nil {
// 		return err
// 	}
// 	for cursor.Next(ctx) {
// 		cursor.Decode()
// 	}
// }
