package dto

type ReservationDto struct {
	Id        int    `json:"reservation_id"`
	HotelName string `json:"name"`
	//Inicio
	//Fin
}

type ReservationsDto []ReservationsDto
