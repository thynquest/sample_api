package routes

import (
	"github.com/gin-gonic/gin"
	"spin.sample.trial/handler"
	"spin.sample.trial/service"
)

func Invoice(router *gin.Engine, service service.Invoice) {
	handler := &handler.InvoiceHandler{
		Service: service,
	}
	router.POST("/invoice", handler.Create)
}

func Invoices(router *gin.Engine, service service.Invoice) {
	handler := &handler.InvoiceHandler{
		Service: service,
	}
	router.GET("/invoices", handler.Retrieve)
}
