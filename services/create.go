package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv ServicesCRUD) Create(ctx context.Context, d dao.StudentDAO) (string, error) {
	ent := entity.StudentEntity{
		StudentID: d.StudentID,
		Name:      d.Name,
		Nickname:  d.Nickname,
		Age:       d.Age,
		Address: &entity.AddressEntity{
			StreetAddress:  d.Address.StreetAddress,
			StreetAddress2: d.Address.StreetAddress2,
			City:           d.Address.City,
			Country:        d.Address.City,
		},
	}
	r, err := sv.Repo.InsertStudent(ctx, ent)
	if err != nil {
		return "No", err
	}
	id := r.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}
