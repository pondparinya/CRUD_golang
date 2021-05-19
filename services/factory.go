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
	CreateStudent(ctx context.Context, d *dao.CreateStudentReq) (*dao.CreateStudentRes, error)
	GetAllStudent(ctx context.Context) ([]*dao.GetStudentRes, error)
	FindByStudentID(ctx context.Context, student_id string) (*dao.GetStudentRes, error)
	UpdateStudentByID(ctx context.Context, d *dao.GetStudentRes) error
	DeleteByStudentID(ctx context.Context, d *dao.DeleteStudentReq) error
}

func NewServuceCRUD(repo repository.IRepository) IServicesCRUD {
	return &ServicesCRUD{
		Repo: repo,
	}
}
