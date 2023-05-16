package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(250);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	Dni      string `gorm:"type:varchar(30);not null;"`
	Email    string `gorm:"type:varchar(150);not null;"`
	Password string `gorm:"type:varchar(30);not null"`
	Admin    int
}

type Users []User
