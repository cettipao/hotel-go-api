package hotels_dto

type HotelDto struct {
	Name           string   `json:"name"`
	RoomsAvailable int      `json:"rooms_available"`
	Description    string   `json:"description"`
	Id             int      `json:"id"`
	Amenities      []string `json:"amenities"`
}

//type HotelsDto []HotelDto

type HotelsDto struct {
	Hotels []HotelDto `json:"hotels"`
}

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
