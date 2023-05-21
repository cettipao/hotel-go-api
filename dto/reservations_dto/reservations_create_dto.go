package reservations_dto

type ReservationCreateDto struct {
	UserId      int    `json:"user_id"`
	HotelId     int    `json:"hotel_id"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
}
