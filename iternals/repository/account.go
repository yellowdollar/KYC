package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
)

func InsertUserData(userID int, a *models.Account) (*models.Account, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	userAccount, err := GetUserAccountByUserID(userID)
	if err != nil {
		return nil, err
	}

	userAccount.Name = a.Name
	userAccount.Surname = a.Surname
	userAccount.Male = a.Male
	userAccount.Age = a.Age

	result := dbcon.Save(&userAccount)
	if result.Error != nil {
		return nil, result.Error
	}

	return userAccount, err
}

func CreateUserAccount(userID int) error {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return err
	}

	var a = models.Account{
		UserID: uint(userID),
	}

	result := dbcon.Create(&a)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserAccountByUserID(userID int) (*models.Account, error) {

	var userAccount models.Account

	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	result := dbcon.Where("user_id = ?", userID).First(&userAccount)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userAccount, nil
}
