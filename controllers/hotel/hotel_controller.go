package userController

import (
	"mvc-go/controllers"
	"mvc-go/dto/hotels_dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto hotels_dto.HotelDetailDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto hotels_dto.HotelsDto
	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func HotelInsert(c *gin.Context) {
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

	var hotelDto hotels_dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func AddAmenitieToHotel(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")
	// Verificar si es admin
	if !controllers.IsAdmin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Debes tener permisos de administrador para realizar esta accion"})
		return
	}

	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	amenitieID, err := strconv.Atoi(c.Param("id_amenitie"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenitie ID"})
		return
	}

	er := service.HotelService.AddAmenitieToHotel(hotelID, amenitieID)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Amenitie added to hotel successfully"})
}
