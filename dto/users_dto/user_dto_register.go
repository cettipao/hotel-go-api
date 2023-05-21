package users_dto

type UserDtoRegister struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Dni      string `json:"dni"`
	Password string `json:"password"`
}
