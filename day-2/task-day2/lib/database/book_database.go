package database

import (
	"strconv"

	"github.com/Jiran03/agmc/task/day2/models"
	"github.com/labstack/echo/v4"
)

var books = []models.Book{
	{
		ID:              1,
		Title:           "title 1",
		Writer:          "writer 1",
		Publisher:       "publisher 1",
		PublicationYear: 2022,
	},
	{
		ID:              2,
		Title:           "title 2",
		Writer:          "writer 2",
		Publisher:       "publisher 2",
		PublicationYear: 2022,
	},
}

//get all Book from DB
func GetBooks() ([]models.Book, error) {
	return books, nil
}

//create new Book
func CreateBook(ctx echo.Context) (models.Book, error) {
	book := models.Book{
		ID: len(books) + 1,
	}

	ctx.Bind(&book)

	return book, nil
}

//get Book by id
func GetBookByID(ctx echo.Context) (models.Book, error) {
	var book models.Book
	id, _ := strconv.Atoi(ctx.Param("id"))

	for i, v := range books {
		if id == v.ID {
			book = books[i]
		}
	}

	return book, nil
}

//update book
func UpdateBook(ctx echo.Context) (models.Book, error) {
	var book models.Book
	newBook := new(models.Book)
	ctx.Bind(&newBook)
	id, _ := strconv.Atoi(ctx.Param("id"))
	for i, v := range books {
		if id == v.ID {
			book = books[i]
		}
	}

	book.Title = newBook.Title
	book.Writer = newBook.Writer
	book.Publisher = newBook.Publisher
	book.PublicationYear = newBook.PublicationYear

	return book, nil
}

//delete book
func DeleteBook(ctx echo.Context) error {
	arrBook := make([]models.Book, 0)
	id, _ := strconv.Atoi(ctx.Param("id"))
	for i, v := range books {
		if id == v.ID {
			arrBook = append(arrBook, books[:i]...)
			arrBook = append(arrBook, books[i+1:]...)
			books = arrBook
		}
	}

	return nil
}
