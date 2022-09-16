package database

import (
	"github.com/Jiran03/agmc/task/day3/config"
	"github.com/Jiran03/agmc/task/day3/models"
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
func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

//get user by id
func GetUserByID(id int) (models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

//update user
func UpdateUser(newUser *models.User, id int) (models.User, error) {
	var user models.User
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
func DeleteUser(id int) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

//login user
func LoginUser(user models.User) (*models.User, error) {
	var err error
	var userValid models.User
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&userValid).Error; err != nil {
		return nil, err
	}

	return &userValid, nil
}
