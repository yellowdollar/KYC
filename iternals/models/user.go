package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login        string  `json:"login"              gorm:"uniqueIndex"`
	Password     string  `json:"-"`
	Role         string  `json:"role"               gorm:"default:'user'"`
	IsIdentified string  `json:"is_identified"      gorm:"default:'unidentified'"`
	Profile      Profile `json:"profile"            gorm:"foreignKey:UserID"`
}

type UserWithoutProfile struct {
	ID           uint   `json:"id"`
	Login        string `json:"login"`
	Role         string `json:"role"`
	IsIdentified string `json:"is_identified"`
}

type UserSignUp struct {
	Login    string `json:"login"`    // User login
	Password string `json:"password"` // User password
}

type UserSignIn struct {
	Login    string `json:"login"`    // UserSignIn login
	Password string `json:"password"` // UserSignIn password
}
