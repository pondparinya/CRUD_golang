package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
)

func (sv ServicesCRUD) FindByStudentID(ctx context.Context, student_id int) (*dao.GetStudentResDAO, error) {
	r, err := sv.Repo.FindByStudentID(ctx, student_id)
	if err != nil {
		return &dao.GetStudentResDAO{}, err
	}
	res := &dao.GetStudentResDAO{
		ID:        r.ID.Hex(),
		StudentID: r.StudentID,
		Name:      r.Name,
		Nickname:  r.Nickname,
		Age:       r.Age,
		Address: dao.AddressDAO{
			StreetAddress:  r.Address.StreetAddress,
			StreetAddress2: r.Address.StreetAddress2,
			City:           r.Address.City,
			Country:        r.Address.Country,
			Zipcode:        r.Address.Zipcode,
			State:          r.Address.State,
		},
	}
	return res, err
}
