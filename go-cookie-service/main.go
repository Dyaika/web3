package main

import (
	"go-cookie-service/maintenance"
	"go-cookie-service/storage-service"
)

func main() {
	maintenance.LoadEnv()
	//cookie_service.Start()
	storage_service.Start()
}
