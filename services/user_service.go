package services

import (
	userCliente "mvc-go/clients/user"

	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDetailDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDtoRegister) (dto.UserDetailDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDetailDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDetailDto dto.UserDetailDto

	if user.Id == 0 {
		return userDetailDto, e.NewBadRequestApiError("user not found")
	}

	userDetailDto.Name = user.Name
	userDetailDto.LastName = user.LastName
	userDetailDto.Dni = user.Dni
	userDetailDto.Email = user.Email
	/*
		for _, reservation := range user.Reservations {
			var dtoReservation dto.ReservationDto

			dtoReservation.Id = reservation.Id
			dtoReservation.HotelName = reservation.Name

			userDetailDto.ReservationsDto = append(userDetailDto.ReservationsDto, dtoReservation)
		}*/

	return userDetailDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.Dni = user.Dni
		userDto.Email = user.Email

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDtoRegister) (dto.UserDetailDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.Dni = userDto.Dni
	user.Email = userDto.Email
	user.Password = userDto.Password
	user.Admin = 0

	user = userCliente.InsertUser(user)

	var userDetailDto dto.UserDetailDto
	userDetailDto.Name = user.Name
	userDetailDto.LastName = user.LastName
	userDetailDto.Dni = user.Dni
	userDetailDto.Email = user.Email

	return userDetailDto, nil
}
