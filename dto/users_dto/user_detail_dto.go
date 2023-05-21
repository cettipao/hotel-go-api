package users_dto

import "mvc-go/dto/reservations_dto"

type UserDetailDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Dni      string `json:"dni"`
	Email    string `json:"email"`
	Admin    int    `json:"admin"`

	ReservationsDto reservations_dto.ReservationsDto `json:"reservas,omitempty"`
}
