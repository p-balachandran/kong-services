package db

import (
	"context"

	"github.com/p-balachandran/kong-services/server/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://${DB_USER}:${DB_PASSWORD}@mongo:27017"

type MongoDB struct {
	mc *mongo.Client
}

func NewMongoDB(ctx context.Context) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		mc: client,
	}, nil
}

func (c *MongoDB) Init(ctx context.Context) error {
	coll := c.mc.Database("pbkong_db").Collection("services")
	docs := []interface{}{
		models.Service{},
	}
	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}

	return nil
}
