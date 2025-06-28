package models

import "gorm.io/gorm"

type Documents struct {
	gorm.Model
	Front     string  `json:"front"`
	Back      string  `json:"back"`
	Selfie    string  `json:"selfie"`
	AccountID int     `json:"account_id"`
	Account   Account `gorm:"foreignKey:AccountID"`
}
