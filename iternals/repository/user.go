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

func GetUserByID(userID int) (*models.User, error) {
	var u models.User

	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	result := dbcon.Where("id = ?", userID).Take(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil

}

func GetUserVerificationStatus(userID int) (string, error) {

	u, err := GetUserByID(userID)
	if err != nil {
		return "", err
	}

	return u.IsIdentified, nil
}

func UpdateUserIdentification(userID int) (*models.User, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	u, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	u.IsIdentified = "pending"

	result := dbcon.Save(&u)
	if result.Error != nil {
		return nil, err
	}

	return u, nil
}
