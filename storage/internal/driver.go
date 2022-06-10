package internal

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Driver interface {
	Connect() (*mongo.Client, context.Context, context.CancelFunc, error)
	Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc)
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
