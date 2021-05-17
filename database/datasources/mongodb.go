package datasources

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Context context.Context
	MongoDB *mongo.Client
}

type IMongoDB interface {
	Disconnect()
}

func NewMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Your URI MongoDB
	option := options.Client().ApplyURI("mongodb+srv://ecommerce:myFirstDatabase@cluster0.idq5h.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	client, err0 := mongo.Connect(ctx, option)
	if err0 != nil {
		log.Fatal(err0)
	}

	return &MongoDB{
		Context: ctx,
		MongoDB: client,
	}

}

func (db MongoDB) Disconnect() {
	db.MongoDB.Disconnect(db.Context)
}
