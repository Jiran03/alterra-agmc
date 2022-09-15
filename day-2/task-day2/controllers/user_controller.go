package controllers

import (
	"net/http"

	"github.com/Jiran03/agmc/task/day2/lib/database"
	"github.com/Jiran03/agmc/task/day2/models"
	"github.com/labstack/echo/v4"
)

//get all users
func GetAllUserController(c echo.Context) error {
	users, err := database.GetAllUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userValue := make([]models.UserValue, 0)
	for i := range users {
		userValue = append(userValue, models.UserValue{
			ID:        int(users[i].ID),
			Name:      users[i].Name,
			Email:     users[i].Email,
			CreatedAt: users[i].CreatedAt,
			UpdatedAt: users[i].UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all users",
		"users":   userValue,
	})
}

//create new user
func CreateUserController(c echo.Context) error {
	user, err := database.CreateUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userValue := models.UserValue{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new user",
		"user":    userValue,
	})
}

//get user by id
func GetUserByIdController(c echo.Context) error {
	user, err := database.GetUserByID(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data tidak ditemukan",
		})
	}

	userValue := models.UserValue{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succes get user by id",
		"user":    userValue,
	})
}

//update user by id
func UpdateUserController(c echo.Context) error {
	user, err := database.UpdateUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userValue := models.UserValue{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succes update user by id",
		"user":    userValue,
	})
}

//delete user by id
func DeleteUserController(c echo.Context) error {
	err := database.DeleteUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "user deleted",
	})
}
