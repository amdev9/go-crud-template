package main

import (
	"go-crud-template/models"
	"go-crud-template/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := models.InitDB("books.db")
	models.Migrate(db)

	api := e.Group("/api")
	routes.SetupRoutes(api)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
