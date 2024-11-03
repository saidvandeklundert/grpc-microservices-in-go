package main

// https://github.com/huseyinbabal/microservices/blob/main/order/cmd/main.go

import (
	"log"

	"github.com/saidvandeklundert/microservices/order/config"
	"github.com/saidvandeklundert/microservices/order/internal/adapters/db"
	"github.com/saidvandeklundert/microservices/order/internal/adapters/grpc"
	"github.com/saidvandeklundert/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}

/*
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/order \
APPLICATION_PORT=3000 \
ENV=development \
go run cmd/main.go
*/
