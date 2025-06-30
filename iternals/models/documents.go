package models

type Documents struct {
	ID        uint    `json:"id"`
	Front     string  `json:"front"`
	Back      string  `json:"back"`
	Selfie    string  `json:"selfie"`
	ProfileID int     `json:"account_id"`
	Profile   Profile `json:"-" gorm:"foreignKey:ProfileID"`
}
