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
	CreateStudent(ctx context.Context, d *dao.CreateStudentDAO) (*dao.CreateStudentResDAO, error)
	GetAllStudent(ctx context.Context) ([]*dao.GetStudentResDAO, error)
	FindByStudentID(ctx context.Context, student_id int) (*dao.GetStudentResDAO, error)
	UpdateStudentByID(ctx context.Context, d *dao.GetStudentResDAO) error
}

func NewServuceCRUD(repo repository.IRepository) IServicesCRUD {
	return &ServicesCRUD{
		Repo: repo,
	}
}
