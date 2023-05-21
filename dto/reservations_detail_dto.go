package dto

type ReservationDetailDto struct {
	Id          int    `json:"reservation_id"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`

	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	UserDni      string `json:"user_dni"`
	UserEmail    string `json:"user_email"`

	HotelName        string `json:"name"`
	HotelDescription string `json:"description"`
}

type ReservationsDetailDto []ReservationsDetailDto
