package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Male    string `json:"male"`
	Age     int    `json:"age"`
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
}
