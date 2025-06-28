package service

import (
	"KYC/iternals/models"
	"KYC/iternals/repository"
	"KYC/utils"
)

func CreateAdmin(a models.Admin) (*models.Admin, error) {

	a.Password = utils.GenerateHash(a.Password)

	admin, err := repository.CreateAdmin(a)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func GetAdminByLogin(login string) (*models.Admin, error) {

	admin, err := repository.GetAdminByLogin(login)
	if err != nil {
		return nil, err
	}

	return admin, nil

}
