package model

type Image struct {
	ID      int    `gorm:"primaryKey"`
	Path    string `gorm:"not null"`
	HotelID int    `gorm:"not null"`
	Hotel   Hotel  `gorm:"foreignkey:HotelID"`
}

type Images []Image
