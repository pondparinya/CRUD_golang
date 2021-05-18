package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
)

func (sv ServicesCRUD) GetAllStudent(ctx context.Context) ([]dao.StudentDAO, error) {
	Res := make([]dao.StudentDAO, 0)
	res, err := sv.Repo.FindAllStudent(ctx)
	if err != nil {
		return []dao.StudentDAO{}, err
	}
	for _, r := range *res {
		data := dao.StudentDAO{
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
		Res = append(Res, data)
	}
	return Res, err
}
