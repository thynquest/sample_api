package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"spin.sample.trial/common"
	"spin.sample.trial/routes"
	"spin.sample.trial/service"
	"spin.sample.trial/storage"
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
