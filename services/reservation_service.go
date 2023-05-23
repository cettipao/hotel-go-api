package services

import (
	hotelClient "mvc-go/clients/hotel"
	reservationClient "mvc-go/clients/reservation"
	"mvc-go/dto/reservations_dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type reservationService struct{}

type reservationServiceInterface interface {
	GetReservationById(id int) (reservations_dto.ReservationDetailDto, e.ApiError)
	GetReservations() (reservations_dto.ReservationsDetailDto, e.ApiError)
	GetReservationsByUser(id int) (reservations_dto.ReservationsDetailDto, e.ApiError)
	//GetReservationsByHotel() (dto.ReservationsDto, e.ApiError)
	InsertReservation(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.ReservationDetailDto, e.ApiError)
	RoomsAvailable(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.RoomsAvailable, e.ApiError)
}

var (
	ReservationService reservationServiceInterface
	layout             = "02/01/2006"
)

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) GetReservationById(id int) (reservations_dto.ReservationDetailDto, e.ApiError) {

	var reservation model.Reservation = reservationClient.GetReservationById(id)
	var reservationDetailDto reservations_dto.ReservationDetailDto

	if reservation.Id == 0 {
		return reservationDetailDto, e.NewBadRequestApiError("reservations_dto not found")
	}

	reservationDetailDto.Id = reservation.Id
	reservationDetailDto.UserName = reservation.User.Name
	reservationDetailDto.InitialDate = reservation.InitialDate.Format(layout)
	reservationDetailDto.FinalDate = reservation.FinalDate.Format(layout)
	reservationDetailDto.UserLastName = reservation.User.LastName
	reservationDetailDto.UserDni = reservation.User.Dni
	reservationDetailDto.UserEmail = reservation.User.Email
	reservationDetailDto.HotelName = reservation.Hotel.Name
	reservationDetailDto.HotelDescription = reservation.Hotel.Description
	/*
		for _, reservations_dto := range users_dto.Reservations {
			var dtoReservation dto.ReservationDto

			dtoReservation.Id = reservations_dto.Id
			dtoReservation.HotelName = reservations_dto.Name

			userDetailDto.ReservationsDto = append(userDetailDto.ReservationsDto, dtoReservation)
		}*/

	return reservationDetailDto, nil
}

func (s *reservationService) GetReservations() (reservations_dto.ReservationsDetailDto, e.ApiError) {

	var reservations = reservationClient.GetReservations()
	var reservationsDetailDto reservations_dto.ReservationsDetailDto

	for _, reservation := range reservations {
		var reservationDetailDto reservations_dto.ReservationDetailDto
		reservationDetailDto.Id = reservation.Id
		reservationDetailDto.UserName = reservation.User.Name
		reservationDetailDto.UserLastName = reservation.User.LastName
		reservationDetailDto.UserDni = reservation.User.Dni
		reservationDetailDto.UserEmail = reservation.User.Email
		reservationDetailDto.HotelName = reservation.Hotel.Name
		reservationDetailDto.HotelDescription = reservation.Hotel.Description
		reservationDetailDto.InitialDate = reservation.InitialDate.Format(layout)
		reservationDetailDto.FinalDate = reservation.FinalDate.Format(layout)
		reservationDetailDto.FinalDate = reservation.FinalDate.Format(layout)

		reservationsDetailDto = append(reservationsDetailDto, reservationDetailDto)
	}

	return reservationsDetailDto, nil
}

func (s *reservationService) GetReservationsByUser(id int) (reservations_dto.ReservationsDetailDto, e.ApiError) {

	var reservations = reservationClient.GetReservationsByUser(id)
	var reservationsDetailDto reservations_dto.ReservationsDetailDto

	for _, reservation := range reservations {
		var reservationDetailDto reservations_dto.ReservationDetailDto
		reservationDetailDto.Id = reservation.Id
		reservationDetailDto.UserName = reservation.User.Name
		reservationDetailDto.UserLastName = reservation.User.LastName
		reservationDetailDto.UserDni = reservation.User.Dni
		reservationDetailDto.UserEmail = reservation.User.Email
		reservationDetailDto.HotelName = reservation.Hotel.Name
		reservationDetailDto.HotelDescription = reservation.Hotel.Description
		reservationDetailDto.InitialDate = reservation.InitialDate.Format(layout)
		reservationDetailDto.FinalDate = reservation.FinalDate.Format(layout)
		reservationDetailDto.FinalDate = reservation.FinalDate.Format(layout)

		reservationsDetailDto = append(reservationsDetailDto, reservationDetailDto)
	}

	return reservationsDetailDto, nil
}

func (s *reservationService) InsertReservation(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.ReservationDetailDto, e.ApiError) {

	var reservation model.Reservation
	var reservationDetailDto reservations_dto.ReservationDetailDto

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

func (s *reservationService) RoomsAvailable(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.RoomsAvailable, e.ApiError) {

	hotelId := reservationDto.HotelId
	initalDate, _ := time.Parse(layout, reservationDto.InitialDate)
	finalDate, _ := time.Parse(layout, reservationDto.FinalDate)
	var reservations = reservationClient.GetReservationsByHotelAndDates(hotelId, initalDate, finalDate)

	var roomsAvailable reservations_dto.RoomsAvailable
	hotel_rooms := hotelClient.GetHotelById(hotelId).RoomsAvailable
	roomsAvailable.Rooms = hotel_rooms - reservations

	return roomsAvailable, nil
}
