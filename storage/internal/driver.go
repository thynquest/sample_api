package internal

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"spin.sample.trial/common"
)

type Driver interface {
	Connect() (*mongo.Client, context.Context, context.CancelFunc, error)
	Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc)
	Insert(ctx context.Context, client *mongo.Client, data interface{}) (interface{}, error)
	Retrieve(ctx context.Context, client *mongo.Client) (*mongo.Cursor, error)
}

type mongoDriver struct {
	uri string
}

func NewDriver(uri string) Driver {
	return &mongoDriver{
		uri: uri,
	}
}

func (m *mongoDriver) Connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	return client, ctx, cancel, err
}

func (m *mongoDriver) Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (m *mongoDriver) Insert(ctx context.Context, client *mongo.Client, data interface{}) (interface{}, error) {
	db := common.GetEnv("DB", "test")
	col := common.GetEnv("COLLECTION", "mycol")
	collection := client.Database(db).Collection(col)
	insert, errInsert := collection.InsertOne(ctx, data)
	if errInsert != nil {
		return nil, errInsert
	}
	return insert.InsertedID, nil
}

func (m *mongoDriver) Retrieve(ctx context.Context, client *mongo.Client) (*mongo.Cursor, error) {
	db := common.GetEnv("DB", "test")
	col := common.GetEnv("COLLECTION", "mycol")
	options := options.Find()
	collection := client.Database(db).Collection(col)
	cursor, errFind := collection.Find(ctx, bson.D{{}}, options)
	return cursor, errFind
}
