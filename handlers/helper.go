package handlers

import (
	"errors"
	"go-crud-template/models"
)

func validateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}

	if book.Author == "" {
		return errors.New("author is required")
	}

	return nil
}
