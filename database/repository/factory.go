package repository

import (
	"context"

	"github.com/pondparinya/CRUD_golang/database/datasources"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

type IRepository interface {
}

func NewRepository(ds *datasources.MongoDB) IRepository {
	return Repository{
		Collection: ds.MongoDB.Database("Database Name").Collection("Collection Name"),
	}
}


func InsertStudent(ctx context.Context , )
