package main

import (
	"go-crud-template/handlers"
	"go-crud-template/models"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := models.InitDB("books.sqlite")
	models.Migrate(db)

	e.Static("/public", "public")

	e.GET("/books", handlers.BooksGet)
	e.GET("/books/add", handlers.BookAddGet)
	e.POST("/books/add", handlers.BookAddPost)
	e.POST("/books/:id/delete", handlers.BookDelete)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
