package storage

import (
	"context"

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

func NewUriStorage(uri string) Storage {
	driver := internal.NewDriver(uri)
	return &mStorage{
		driver: driver,
	}
}

func NewDriverStorage(d internal.Driver) Storage {
	return &mStorage{
		driver: d,
	}
}

func (m *mStorage) Insert(model interface{}) (interface{}, error) {
	ctx, cancel, err := m.driver.Connect()
	if err != nil {
		return nil, err
	}
	defer m.driver.Close(ctx, cancel)
	return m.driver.Insert(ctx, model)
}

func (m *mStorage) Retrieve() ([]InvoiceData, error) {
	ctx, cancel, err := m.driver.Connect()
	if err != nil {
		return nil, err
	}
	defer m.driver.Close(ctx, cancel)
	cursor, errFind := m.driver.Retrieve(ctx)
	if errFind != nil {
		return nil, errFind
	}
	defer cursor.Close(context.TODO())
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
