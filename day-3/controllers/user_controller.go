package controllers

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/agmc/task/day3/lib/database"
	"github.com/Jiran03/agmc/task/day3/middlewares"
	"github.com/Jiran03/agmc/task/day3/models"
	"github.com/labstack/echo/v4"
)

//get all users
func GetAllUserController(ctx echo.Context) error {
	users, err := database.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userValue := make([]models.UserResponse, 0)
	for i := range users {
		userValue = append(userValue, models.UserResponse{
			ID:        int(users[i].ID),
			Name:      users[i].Name,
			Email:     users[i].Email,
			CreatedAt: users[i].CreatedAt,
			UpdatedAt: users[i].UpdatedAt,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all users",
		"users":   userValue,
	})
}

//create new user
func CreateUserController(ctx echo.Context) error {
	user := models.User{}
	ctx.Bind(&user)
	user, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userValue := models.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create new user",
		"user":    userValue,
	})
}

//get user by id
func GetUserByIdController(ctx echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := database.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user.ID == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data tidak ditemukan",
		})
	}

	userValue := models.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succes get user by id",
		"user":    userValue,
	})
}

//update user by id
func UpdateUserController(ctx echo.Context) error {
	newUser := new(models.User)
	ctx.Bind(&newUser)
	id, _ := strconv.Atoi(ctx.Param("id"))
	userData, err := database.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID, userEmail := middlewares.ExtractToken(ctx)
	if userID != id && userEmail != userData.Email {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "anda tidak bisa mengakses data ini"})
	}

	user, err := database.UpdateUser(newUser, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userValue := models.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succes update user by id",
		"user":    userValue,
	})
}

//delete user by id
func DeleteUserController(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userData, err := database.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID, userEmail := middlewares.ExtractToken(ctx)
	if userID != id && userEmail != userData.Email {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "anda tidak bisa mengakses data ini"})
	}

	err = database.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "user deleted",
	})
}

func LoginUserController(ctx echo.Context) error {
	var err error
	user := models.User{}
	ctx.Bind(&user)
	userValid, err := database.LoginUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var token string
	token, err = middlewares.CreateToken(int(userValid.ID), userValid.Email)
	userResponse := models.UserResponse{
		ID:        int(userValid.ID),
		Name:      userValid.Name,
		Email:     userValid.Email,
		CreatedAt: userValid.CreatedAt,
		UpdatedAt: userValid.UpdatedAt}
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success login",
		"user":    userResponse,
		"token":   token,
	})
}
