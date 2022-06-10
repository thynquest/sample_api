package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"u-trade.sample.trial/common"
	"u-trade.sample.trial/routes"
	"u-trade.sample.trial/service"
	"u-trade.sample.trial/storage"
)

func main() {

	host := common.GetEnv("API_HOST", "localhost")
	port := common.GetEnv("PORT", "8090")
	storageUrl := common.GetEnv("STORAGE", "mongodb://localhost:27017")
	target := fmt.Sprintf("%s:%s", host, port)
	r := gin.Default()
	dbStorage := storage.NewStorage(storageUrl)
	invoiceService := service.NewInvoiceSvc(dbStorage)
	routes.Invoice(r, invoiceService)
	routes.Invoices(r, invoiceService)
	if errRun := r.Run(target); errRun != nil {
		panic(errRun)
	}
}
