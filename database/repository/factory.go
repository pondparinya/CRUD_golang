package repository

import (
	"context"

	"github.com/pondparinya/CRUD_golang/database/datasources"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

type IRepository interface {
	InsertStudent(ctx context.Context, ent entity.StudentEntity) (*mongo.InsertOneResult, error)
}

func NewRepository(ds *datasources.MongoDB) IRepository {
	return &Repository{
		Collection: ds.MongoDB.Database("Test").Collection("student"),
	}
}

func (repo Repository) InsertStudent(ctx context.Context, ent entity.StudentEntity) (*mongo.InsertOneResult, error) {
	return repo.Collection.InsertOne(ctx, &ent)
}
