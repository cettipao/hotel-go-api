package amenitieController

import (
	"mvc-go/controllers"
	amenities_dto "mvc-go/dto/hotels_dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAmenitieById(c *gin.Context) {
	log.Debug("Amenitie ID to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var amenitieDto amenities_dto.AmenitieDto

	amenitieDto, err := service.AmenitieService.GetAmenitieById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, amenitieDto)
}

func GetAmenities(c *gin.Context) {
	var amenitiesDto amenities_dto.AmenitiesDto
	amenitiesDto, err := service.AmenitieService.GetAmenities()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, amenitiesDto)
}

func InsertAmenitie(c *gin.Context) {
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

	var amenitieDto amenities_dto.AmenitieDto
	err := c.BindJSON(&amenitieDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	amenitieDto, er := service.AmenitieService.InsertAmenitie(amenitieDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, amenitieDto)
}
