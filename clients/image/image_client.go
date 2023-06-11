package clients

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

var Db *gorm.DB

func GetImageById(id int) model.Image {
	var image model.Image

	Db.Where("id = ?", id).First(&image)
	log.Debug("Image: ", image)

	return image
}

func GetImagesByHotelId(hotelID int) model.Images {
	var images model.Images

	Db.Where("hotel_id = ?", hotelID).Find(&images)
	log.Debug("Images: ", images)

	return images
}

func GetImages() model.Images {
	var images model.Images

	Db.Find(&images)
	log.Debug("Images: ", images)

	return images
}

func InsertImage(image model.Image) model.Image {
	result := Db.Create(&image)

	if result.Error != nil {
		// TODO: Manejar errores
		log.Error("")
	}
	log.Debug("Image Created: ", image.ID)
	return image
}

func DeleteImageById(id int) e.ApiError {

	err := Db.Delete(&model.Image{}, id).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Image not found")
		}

		return e.NewBadRequestApiError("Failed to delete Image")
	}

	return nil // Sin errores, se eliminó el amenitie correctamente
}