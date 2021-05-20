package gateways

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (sv HTTP) FindByID(c echo.Context) error {
	ID := c.QueryParam("_id")
	res, err := sv.Service.FindStudentByID(c.Request().Context(), ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, res)
}
