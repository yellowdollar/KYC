package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
)

func InsertProfileData(userID int, a *models.Profile) (*models.Profile, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	userProfile, err := GetUserProfileByID(userID)
	if err != nil {
		return nil, err
	}

	userProfile.Name = a.Name
	userProfile.Surname = a.Surname
	userProfile.Gender = a.Gender
	userProfile.Age = a.Age

	result := dbcon.Save(&userProfile)
	if result.Error != nil {
		return nil, result.Error
	}

	return userProfile, err
}

func CreateUserProfile(userID int) error {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return err
	}

	var a = models.Profile{
		UserID: uint(userID),
	}

	result := dbcon.Create(&a)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserProfileByID(userID int) (*models.Profile, error) {

	var userProfile models.Profile

	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	result := dbcon.Where("user_id = ?", userID).First(&userProfile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userProfile, nil
}
