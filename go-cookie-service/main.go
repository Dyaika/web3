package main

import (
	"go-cookie-service/cookie-service"
	"go-cookie-service/maintenance"
	"go-cookie-service/storage-service"
	"net/http"
)

func main() {
	maintenance.LoadEnv()
	maintenance.InitializeLogging()
	cookie_service.Start()
	storage_service.Start()
	port := maintenance.GetEnv("PORT", "8080")
	maintenance.LogData("Server is running on port " + port + "\n")
	http.ListenAndServe(":"+port, nil)
}
