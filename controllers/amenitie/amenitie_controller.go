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

	// Verificar si alguno de los campos está vacío
	if controllers.IsEmptyField(amenitieDto.Name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios esta vacio o no se envio"})
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

func DeleteAmenitieById(c *gin.Context) {
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

	// Obtiene el ID del amenitie de los parámetros de la solicitud
	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenitie ID"})
		return
	}

	// Llama al servicio para eliminar el hotel por su ID
	err = service.AmenitieService.DeleteAmenitieById(hotelID)
	if err != nil {
		// Verifica si se produjo un error específico de "hotel no encontrado"
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Si no se produjo ningún error, devuelve una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Amenitie deleted successfully"})
}

func GetHotelAmenities(c *gin.Context) {
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

	// Obtener los hoteles y amenidades desde el servicio
	hotelsDto, err := service.HotelService.GetHotels()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	// Obtener la lista de hoteles y amenidades desde los objetos Dto
	hotels := hotelsDto.Hotels

	// Crear la estructura de datos para las relaciones
	hotelAmenities := make([]map[string]interface{}, 0)

	// Recorrer los hoteles y amenidades para generar las relaciones
	for _, hotel := range hotels {

		amenitiesDto, err := service.AmenitieService.GetAmenitiesByHotelId(hotel.Id)
		if err != nil {
			c.JSON(err.Status(), err)
			return
		}
		amenities := amenitiesDto.Amenities
		for _, amenitie := range amenities {
			hotelAmenity := make(map[string]interface{})
			hotelAmenity["hotel_name"] = hotel.Name
			hotelAmenity["hotel_id"] = hotel.Id
			hotelAmenity["amenitie"] = amenitie.Name
			hotelAmenity["amenitie_id"] = amenitie.Id

			hotelAmenities = append(hotelAmenities, hotelAmenity)
		}
	}

	// Retornar el resultado como JSON
	c.JSON(http.StatusOK, gin.H{
		"hotel_amenitie": hotelAmenities,
	})
}
