package model

type Hotel struct {
	Id             int    `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(50);not null"`
	RoomsAvailable int
	Description    string `gorm:"type:varchar(250);not null"`
}

type Hotels []Hotel