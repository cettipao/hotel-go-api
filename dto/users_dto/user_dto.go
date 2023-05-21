package users_dto

type UserDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Dni      string `json:"dni"`
	Id       int    `json:"id"`
}

type UsersDto []UserDto
