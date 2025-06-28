package db

import (
	"KYC/iternals/configs"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnDB() (*gorm.DB, error) {
	cfg := configs.AppSettings.PostgresParams

	dsn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		cfg.Host,
		cfg.Port,
		cfg.User,
		os.Getenv("DB_PASSWORD"),
		cfg.DatabaseName)

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error while connecting DB: %w", err)
	}

	return db, nil
}

func CloseDBConn() error {
	sqlDB, err := db.DB()

	if err != nil {
		return fmt.Errorf("failed to get raw Connection to DB: %w", err)
	}

	return sqlDB.Close()
}

func GetDBConn() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("there is no database connections initialized")
	}

	return db, nil
}
