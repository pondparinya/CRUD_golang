package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pondparinya/CRUD_golang/database/datasources"
	"github.com/pondparinya/CRUD_golang/database/repository"
	gateways "github.com/pondparinya/CRUD_golang/gateways"
	"github.com/pondparinya/CRUD_golang/services"
	validator "gopkg.in/go-playground/validator.v9"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	e := echo.New()

	// Set echo with validator
	e.Validator = &RequestValidator{Validator: validator.New()}
	db := datasources.NewMongoDB()
	repo := repository.NewRepository(db)
	sv := services.NewServuceCRUD(repo)

	routG := e.Group("/")

	gateways.NewHTTP(routG, sv)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
