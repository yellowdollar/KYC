package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
}
