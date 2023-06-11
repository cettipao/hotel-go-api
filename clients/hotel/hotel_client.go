package clients

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	//Db.Where("id = ?", id).Preload("Address").Preload("Telephones").First(&users_dto)
	Db.Where("id = ?", id).Preload("Amenities").Preload("Images").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotels() model.Hotels {
	var hotels model.Hotels
	//Db.Preload("Address").Find(&users)
	Db.Preload("Amenities").Preload("Images").Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}

func InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}

func UpdateHotel(hotel model.Hotel) {
	Db.Save(&hotel)
	log.Debug("Hotel Updated: ", hotel.Id)
}

func IsAmenitieAlreadyLinked(hotelID, amenitieID int) bool {
	var count int
	err := Db.Table("hotel_amenities").
		Where("hotel_id = ? AND amenitie_id = ?", hotelID, amenitieID).
		Count(&count).
		Error

	return err == nil && count > 0
}
