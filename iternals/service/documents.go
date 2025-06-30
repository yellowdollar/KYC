package service

import (
	"KYC/iternals/models"
	"KYC/iternals/repository"
)

func UpdateProfileDocuments(profileID int, d models.Documents) (*models.Documents, error) {
	result, err := repository.UpdateProfileDocuments(profileID, d)
	if err != nil {
		return nil, err
	}

	return result, err
}
