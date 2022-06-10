package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"u-trade.sample.trial/model"
	"u-trade.sample.trial/service"
)

type InvoiceHandler struct {
	Service service.Invoice
}

type Invoice struct {
	Amount    float64 `json:"amount"`
	Company   string  `json:"company"`
	IssueDate string  `json:"issuedate"`
	DueDate   string  `json:"duedate"`
}

func (i *InvoiceHandler) Create(c *gin.Context) {
	var data Invoice
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	result, errResult := i.Service.Create(data.Company, data.Amount, data.IssueDate, data.DueDate)
	if errResult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errResult.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (i *InvoiceHandler) Retrieve(c *gin.Context) {
	items, errResults := i.Service.Retrieve()
	if errResults != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errResults.Error()})
		return
	}
	results := make([]model.Invoice, len(items))
	for i, item := range items {
		result := model.Invoice{
			Amount:    item.Amount,
			Company:   item.Company,
			IssueDate: item.IssueDate,
			DueDate:   item.DueDate,
		}
		results[i] = result
	}
	c.JSON(http.StatusOK, results)
}
