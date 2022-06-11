package service

import (
	"spin.sample.trial/storage"
)

type Invoice interface {
	Create(companyName string, amount float64, issueDate, dueDate string) (bool, error)
	Retrieve() ([]storage.InvoiceData, error)
}

type invoice struct {
	storage storage.Storage
}

func NewInvoiceSvc(storage storage.Storage) Invoice {
	return &invoice{
		storage: storage,
	}
}

func (i *invoice) Create(companyName string, amount float64, issueDate, dueDate string) (bool, error) {
	data := &storage.InvoiceData{
		Amount:    amount,
		Company:   companyName,
		IssueDate: issueDate,
		DueDate:   dueDate,
	}
	result, errResult := i.storage.Insert(data)
	if errResult != nil {
		return false, errResult
	}
	return (result != nil), nil
}

func (i *invoice) Retrieve() ([]storage.InvoiceData, error) {
	results, errResults := i.storage.Retrieve()
	return results, errResults
}
