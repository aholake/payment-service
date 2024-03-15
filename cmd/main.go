package main

import (
	"log"

	"github.com/aholake/payment-service/config"
	"github.com/aholake/payment-service/internal/adapters/db"
	adapters "github.com/aholake/payment-service/internal/adapters/gprc"
	"github.com/aholake/payment-service/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewDBAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("unable to connect DB: %v", err)
	}

	apiPort := api.NewApplication(dbAdapter)
	server := adapters.NewGRPCAdapter(int32(config.GetApplicationPort()), apiPort)

	server.Run()
}
