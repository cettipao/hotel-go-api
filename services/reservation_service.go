package services

import (
	reservationClient "mvc-go/clients/reservation"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type reservationService struct{}

type reservationServiceInterface interface {
	GetReservationById(id int) (dto.ReservationDetailDto, e.ApiError)
	GetReservations() (dto.ReservationsDetailDto, e.ApiError)
	//GetReservationsByUser() (dto.ReservationsDto, e.ApiError)
	//GetReservationsByHotel() (dto.ReservationsDto, e.ApiError)
	InsertReservation(userDto dto.ReservationCreateDto) (dto.ReservationDetailDto, e.ApiError)
}

var (
	ReservationService reservationServiceInterface
)

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) GetReservationById(id int) (dto.ReservationDetailDto, e.ApiError) {

	var reservation model.Reservation = reservationClient.GetReservationById(id)
	var reservationDetailDto dto.ReservationDetailDto

	if reservation.Id == 0 {
		return reservationDetailDto, e.NewBadRequestApiError("reservation not found")
	}

	reservationDetailDto.Id = reservation.Id
	reservationDetailDto.UserName = reservation.User.Name
	reservationDetailDto.UserLastName = reservation.User.LastName
	reservationDetailDto.UserDni = reservation.User.Dni
	reservationDetailDto.UserEmail = reservation.User.Email
	reservationDetailDto.HotelName = reservation.Hotel.Name
	reservationDetailDto.HotelDescription = reservation.Hotel.Description
	/*
		for _, reservation := range user.Reservations {
			var dtoReservation dto.ReservationDto

			dtoReservation.Id = reservation.Id
			dtoReservation.HotelName = reservation.Name

			userDetailDto.ReservationsDto = append(userDetailDto.ReservationsDto, dtoReservation)
		}*/

	return reservationDetailDto, nil
}

func (s *reservationService) GetReservations() (dto.ReservationsDetailDto, e.ApiError) {

	//var reservations model.Reservations = reservationClient.GetReservations()
	var reservationsDetailDto dto.ReservationsDetailDto
	/*
		for _, reservation := range reservations {
			var reservationDetailDto dto.ReservationDetailDto
			reservationDetailDto.Id = reservation.Id
			reservationDetailDto.UserName = reservation.User.Name
			reservationDetailDto.UserLastName = reservation.User.LastName
			reservationDetailDto.UserDni = reservation.User.Dni
			reservationDetailDto.UserEmail = reservation.User.Email
			reservationDetailDto.HotelName = reservation.Hotel.Name
			reservationDetailDto.HotelDescription = reservation.Hotel.Description

			usersDto = append(usersDto, userDto)
		}*/

	return reservationsDetailDto, nil
}

func (s *reservationService) InsertReservation(reservationDto dto.ReservationCreateDto) (dto.ReservationDetailDto, e.ApiError) {

	var reservation model.Reservation

	reservation.HotelId = reservationDto.HotelId
	reservation.UserId = reservationDto.UserId

	reservation = reservationClient.InsertReservation(reservation)

	var reservationDetailDto dto.ReservationDetailDto
	reservationDetailDto, _ = s.GetReservationById(reservation.Id)

	return reservationDetailDto, nil
}
