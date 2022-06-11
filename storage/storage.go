package storage

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"spin.sample.trial/common"
	"spin.sample.trial/storage/internal"
)

type InvoiceData struct {
	Amount    float64
	Company   string
	IssueDate string
	DueDate   string
}

type Storage interface {
	Insert(model interface{}) (interface{}, error)
	Retrieve() ([]InvoiceData, error)
}

type mStorage struct {
	driver internal.Driver
}

func NewStorage(uri string) Storage {
	driver := internal.NewDriver(uri)
	return &mStorage{
		driver: driver,
	}
}

func (m *mStorage) Insert(model interface{}) (interface{}, error) {
	client, ctx, cancel, err := m.driver.Connect()
	if err != nil {
		panic(err)
	}
	db := common.GetEnv("DB", "test")
	col := common.GetEnv("COLLECTION", "mycol")
	defer m.driver.Close(client, ctx, cancel)
	collection := client.Database(db).Collection(col)
	insert, errInsert := collection.InsertOne(ctx, model)

	return insert.InsertedID, errInsert
}

func (m *mStorage) Retrieve() ([]InvoiceData, error) {
	client, ctx, cancel, err := m.driver.Connect()
	if err != nil {
		panic(err)
	}
	db := common.GetEnv("DB", "test")
	col := common.GetEnv("COLLECTION", "mycol")
	defer m.driver.Close(client, ctx, cancel)
	options := options.Find()
	collection := client.Database(db).Collection(col)
	cursor, errFind := collection.Find(ctx, bson.D{{}}, options)
	if errFind != nil {
		return nil, errFind
	}
	var results []InvoiceData
	for cursor.Next(ctx) {
		var item InvoiceData
		errDec := cursor.Decode(&item)
		if errDec != nil {
			return nil, errDec
		}
		results = append(results, item)
	}
	return results, nil
}
