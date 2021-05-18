package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
	"github.com/pondparinya/CRUD_golang/database/repository"
)

type ServicesCRUD struct {
	Repo repository.IRepository
}

type IServicesCRUD interface {
	CreateStudent(ctx context.Context, d dao.StudentDAO) (dao.StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]dao.StudentDAO, error)
}

func NewServuceCRUD(repo repository.IRepository) IServicesCRUD {
	return &ServicesCRUD{
		Repo: repo,
	}
}
