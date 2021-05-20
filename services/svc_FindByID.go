package services

import (
	"context"
	"fmt"

	"github.com/pondparinya/CRUD_golang/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sv ServicesCRUD) FindStudentByID(ctx context.Context, ID string) (*dao.GetStudentRes, error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return &dao.GetStudentRes{}, err
	}

	r, err := sv.Repo.FindByID(ctx, id)
	if err != nil {
		return &dao.GetStudentRes{}, err
	}
	fmt.Println(r)
	res := &dao.GetStudentRes{
		ID:        r.ID.Hex(),
		StudentID: r.StudentID,
		Name:      r.Name,
		Nickname:  r.Nickname,
		Age:       r.Age,
		Address: dao.AddressDAO{
			StreetAddress:  r.Address.StreetAddress,
			StreetAddress2: r.Address.StreetAddress2,
			State:          r.Address.State,
			City:           r.Address.City,
			Zipcode:        r.Address.Zipcode,
			Country:        r.Address.Country,
		},
	}
	return res, err

}
