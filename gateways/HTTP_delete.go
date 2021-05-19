package gateways

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pondparinya/CRUD_golang/dao"
	"gopkg.in/go-playground/validator.v9"
)

func (sv HTTP) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(dao.DeleteStudentReq)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var Errors []string
	if err := c.Validate(r); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			err := fmt.Sprintf("Error invalid paremeter (%v)", err.Field())
			Errors = append(Errors, err)
		}
		return c.JSON(http.StatusBadRequest, map[string][]string{"error": Errors})
	}
	err := sv.Service.DeleteByStudentID(ctx, &dao.DeleteStudentReq{
		StudentID: r.StudentID,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, map[string]string{"message": "delete complete"})
}
