package users_dto

type UserDetailDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Dni      string `json:"dni"`
	Email    string `json:"email"`
	Admin    int    `json:"admin"`
}
