package clients

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

var Db *gorm.DB

func GetAmenitieById(id int) model.Amenitie {
	var amenitie model.Amenitie

	Db.Where("id = ?", id).First(&amenitie)
	log.Debug("Amenitie: ", amenitie)

	return amenitie
}

func UpdateAmenitie(amenitie model.Amenitie) model.Amenitie {
	result := Db.Save(&amenitie)
	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Amenitie Updated: ", amenitie.Id)
	return amenitie
}

func GetAmenities() model.Amenities {
	var amenities model.Amenities

	Db.Find(&amenities)
	log.Debug("Amenities: ", amenities)

	return amenities
}

func InsertAmenitie(amenitie model.Amenitie) model.Amenitie {
	result := Db.Create(&amenitie)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Amenitie Created: ", amenitie.Id)
	return amenitie
}

func GetAmenitiesByHotelId(hotelId int) model.Amenities {
	var amenities model.Amenities

	Db.Table("amenities").
		Joins("JOIN hotel_amenities ON amenities.id = hotel_amenities.amenitie_id").
		Where("hotel_amenities.hotel_id = ?", hotelId).
		Find(&amenities)
	log.Debug("Amenities by Hotel ID: ", amenities)

	return amenities
}

func DeleteAmenitieById(id int) e.ApiError {
	// Lógica para eliminar el hotel por su ID en la capa de datos o API externa
	err := Db.Delete(&model.Amenitie{}, id).Error
	if err != nil {
		// Manejar el error de eliminación del hotel según sea necesario
		// Por ejemplo, puedes verificar si el error es de "registro no encontrado"
		// y devolver un error personalizado utilizando la función NewNotFoundError
		// de tu paquete de excepciones.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Hotel not found")
		}
		// Otros errores de eliminación del hotel
		return e.NewBadRequestApiError("Failed to delete amenitie")
	}

	return nil // Sin errores, se eliminó el amenitie correctamente
}
