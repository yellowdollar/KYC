package service

import (
	"KYC/iternals/models"
	"KYC/iternals/repository"
)

func UpdateProfileInfo(userID int, a models.Profile) (*models.Profile, error) {
	result, err := repository.InsertProfileData(userID, &a)
	if err != nil {
		return nil, err
	}

	// creating documents null row with accountID
	_, err = repository.CreateProfileDocuments(int(result.ID))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProfileByUserID(userID int) (*models.Profile, error) {
	result, err := repository.GetUserProfileByID(userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
