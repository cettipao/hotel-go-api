package dto

type HotelDetailDto struct {
	Name           string `json:"name"`
	RoomsAvailable int    `json:"rooms_available"`
	Description    string `json:"description"`

	ReservationsDto ReservationsDto `json:"reservas,omitempty"`
}

type HotelsDetailDto []HotelDetailDto

/*
package model

type Hotel struct {
	Id             int    `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(50);not null"`
	RoomsAvailable int
	Description    string `gorm:"type:varchar(250);not null"`
}

type Hotels []User


*/
