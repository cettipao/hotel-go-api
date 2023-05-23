package services

import (
	log "github.com/sirupsen/logrus"
	userCliente "mvc-go/clients/user"
	reservations_dto "mvc-go/dto/reservations_dto"
	"mvc-go/dto/users_dto"
	"mvc-go/model"
	"mvc-go/services/login"
	e "mvc-go/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (users_dto.UserDetailDto, e.ApiError)
	GetUsers() (users_dto.UsersDto, e.ApiError)
	InsertUser(userDto users_dto.UserDtoRegister) (users_dto.UserDetailDto, e.ApiError)
	UserLogin(userDto users_dto.UserLoginDto) (users_dto.UserLoginResponseDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (users_dto.UserDetailDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDetailDto users_dto.UserDetailDto

	if user.Id == 0 {
		return userDetailDto, e.NewBadRequestApiError("users_dto not found")
	}

	userDetailDto.Name = user.Name
	userDetailDto.LastName = user.LastName
	userDetailDto.Dni = user.Dni
	userDetailDto.Email = user.Email
	userDetailDto.Admin = user.Admin

	log.Debug(len(user.Reservations))

	for _, reservation := range user.Reservations {
		var dtoReservation reservations_dto.ReservationDto

		dtoReservation.Id = reservation.Id
		dtoReservation.HotelName = reservation.Hotel.Name

		userDetailDto.ReservationsDto = append(userDetailDto.ReservationsDto, dtoReservation)
	}

	return userDetailDto, nil
}

func (s *userService) GetUsers() (users_dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto users_dto.UsersDto

	for _, user := range users {
		var userDto users_dto.UserDto
		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.Dni = user.Dni
		userDto.Email = user.Email

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto users_dto.UserDtoRegister) (users_dto.UserDetailDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.Dni = userDto.Dni
	user.Email = userDto.Email
	hash, _ := login.HashPassword(userDto.Password)
	user.Password = hash
	user.Admin = 0

	user = userCliente.InsertUser(user)

	var userDetailDto users_dto.UserDetailDto
	userDetailDto.Name = user.Name
	userDetailDto.LastName = user.LastName
	userDetailDto.Dni = user.Dni
	userDetailDto.Email = user.Email

	return userDetailDto, nil
}

func (s *userService) UserLogin(userDto users_dto.UserLoginDto) (users_dto.UserLoginResponseDto, e.ApiError) {
	var loginResponse users_dto.UserLoginResponseDto
	user := userCliente.GetUserByEmail(userDto.Email)
	if user.Id == 0 {
		return loginResponse, e.NewBadRequestApiError("User not found")
	}
	if !login.CheckPasswordHash(userDto.Password, user.Password) {
		//Retornar Api error de contraseña incorrecta
		return loginResponse, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}
	token, err := login.GenerateToken(user.Id, user.Email)
	if err != nil {
		// Retornar Api error de error al generar el token
		return loginResponse, e.NewInternalServerApiError("Error al generar el token", err)
	}
	loginResponse.Token = token
	return loginResponse, nil
}
