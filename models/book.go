package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `db:"title" form:"title"`
	Author string `db:"author" form:"author"`
}

func Save(book *Book) error {
	return DB.Create(book).Error
}

func FindById(id string) (*Book, error) {
	var book Book
	err := DB.First(&book, id).Error

	fmt.Println("_____>", err)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func Delete(book *Book) error {
	return DB.Delete(book).Error
}

func FindAll() ([]Book, error) {
	var books []Book
	err := DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// func UpdateBook(id uint, updatedBook *Book) (*Book, error) {
// 	// Parse and validate the ID
// 	book, err := FindById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Update the book fields if they are provided in the request
// 	if updatedBook.Title != "" {
// 		book.Title = updatedBook.Title
// 	}
// 	if updatedBook.Author != "" {
// 		book.Author = updatedBook.Author
// 	}

// 	err = DB.Save(book).Error

// 	return book, err
// }
