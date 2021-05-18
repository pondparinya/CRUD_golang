package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv ServicesCRUD) CreateStudent(ctx context.Context, d dao.StudentDAO) (dao.StudentResponse, error) {
	ent := entity.StudentEntity{
		StudentID: d.StudentID,
		Name:      d.Name,
		Nickname:  d.Nickname,
		Age:       d.Age,
		Address: &entity.AddressEntity{
			StreetAddress:  d.Address.StreetAddress,
			StreetAddress2: d.Address.StreetAddress2,
			State:          d.Address.State,
			Zipcode:        d.Address.Zipcode,
			City:           d.Address.City,
			Country:        d.Address.City,
		},
	}
	r, err := sv.Repo.InsertStudent(ctx, ent)
	if err != nil {
		return dao.StudentResponse{}, err
	}
	id := r.InsertedID.(primitive.ObjectID).Hex()
	return dao.StudentResponse{
		ID: id,
	}, err
}
