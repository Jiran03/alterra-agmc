package database

import (
	"strconv"

	"github.com/Jiran03/agmc/task/day2/config"
	"github.com/Jiran03/agmc/task/day2/models"
	"github.com/labstack/echo/v4"
)

//get all user from DB
func GetAllUser() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

//create new user
func CreateUser(ctx echo.Context) (models.User, error) {
	user := models.User{}
	ctx.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

//get user by id
func GetUserByID(ctx echo.Context) (models.User, error) {
	var user models.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

//update user
func UpdateUser(ctx echo.Context) (models.User, error) {
	var user models.User
	newUser := new(models.User)
	ctx.Bind(&newUser)
	id, _ := strconv.Atoi(ctx.Param("id"))

	config.DB.First(&user, id)
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password
	if err := config.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

//delete user
func DeleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
