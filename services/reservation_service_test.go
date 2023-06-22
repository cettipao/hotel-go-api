package services

import (
	"github.com/stretchr/testify/assert"
	hotelClient "mvc-go/clients/hotel"
	reservationClient "mvc-go/clients/reservation"
	"mvc-go/dto/reservations_dto"
	"mvc-go/model"
	"testing"
	"time"
)

type TestReservations struct {
}

func (testReservations TestReservations) GetReservationsByHotelAndDate(idHotel int, date time.Time) int {
	return 4
}

func (testReservations TestReservations) InsertReservation(reservation model.Reservation) model.Reservation {
	reservationCreated := model.Reservation{
		Id:          1,
		InitialDate: reservation.InitialDate,
		FinalDate:   reservation.FinalDate,
		User:        reservation.User,
		UserId:      reservation.UserId,
		Hotel:       reservation.Hotel,
		HotelId:     reservation.HotelId,
	}
	return reservationCreated
}

func (testReservations TestReservations) GetReservationById(id int) model.Reservation {
	parsedTime, _ := time.Parse(layout, "01/01/2023")
	reservation := model.Reservation{
		Id:          id,
		InitialDate: parsedTime,
		FinalDate:   parsedTime,
		User:        model.User{},
		UserId:      0,
		Hotel:       model.Hotel{},
		HotelId:     0,
	}

	return reservation
}

func TestInsertReservation(t *testing.T) {
	//prepare
	reservation := reservations_dto.ReservationCreateDto{
		UserId:      1,
		HotelId:     123,
		InitialDate: "01/01/2023",
		FinalDate:   "01/01/2023",
	}
	reservationClient.MyClient = TestReservations{}

	//act
	reservationCreated, _ := ReservationService.InsertReservation(reservation)

	//assert
	assert.Equal(t, 1, reservationCreated.Id)
	assert.Equal(t, reservation.InitialDate, reservationCreated.InitialDate)

}

func TestRoomsAvailable(t *testing.T) {
	//prepare
	initialDate := "01/01/2023"
	finalDate := "15/01/2023"
	reservationClient.MyClient = TestReservations{}
	hotelClient.MyClient = TestHotels{}

	//act
	roomsAvailable, _ := ReservationService.RoomsAvailable(initialDate, finalDate)

	//assert
	assert.Equal(t, 6, roomsAvailable.Rooms[0].RoomsAvailable)
	assert.Equal(t, 6, roomsAvailable.Rooms[1].RoomsAvailable)
}

func (testReservations TestReservations) GetReservationsByUser(idUser int) model.Reservations {
	return nil
}
func (testReservations TestReservations) GetReservationsByHotel(idHotel int) model.Reservations {
	return nil
}
func (testReservations TestReservations) GetReservationsByHotelAndDates(idHotel int, initialDate time.Time, finalDate time.Time) int {
	return 0
}
