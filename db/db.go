package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     *sync.Once
)

func GetClient() (*mongo.Client, error) {
	var err error

	clientOnce.Do(func() {
		uri := os.Getenv("MONGO_URI")
		clientOptions := options.Client().ApplyURI(uri)
		clientInstance, err = mongo.Connect(context.TODO(), clientOptions)
	})

	return clientInstance, err
}

func GetCollection(name string) (*mongo.Collection, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}
	return client.Database("blogdb").Collection(name), nil
}
