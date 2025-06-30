package service

import (
	"KYC/iternals/models"
	"KYC/iternals/repository"
)

func CheckRole(userID int) (bool, error) {
	userRole, err := repository.CheckRoleByUserID(userID)
	if err != nil {
		return false, err
	}

	if userRole != "admin" {
		return false, nil
	}

	return true, nil
}

func GetUsers(userID string, filterStatus string) ([]models.User, error) {
	users, err := repository.GetUsers(userID, filterStatus)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func ConfirmUserIdentifyStatus(userID int) (*models.User, error) {
	user, err := repository.ConfirmUserIdentifyStatus(userID)
	if err != nil {
		return nil, err
	}

	return user, err
}
