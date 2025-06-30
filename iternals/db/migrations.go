package db

import "KYC/iternals/models"

func InitMigrations() error {

	// migrate tb_user
	if err := db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Documents{}); err != nil {
		return err
	}

	return nil
}
