package controllers

import (
	"net/http"

	"github.com/Jiran03/agmc/task/day4/lib/database"
	"github.com/labstack/echo/v4"
)

//get all books
func GetAllBookController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all Books",
		"books":   books,
	})
}

//create new book
func CreateBookController(c echo.Context) error {
	book, err := database.CreateBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new book",
		"book":    book,
	})
}

//get book by id
func GetBookByIdController(c echo.Context) error {
	book, err := database.GetBookByID(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if book.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data tidak ditemukan",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succes get book by id",
		"book":    book,
	})
}

//update book by id
func UpdateBookController(c echo.Context) error {
	book, err := database.UpdateBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succes update book by id",
		"book":    book,
	})
}

//delete book by id
func DeleteBookController(c echo.Context) error {
	err := database.DeleteBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "book deleted",
	})
}
