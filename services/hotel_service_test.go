package services

import (
	"github.com/stretchr/testify/assert"
	reservationClient "mvc-go/clients/reservation"
	"testing"
	"time"
)

type TestReservations struct {
}

func (testReservations TestReservations) GetReservationsByHotelAndDate(idHotel int, date time.Time) int {
	return 4
}

func TestRoomsAvailable(t *testing.T) {
	//prepare
	initialDate := "01/01/2023"
	finalDate := "15/01/2023"
	reservationClient.MyClient = TestReservations{} // Asigna una instancia de TestReservations directamente a MyClient

	//act
	roomsAvailable, _ := ReservationService.RoomsAvailable(initialDate, finalDate)

	//assert
	assert.Equal(t, roomsAvailable.Rooms, 4)
}
