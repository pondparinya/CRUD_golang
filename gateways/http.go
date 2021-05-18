package gateways

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pondparinya/CRUD_golang/dao"
	"github.com/pondparinya/CRUD_golang/services"
	"gopkg.in/go-playground/validator.v9"
)

type HTTP struct {
	Service services.IServicesCRUD
}

func NewHTTP(g *echo.Group, sv services.IServicesCRUD) {
	h := &HTTP{
		Service: sv,
	}
	g.POST("", h.HTTP_Create)

}

func (sv HTTP) HTTP_Create(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(dao.StudentDAO)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var Errors []string
	if err := c.Validate(r); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			err := fmt.Sprintf("Error invalid parameter (%v)", err.Field())
			Errors = append(Errors, err)
		}
		return c.JSON(http.StatusBadRequest, map[string][]string{
			"error": Errors,
		})
	}
	res, err := sv.Service.CreateStudent(ctx, dao.StudentDAO{
		StudentID: r.StudentID,
		Name:      r.Name,
		Nickname:  r.Nickname,
		Age:       r.Age,
		Address:   r.Address,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(201, res)
}
