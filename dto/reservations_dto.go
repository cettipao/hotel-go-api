package dto

type ReservationDto struct {
	Id int `json:"reservation_id"`

	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	UserDni      string `json:"user_dni"`
	UserEmail    string `json:"user_email"`

	HotelName        string `json:"name"`
	HotelDescription string `json:"description"`
}

type ReservationsDto []ReservationsDto
