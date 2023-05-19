package model

type Reservation struct {
	Id int `gorm:"primaryKey"`
	//Inicio
	//Fin

	User   User `gorm:"foreignkey:UserId"`
	UserId int

	Hotel   Hotel `gorm:"foreignkey:HotelId"`
	HotelId int
}

type Reservations []User
