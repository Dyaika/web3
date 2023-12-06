package storage

import (
	"context"
	"go-cookie-service/maintenance"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection(uri string, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
		return nil, err
	}

	db := client.Database(dbName)
	return db, nil
}
