package reservationController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetReservationById(c *gin.Context) {
	log.Debug("Reservation id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservationDetailDto dto.ReservationDetailDto

	reservationDetailDto, err := service.ReservationService.GetReservationById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservationDetailDto)
}

func GetReservations(c *gin.Context) {
	var reservationsDetailDto dto.ReservationsDetailDto
	reservationsDetailDto, err := service.ReservationService.GetReservations()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservationsDetailDto)
}

func InsertReservation(c *gin.Context) {
	var reservationCreateDto dto.ReservationCreateDto
	err := c.BindJSON(&reservationCreateDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservationDetailDto dto.ReservationDetailDto
	reservationDetailDto, er := service.ReservationService.InsertReservation(reservationCreateDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservationDetailDto)
}

func RoomsAvailable(c *gin.Context) {
	var reservationCreateDto dto.ReservationCreateDto
	err := c.BindJSON(&reservationCreateDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var roomsAvailable dto.RoomsAvailable

	roomsAvailable, er := service.ReservationService.RoomsAvailable(reservationCreateDto)

	if err != nil {
		c.JSON(er.Status(), err)
		return
	}
	c.JSON(http.StatusOK, roomsAvailable)
}
