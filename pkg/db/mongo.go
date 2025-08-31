package db

import (
	"context"
	"log"
	"time"

	"github.com/ArjunDev17/course-content-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo(ctx context.Context) (*mongo.Client, error) {
	if Client != nil {
		return Client, nil
	}

	uri := config.Cfg.Mongo.URI
	clientOpts := options.Client().ApplyURI(uri)

	ctxConn, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxConn, clientOpts)
	if err != nil {
		return nil, err
	}

	// ping
	ctxPing, cancelPing := context.WithTimeout(ctx, 2*time.Second)
	defer cancelPing()
	if err := client.Ping(ctxPing, nil); err != nil {
		return nil, err
	}

	Client = client
	log.Println("Connected to MongoDB:", uri)
	return Client, nil
}

func CoursesCollection() *mongo.Collection {
	return Client.Database(config.Cfg.Mongo.Database).Collection(config.Cfg.Mongo.CoursesCollection)
}
