package clients

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
)

var Db *gorm.DB

func GetAmenitieById(id int) model.Amenitie {
	var amenitie model.Amenitie

	Db.Where("id = ?", id).First(&amenitie)
	log.Debug("Amenitie: ", amenitie)

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
