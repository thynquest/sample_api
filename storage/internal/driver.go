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
	Connect() (context.Context, context.CancelFunc, error)
	Close(ctx context.Context, cancel context.CancelFunc)
	Insert(ctx context.Context, data interface{}) (interface{}, error)
	Retrieve(ctx context.Context) (*mongo.Cursor, error)
}

type mongoDriver struct {
	uri    string
	client *mongo.Client
}

func NewDriver(uri string) Driver {
	return &mongoDriver{
		uri: uri,
	}
}

func (m *mongoDriver) Connect() (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	m.client = client
	return ctx, cancel, err
}

func (m *mongoDriver) Close(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := m.client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (m *mongoDriver) Insert(ctx context.Context, data interface{}) (interface{}, error) {
	db := common.GetEnv("DB", "test")
	col := common.GetEnv("COLLECTION", "mycol")
	collection := m.client.Database(db).Collection(col)
	insert, errInsert := collection.InsertOne(ctx, data)
	if errInsert != nil {
		return nil, errInsert
	}
	return insert.InsertedID, nil
}

func (m *mongoDriver) Retrieve(ctx context.Context) (*mongo.Cursor, error) {
	db := common.GetEnv("DB", "test")
	col := common.GetEnv("COLLECTION", "mycol")
	options := options.Find()
	collection := m.client.Database(db).Collection(col)
	cursor, errFind := collection.Find(ctx, bson.D{{}}, options)
	return cursor, errFind
}
