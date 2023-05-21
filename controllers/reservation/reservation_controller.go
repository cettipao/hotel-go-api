package reservationController

import (
	"mvc-go/controllers"
	"mvc-go/dto/reservations_dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetReservationById(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")
	//Verificar si es admin
	if !controllers.IsAdmin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Debes tener permisos de administrador para realizar esta accion"})
		return
	}

	log.Debug("Reservation id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservationDetailDto reservations_dto.ReservationDetailDto

	reservationDetailDto, err := service.ReservationService.GetReservationById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservationDetailDto)
}

func GetReservations(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")
	//Verificar si es admin
	if !controllers.IsAdmin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Debes tener permisos de administrador para realizar esta accion"})
		return
	}

	var reservationsDetailDto reservations_dto.ReservationsDetailDto
	reservationsDetailDto, err := service.ReservationService.GetReservations()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservationsDetailDto)
}

func InsertReservation(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var reservationCreateDto reservations_dto.ReservationCreateDto
	err := c.BindJSON(&reservationCreateDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservationDetailDto reservations_dto.ReservationDetailDto
	reservationDetailDto, er := service.ReservationService.InsertReservation(reservationCreateDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservationDetailDto)
}

// NO NECESITO AUTH
func RoomsAvailable(c *gin.Context) {
	var reservationCreateDto reservations_dto.ReservationCreateDto
	err := c.BindJSON(&reservationCreateDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var roomsAvailable reservations_dto.RoomsAvailable

	roomsAvailable, er := service.ReservationService.RoomsAvailable(reservationCreateDto)

	if err != nil {
		c.JSON(er.Status(), err)
		return
	}
	c.JSON(http.StatusOK, roomsAvailable)
}
