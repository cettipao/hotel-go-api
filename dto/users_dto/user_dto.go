package users_dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Dni      string `json:"dni"`
}

type UsersDto struct {
	Users []UserDto `json:"users"`
}
