package dto

type UserDetailDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Dni      string `json:"dni"`
	Email    string `json:"email"`

	ReservationsDto ReservationsDto `json:"reservas,omitempty"`
}
