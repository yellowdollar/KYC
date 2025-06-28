package db

import "KYC/iternals/models"

func InitMigrations() error {

	// migrate tb_user
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	// migrate tb_account
	if err := db.AutoMigrate(&models.Account{}); err != nil {
		return err
	}

	// migrate tb_documents
	if err := db.AutoMigrate(&models.Documents{}); err != nil {
		return err
	}

	// migrate tb_admin
	if err := db.AutoMigrate(&models.Admin{}); err != nil {
		return err
	}

	return nil
}
