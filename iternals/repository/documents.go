package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
)

func CreateProfileDocuments(profileID int) (*models.Documents, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var d = models.Documents{
		ProfileID: profileID,
	}

	result := dbcon.Create(&d)
	if result.Error != nil {
		return nil, err
	}

	return &d, nil
}

func GetDocumentsByProfileID(profileID int) (*models.Documents, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var d models.Documents

	result := dbcon.Where("profile_id = ?", profileID).Take(&d)
	if result.Error != nil {
		return nil, result.Error
	}

	return &d, nil
}

func UpdateProfileDocuments(profileID int, d models.Documents) (*models.Documents, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	docs, err := GetDocumentsByProfileID(profileID)
	if err != nil {
		return nil, err
	}

	docs.Front = d.Front
	docs.Back = d.Back
	docs.Selfie = d.Selfie

	result := dbcon.Save(&docs)
	if result.Error != nil {
		return nil, result.Error
	}

	return docs, nil
}
