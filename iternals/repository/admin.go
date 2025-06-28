package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
)

func CreateAdmin(a models.Admin) (*models.Admin, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	result := dbcon.Create(&a)
	if result.Error != nil {
		return nil, err
	}

	return &a, err
}

func GetAdminByLogin(adminLogin string) (*models.Admin, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var admin = models.Admin{}

	result := dbcon.Where("login = ?", adminLogin).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}
