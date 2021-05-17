package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pondparinya/CRUD_golang/database/datasources"
	"github.com/pondparinya/CRUD_golang/database/repository"
	gateways "github.com/pondparinya/CRUD_golang/gateways"
	"github.com/pondparinya/CRUD_golang/services"
)

func main() {
	e := echo.New()
	db := datasources.NewMongoDB()

	repo := repository.NewRepository(db)

	sv := services.NewServuceCRUD(repo)

	routG := e.Group("/")

	gateways.NewHTTP(routG, sv)

	e.Logger.Fatal(e.Start(":3000"))

}
