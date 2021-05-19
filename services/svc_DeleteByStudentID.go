package services

import (
	"context"

	"github.com/pondparinya/CRUD_golang/dao"
)

func (sv ServicesCRUD) DeleteByStudentID(ctx context.Context, d *dao.DeleteStudentReq) error {
	err := sv.Repo.DeleteByStudentID(ctx, d.StudentID)
	if err.Err().Error() == "mongo: no documents in result" {
		return err.Err()
	}
	return nil
}
