package imageController

import (
	"mvc-go/controllers"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetImageById(c *gin.Context) {
	log.Debug("Image id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	imageDto, err := service.ImageService.GetImageById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imageDto)
}

func GetImagesByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load images: " + c.Param("id"))

	hotelID, _ := strconv.Atoi(c.Param("id"))
	imagesDto, err := service.ImageService.GetImagesByHotelId(hotelID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}

func GetImages(c *gin.Context) {

	imagesDto, err := service.ImageService.GetImages()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}

func InsertImage(c *gin.Context) {
	// Obtener el ID del hotel del contexto o de los parámetros de la ruta, según sea necesario
	hotelID, erint := strconv.Atoi(c.Param("id"))
	if erint != nil {
		// Manejar el error si la conversión falla
		c.JSON(http.StatusBadRequest, gin.H{"error": erint.Error()})
		return
	}

	// Obtener los datos de la imagen del cuerpo de la solicitud
	imageFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar la imagen y manejar la lógica de relación con el hotel
	imageDto, er := service.ImageService.InsertImage(hotelID, imageFile)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, imageDto)
}

func DeleteImageById(c *gin.Context) {
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
	imageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenitie ID"})
		return
	}

	// Llama al servicio para eliminar el hotel por su ID
	err = service.ImageService.DeleteImageById(imageId)
	if err != nil {
		// Verifica si se produjo un error específico de "hotel no encontrado"
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Si no se produjo ningún error, devuelve una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Amenitie deleted successfully"})
}
