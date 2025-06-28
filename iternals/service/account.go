package service

import (
	"KYC/iternals/models"
	"KYC/iternals/repository"
)

func UpdateAccountInfo(userID int, a models.Account) (*models.Account, error) {
	result, err := repository.InsertUserData(userID, &a)
	if err != nil {
		return nil, err
	}

	// creating documents null row with accountID
	_, err = repository.CreateAccountDocuments(int(result.ID))
	if err != nil {
		return nil, err
	}

	return result, nil
}
