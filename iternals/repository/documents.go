package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
)

func CreateAccountDocuments(accountID int) (*models.Documents, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var d = models.Documents{
		AccountID: accountID,
	}

	result := dbcon.Create(&d)
	if result.Error != nil {
		return nil, err
	}

	return &d, nil
}

func GetAccountDocumentsByAccountId(accountID int) (*models.Documents, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var d models.Documents

	result := dbcon.Where("account_id = ?", accountID).First(&d)
	if result.Error != nil {
		return nil, err
	}

	return &d, nil
}

func InsertAccountDocuments(accountID int, d models.Documents) (*models.Documents, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	accountDocs, err := GetAccountDocumentsByAccountId(accountID)
	if err != nil {
		return nil, err
	}

	accountDocs.Front = d.Front
	accountDocs.Back = d.Back
	accountDocs.Selfie = d.Selfie

	result := dbcon.Save(&accountDocs)
	if result.Error != nil {
		return nil, result.Error
	}

	return accountDocs, nil
}
