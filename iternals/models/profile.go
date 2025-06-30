package models

type Profile struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"-" gorm:"foreignKey:UserID"`
}

type ProfileUpdate struct {
	Name    string `json:"name"`    // Profile name
	Surname string `json:"surname"` // Profile surname
	Gender  string `json:"gender"`  // Profile gender
	Age     int    `json:"age"`     // Profile age
}
