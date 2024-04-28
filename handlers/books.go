package handlers

import (
	"fmt"
	"net/http"

	"go-crud-template/lib"
	"go-crud-template/models"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func BooksGet(c echo.Context) error {
	books, err := models.FindAll()
	if err != nil {
		log.Error("Failed to retrieve books. ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve books.")
	}

	return lib.Render(c, 200, BooksList(books))
}

func BookAddGet(c echo.Context) error {
	return lib.Render(c, 200, BookAdd(nil, nil))
}

func BookAddPost(c echo.Context) (err error) {
	book := new(models.Book)

	if err = c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(book); err != nil {
		return err
	}

	err = models.Save(book)

	if err != nil {
		log.Error("Failed to create the book. ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create the book.")
	}

	if err == nil {
		return lib.HtmxRedirect(c, "/books")
	}

	fmt.Println("Bookadd", book)
	component := lib.HtmxRender(
		c,
		func() templ.Component { return BookAddForm(book, err) },
		func() templ.Component { return BookAdd(book, err) },
	)

	return lib.Render(c, 200, component)
}

func BookDelete(c echo.Context) error {

	id := c.Param("id")

	book, err := models.FindById(id)

	fmt.Println("----<", book, err)
	if err == nil {
		err = models.Delete(book)
	}

	if err != nil {
		c.NoContent(400)
		return nil
	}

	return lib.HtmxRedirect(c, "/books")
}

// func GetBook(c echo.Context) error {
// 	fmt.Println("c", c)

// 	id, err := getIntId(c)
// 	if err != nil {
// 		log.Error("Invalid book ID: ", err.Error())
// 		return c.JSON(http.StatusBadRequest, "Invalid book ID")
// 	}

// 	book, err := models.GetBookByID(id)
// 	if err != nil {
// 		log.Error("Book not found ", err.Error())
// 		return c.JSON(http.StatusNotFound, "Book not found")
// 	}

// 	return c.JSON(http.StatusOK, book)
// }

// func UpdateBook(c echo.Context) error {
// 	id, err := getIntId(c)
// 	if err != nil {
// 		log.Error("Invalid book ID: ", err.Error())
// 		return c.JSON(http.StatusBadRequest, "Invalid book ID")
// 	}
// 	// Create a new dto.Book instance and bind the request data to it
// 	var updatedBook dto.Book
// 	if err := c.Bind(&updatedBook); err != nil {
// 		log.Error("Invalid request data ", err.Error())
// 		return c.JSON(http.StatusBadRequest, "Invalid request data")
// 	}

// 	book := new(models.Book)
// 	if book, err = models.UpdateBook(id, &updatedBook); err != nil {
// 		log.Error("Failed to update the book ", err.Error())
// 		return c.JSON(http.StatusInternalServerError, "Failed to update the book.")
// 	}

// 	return c.JSON(http.StatusOK, book)
// }

// func DeleteBook(c echo.Context) error {
// 	id, err := getIntId(c)
// 	if err != nil {
// 		log.Error("Invalid book ID: ", err.Error())
// 		return c.JSON(http.StatusBadRequest, "Invalid book ID")
// 	}

// 	if err := models.DeleteBook(id); err != nil {
// 		log.Error("Book not found ", err.Error())
// 		return c.JSON(http.StatusNotFound, "Book not found")
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }
