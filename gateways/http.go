package gateways

import (
	"github.com/labstack/echo/v4"
	"github.com/pondparinya/CRUD_golang/services"
)

type HTTP struct {
	Service services.IServicesCRUD
}

func NewHTTP(g *echo.Group, sv services.IServicesCRUD) {
	h := &HTTP{
		Service: sv,
	}
	g.POST("", h.HTTP_Create)
	g.GET("", h.GetAll)
	g.GET("students", h.FindByID)
	g.DELETE("student", h.Delete)
	g.GET("student", h.FindByStudentID)
	g.PUT("student", h.UpdateStudent)

}
