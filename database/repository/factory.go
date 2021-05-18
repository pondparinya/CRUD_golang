package repository

import (
	"context"

	"github.com/pondparinya/CRUD_golang/database/datasources"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Collection *mongo.Collection
}

type IRepository interface {
	InsertStudent(ctx context.Context, ent entity.StudentEntity) (*mongo.InsertOneResult, error)
	FindAllStudent(ctx context.Context) (*[]entity.StudentEntity, error)
}

func NewRepository(ds *datasources.MongoDB) IRepository {
	return &Repository{
		Collection: ds.MongoDB.Database("Test").Collection("student"),
	}
}

func (repo Repository) InsertStudent(ctx context.Context, ent entity.StudentEntity) (*mongo.InsertOneResult, error) {
	return repo.Collection.InsertOne(ctx, &ent)
}

func (repo Repository) FindAllStudent(ctx context.Context) (*[]entity.StudentEntity, error) {
	option := options.Find()
	cur, err := repo.Collection.Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		return &[]entity.StudentEntity{}, err
	}
	var ent []entity.StudentEntity
	err = cur.All(ctx, &ent)
	return &ent, err
	// res := make([]entity.StudentEntity, 0)
	// for cur.Next(ctx) {
	// 	var item entity.StudentEntity
	// 	err := cur.Decode(&item)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	res = append(res, item)

	// }

}
