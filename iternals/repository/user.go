package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
)

func CreateUser(u models.User) (*models.User, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	result := dbcon.Create(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func GetUserByLogin(userLogin string) (*models.User, error) {
	var user models.User

	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	result := dbcon.Where("login = ?", userLogin).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
