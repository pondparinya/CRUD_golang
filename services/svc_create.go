package services

import (
	"context"
	"fmt"

	"github.com/pondparinya/CRUD_golang/dao"
	"github.com/pondparinya/CRUD_golang/database/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv ServicesCRUD) CreateStudent(ctx context.Context, d *dao.CreateStudentReq) (*dao.CreateStudentRes, error) {
	res, _ := sv.Repo.FindByStudentID(ctx, d.StudentID)
	if res != nil {
		return &dao.CreateStudentRes{}, fmt.Errorf("This StudentID is Already")
	}

	ent := &entity.StudentEntity{
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
		return &dao.CreateStudentRes{}, err
	}
	id := r.InsertedID.(primitive.ObjectID).Hex()
	return &dao.CreateStudentRes{
		ID: id,
	}, err
}
