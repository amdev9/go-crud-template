package routes

import (
	"go-crud-template/handlers"
	"go-crud-template/models"
	"go-crud-template/repository"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(g *echo.Group) {
	// Routes for the users module
	repo := repository.NewGormBookRepository(models.DB)
	bookHandler := handlers.NewBookHandler(repo)

	g.GET("/books", bookHandler.GetAllBooks)
	g.POST("/books", bookHandler.CreateBook)
	g.GET("/books/:id", bookHandler.GetBook)
	g.PUT("/books/:id", bookHandler.UpdateBook)
	g.DELETE("/books/:id", bookHandler.DeleteBook)
}
