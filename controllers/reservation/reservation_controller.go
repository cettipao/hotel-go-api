package reservationController

import (
	"mvc-go/controllers"
	"mvc-go/dto/reservations_dto"
	service "mvc-go/services"
	"net/http"
	"strconv"
	"time"

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

func GetReservationsByUser(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")

	// Obtener los query params de la solicitud
	hotelID := c.Query("hotel_id")
	date := c.Query("date")

	var reservationsDto reservations_dto.ReservationsDto
	reservationsDto, err := service.ReservationService.GetReservationsByUser(userID, hotelID, date)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservationsDto)
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

	// Verificar si alguno de los campos está vacío
	if controllers.IsEmptyField(reservationCreateDto.InitialDate) ||
		controllers.IsEmptyField(reservationCreateDto.FinalDate) ||
		reservationCreateDto.HotelId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios esta vacio o no se envio"})
		return
	}

	// Parsear las fechas iniciales y finales
	layout := "02/01/2006"
	initialDate, err := time.Parse(layout, reservationCreateDto.InitialDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha inicial inválida"})
		return
	}

	finalDate, err := time.Parse(layout, reservationCreateDto.FinalDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha final inválida"})
		return
	}

	// Verificar si la fecha final es mayor que la fecha inicial
	if finalDate.Before(initialDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha final debe ser mayor que la fecha inicial"})
		return
	}

	//Verificar que haya rooms disponibles
	rooms_available, _ := service.ReservationService.RoomsAvailableHotel(reservationCreateDto)
	if rooms_available.Rooms <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No hay habitaciones disponibles para esa fecha"})
		return
	}

	userID := c.GetInt("user_id")
	reservationCreateDto.UserId = userID

	var reservationDetailDto reservations_dto.ReservationDetailDto
	reservationDetailDto, er := service.ReservationService.InsertReservation(reservationCreateDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservationDetailDto)
}

func RoomsAvailable(c *gin.Context) {
	// Obtener los parámetros de consulta
	params := c.Request.URL.Query()

	// Obtener los valores de los parámetros
	initialDate := params.Get("initial_date")
	finalDate := params.Get("final_date")

	roomsAvailable, err := service.ReservationService.RoomsAvailable(initialDate, finalDate)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, roomsAvailable)
}
