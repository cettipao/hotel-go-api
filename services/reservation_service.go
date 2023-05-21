package services

import (
	hotelClient "mvc-go/clients/hotel"
	reservationClient "mvc-go/clients/reservation"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type reservationService struct{}

type reservationServiceInterface interface {
	GetReservationById(id int) (dto.ReservationDetailDto, e.ApiError)
	GetReservations() (dto.ReservationsDetailDto, e.ApiError)
	//GetReservationsByUser() (dto.ReservationsDto, e.ApiError)
	//GetReservationsByHotel() (dto.ReservationsDto, e.ApiError)
	InsertReservation(reservationDto dto.ReservationCreateDto) (dto.ReservationDetailDto, e.ApiError)
	RoomsAvailable(reservationDto dto.ReservationCreateDto) (dto.RoomsAvailable, e.ApiError)
}

var (
	ReservationService reservationServiceInterface
	layout             = "02/01/2006"
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
	reservationDetailDto.InitialDate = reservation.InitialDate.Format("02/01/2006")
	reservationDetailDto.FinalDate = reservation.FinalDate.Format("02/01/2006")
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
	var reservationDetailDto dto.ReservationDetailDto

	reservation.HotelId = reservationDto.HotelId
	reservation.UserId = reservationDto.UserId
	parsedTime, _ := time.Parse(layout, reservationDto.InitialDate)
	reservation.InitialDate = parsedTime
	parsedTime, _ = time.Parse(layout, reservationDto.FinalDate)
	reservation.FinalDate = parsedTime

	reservation = reservationClient.InsertReservation(reservation)

	reservationDetailDto, _ = s.GetReservationById(reservation.Id)

	return reservationDetailDto, nil
}

func (s *reservationService) RoomsAvailable(reservationDto dto.ReservationCreateDto) (dto.RoomsAvailable, e.ApiError) {

	hotelId := reservationDto.HotelId
	initalDate, _ := time.Parse(layout, reservationDto.InitialDate)
	finalDate, _ := time.Parse(layout, reservationDto.FinalDate)
	var reservations = reservationClient.GetReservationsByHotelAndDates(hotelId, initalDate, finalDate)

	var roomsAvailable dto.RoomsAvailable
	hotel_rooms := hotelClient.GetHotelById(hotelId).RoomsAvailable
	roomsAvailable.Rooms = hotel_rooms - reservations

	return roomsAvailable, nil
}
