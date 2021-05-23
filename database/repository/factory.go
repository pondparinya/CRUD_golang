package repository

import (
	"context"
	"fmt"

	"github.com/pondparinya/CRUD_golang/database/datasources"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Collection *mongo.Collection
}

type IRepository interface {
	InsertStudent(ctx context.Context, ent *entity.StudentEntity) (*mongo.InsertOneResult, error)
	FindAllStudent(ctx context.Context) (*[]entity.StudentEntity, error)
	FindByStudentID(ctx context.Context, studentID string) (*entity.StudentEntity, error)
	FindByID(ctx context.Context, ID primitive.ObjectID) (*entity.StudentEntity, error)
	UpdateStudent(ctx context.Context, ent *entity.StudentEntity) (*mongo.UpdateResult, error)
	DeleteByStudentID(ctx context.Context, StudentID string) *mongo.SingleResult
}

func NewRepository(ds *datasources.MongoDB) IRepository {
	return &Repository{
		Collection: ds.MongoDB.Database("DATABASE_NAME").Collection("COLLECTION_NAME"),
	}
}

func (repo Repository) InsertStudent(ctx context.Context, ent *entity.StudentEntity) (*mongo.InsertOneResult, error) {
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
}

func (repo Repository) FindByStudentID(ctx context.Context, studentID string) (*entity.StudentEntity, error) {
	var ent *entity.StudentEntity
	filter := bson.M{"student_id": studentID}
	err := repo.Collection.FindOne(ctx, filter).Decode(&ent)
	return ent, err
}

func (repo Repository) UpdateStudent(ctx context.Context, ent *entity.StudentEntity) (*mongo.UpdateResult, error) {
	fmt.Println(ent)
	return repo.Collection.UpdateByID(ctx, bson.M{
		"_id": ent.ID,
	},
		bson.M{"$set": bson.M{
			"student_id": ent.StudentID,
			"age":        ent.Age,
			"name":       ent.Name,
			"nickname":   ent.Nickname,
			"address":    ent.Address,
		},
		},
	)
}

func (repo Repository) DeleteByStudentID(ctx context.Context, StudentID string) *mongo.SingleResult {
	filter := bson.M{"student_id": StudentID}
	return repo.Collection.FindOneAndDelete(ctx, filter)
}

func (repo Repository) FindByID(ctx context.Context, ID primitive.ObjectID) (*entity.StudentEntity, error) {
	var ent *entity.StudentEntity
	err := repo.Collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&ent)
	return ent, err
}
