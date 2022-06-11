package storage

import (
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
	defer m.driver.Close(client, ctx, cancel)
	return m.driver.Insert(ctx, client, model)

}

func (m *mStorage) Retrieve() ([]InvoiceData, error) {
	client, ctx, cancel, err := m.driver.Connect()
	if err != nil {
		panic(err)
	}
	defer m.driver.Close(client, ctx, cancel)
	cursor, errFind := m.driver.Retrieve(ctx, client)
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
