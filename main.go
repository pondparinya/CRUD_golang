package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pondparinya/CRUD_golang/database/datasources"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
