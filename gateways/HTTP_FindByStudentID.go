package gateways

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (sv HTTP) FindByStudentID(c echo.Context) error {
	StudentID := c.QueryParam("student_id")
	id, _ := strconv.Atoi(StudentID)
	res, err := sv.Service.FindByStudentID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, res)
}
