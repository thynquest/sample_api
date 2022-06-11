package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"spin.sample.trial/service"
	"spin.sample.trial/storage"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type MockStorage struct{}

func (m *MockStorage) Insert(model interface{}) (interface{}, error) {
	return struct{}{}, nil
}

func (m *MockStorage) Retrieve() ([]storage.InvoiceData, error) {
	mockitem := storage.InvoiceData{
		Amount:    12345,
		Company:   "MyMockCompany",
		IssueDate: "2022-04-25",
		DueDate:   "2022-04-29",
	}
	mockresult := []storage.InvoiceData{
		mockitem,
	}
	return mockresult, nil
}

func TestCreateInvoice(t *testing.T) {
	ms := &MockStorage{}
	invoiceSvc := service.NewInvoiceSvc(ms)
	testHandler := &InvoiceHandler{
		Service: invoiceSvc,
	}
	r := SetUpRouter()
	r.POST("/invoice", testHandler.Create)
	payload := Invoice{
		Amount:    1234,
		Company:   "myTestCompany",
		IssueDate: "2022-04-25",
		DueDate:   "2022-04-29",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/invoice", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if http.StatusCreated != w.Code {
		t.Errorf("the status should be %v got %v", http.StatusCreated, w.Code)
	}
}

func TestRetrieveInvoice(t *testing.T) {
	ms := &MockStorage{}
	invoiceSvc := service.NewInvoiceSvc(ms)
	testHandler := &InvoiceHandler{
		Service: invoiceSvc,
	}
	r := SetUpRouter()
	r.GET("/invoices", testHandler.Retrieve)
	req, _ := http.NewRequest("GET", "/invoices", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var invoices []Invoice
	json.Unmarshal(w.Body.Bytes(), &invoices)
	if http.StatusOK != w.Code {
		t.Errorf("the status should be %v got %v", http.StatusCreated, w.Code)
	}
	if len(invoices) != 1 {
		t.Errorf("there should be 1 invoice got %d", len(invoices))
	}
}
