package repository

import (
	"KYC/iternals/db"
	"KYC/iternals/models"
	"fmt"
	"strconv"
)

func CheckRoleByUserID(userID int) (string, error) {

	dbcon, err := db.GetDBConn()
	if err != nil {
		return "", err
	}

	var u models.User

	result := dbcon.Where("id = ?", userID).Take(&u)
	if result.Error != nil {
		return "", result.Error
	}

	return u.Role, nil
}

func GetUsers(userID string, filterStatus string) ([]models.User, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	var users []models.User

	query := dbcon.Preload("Profile").Where("role <> ?", "admin")

	if userID != "" {
		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			return nil, err
		}

		query = query.Where("id = ?", userIDint)
	} else if filterStatus != "" {
		query = query.Where("is_identified = ?", filterStatus)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func ConfirmUserIdentifyStatus(userID int) (*models.User, error) {
	dbcon, err := db.GetDBConn()
	if err != nil {
		return nil, err
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.IsIdentified == "pending" {
		user.IsIdentified = "identified"
	} else {
		return nil, fmt.Errorf("cannot confirm user: status is: %s", user.IsIdentified)
	}

	result := dbcon.Save(&user)
	if result.Error != nil {
		return nil, err
	}

	return user, nil
}
