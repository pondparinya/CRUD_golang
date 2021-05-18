package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv ServicesCRUD) UpdateStudentByID(ctx context.Context, d *dao.GetStudentResDAO) error {
	id, err := primitive.ObjectIDFromHex(d.ID)
	ent := &entity.StudentEntity{
		ID:        id,
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
			Country:        d.Address.Country,
		},
	}

	_, err = sv.Repo.UpdateStudent(ctx, ent)
	if err != nil {
		return err
	}
	return err
}
