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
	Create(ctx context.Context, d dao.StudentDAO) (string, error)
}

func NewServuceCRUD(repo repository.IRepository) IServicesCRUD {
	return &ServicesCRUD{
		Repo: repo,
	}
}
