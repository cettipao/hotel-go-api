package services

import (
	log "github.com/sirupsen/logrus"
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
	DeleteReservationById(id int) e.ApiError
	GetReservations() (reservations_dto.ReservationsDetailDto, e.ApiError)
	InsertReservation(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.ReservationDetailDto, e.ApiError)
	RoomsAvailable(initialDate string, finalDate string) (reservations_dto.RoomsResponse, e.ApiError)
	RoomsAvailableHotel(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.RoomsAvailable, e.ApiError)
	GetFilteredReservations(hotelID int, userID int, startDate string, endDate string) (reservations_dto.ReservationsDetailDto, error)
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

	return reservationDetailDto, nil
}

func (s *reservationService) DeleteReservationById(id int) e.ApiError {

	err := reservationClient.DeleteReservationById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *reservationService) GetReservations() (reservations_dto.ReservationsDetailDto, e.ApiError) {
	var reservations = reservationClient.GetReservations()
	var reservationsDetailDto reservations_dto.ReservationsDetailDto
	reservationsDetailDto.Reservations = []reservations_dto.ReservationDetailDto{}

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

		reservationsDetailDto.Reservations = append(reservationsDetailDto.Reservations, reservationDetailDto)
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

func (s *reservationService) RoomsAvailable(initialDate string, finalDate string) (reservations_dto.RoomsResponse, e.ApiError) {
	var response reservations_dto.RoomsResponse

	// Obtener todos los hoteles
	hotels := hotelClient.GetHotels()

	// Iterar sobre cada hotel y obtener las habitaciones disponibles
	for _, hotel := range hotels {
		hotelId := hotel.Id

		// Obtener las reservas para el hotel y las fechas dadas
		initalDate, _ := time.Parse(layout, initialDate)
		finalDate, _ := time.Parse(layout, finalDate)
		reservations := reservationClient.GetReservationsByHotelAndDates(hotelId, initalDate, finalDate)

		// Calcular las habitaciones disponibles
		roomsAvailable := hotel.RoomsAvailable - reservations

		// Agregar la información del hotel y las habitaciones disponibles a la respuesta
		roomInfo := reservations_dto.RoomInfo{
			Name:           hotel.Name,
			RoomsAvailable: roomsAvailable,
			Id:             hotel.Id,
		}
		response.Rooms = append(response.Rooms, roomInfo)
	}

	return response, nil
}

func (s *reservationService) RoomsAvailableHotel(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.RoomsAvailable, e.ApiError) {

	hotelId := reservationDto.HotelId
	initalDate, _ := time.Parse(layout, reservationDto.InitialDate)
	finalDate, _ := time.Parse(layout, reservationDto.FinalDate)
	var reservations = reservationClient.GetReservationsByHotelAndDates(hotelId, initalDate, finalDate)

	var roomsAvailable reservations_dto.RoomsAvailable
	hotel_rooms := hotelClient.GetHotelById(hotelId).RoomsAvailable
	roomsAvailable.Rooms = hotel_rooms - reservations

	return roomsAvailable, nil
}

// En el archivo service/reservation_service.go

func (s *reservationService) GetFilteredReservations(hotelID int, userID int, startDate string, endDate string) (reservations_dto.ReservationsDetailDto, error) {
	// Realiza la lógica para filtrar las reservas según los parámetros proporcionados
	newLayout := "2006-01-02"
	// Crea una variable para almacenar las reservas filtradas
	var filteredReservations = make([]reservations_dto.ReservationDetailDto, 0)

	// Obtén todas las reservas existentes desde tu fuente de datos (por ejemplo, una base de datos)
	allReservations := reservationClient.GetReservations()

	// Aplica los filtros si se proporcionan
	for _, reservation := range allReservations {
		// Verifica si el ID del hotel coincide con el parámetro proporcionado
		if hotelID != 0 && reservation.HotelId != hotelID {
			continue
		}

		// Verifica si el ID del usuario coincide con el parámetro proporcionado
		if userID != 0 && reservation.UserId != userID {
			continue
		}

		// Verifica si la fecha de inicio coincide con el parámetro proporcionado
		if startDate != "" {
			startTime, err := time.Parse(newLayout, startDate)
			if err != nil {
				log.Error(err)
				return reservations_dto.ReservationsDetailDto{}, err
			}

			if reservation.InitialDate.Before(startTime) {
				continue
			}
		}

		// Verifica si la fecha de finalización coincide con el parámetro proporcionado
		if endDate != "" {
			log.Error(endDate)
			endTime, err := time.Parse(newLayout, endDate)
			if err != nil {
				log.Error(err)
				return reservations_dto.ReservationsDetailDto{}, err
			}

			if reservation.FinalDate.After(endTime) {
				continue
			}
		}

		// Si la reserva pasa todos los filtros, agrégala a las reservas filtradas
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
		filteredReservations = append(filteredReservations, reservationDetailDto)
	}

	// Crea una instancia de ReservationsDetailDto y asigna las reservas filtradas
	reservationsDto := reservations_dto.ReservationsDetailDto{
		Reservations: filteredReservations,
	}

	return reservationsDto, nil
}
