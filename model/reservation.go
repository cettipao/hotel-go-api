package model

import "time"

type Reservation struct {
	Id          int       `gorm:"primaryKey"`
	InitialDate time.Time `gorm:"column:initial_date;not null"`
	FinalDate   time.Time `gorm:"column:final_date;not null"`

	User   User `gorm:"foreignkey:UserId"`
	UserId int

	Hotel   Hotel `gorm:"foreignkey:HotelId"`
	HotelId int
}

type Reservations []Reservation
