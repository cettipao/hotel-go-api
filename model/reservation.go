package model

type Reservation struct {
	Id int `gorm:"primaryKey"`

	User   User `gorm:"foreignkey:UserId"`
	UserId int

	Hotel   Hotel `gorm:"foreignkey:HotelId"`
	HotelId int

	Name           string `gorm:"type:varchar(50);not null"`
	RoomsAvailable int
	Description    string `gorm:"type:varchar(250);not null"`
}

type Reservations []User
