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
	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotels() model.Hotels {
	var hotels model.Hotels
	//Db.Preload("Address").Find(&users)
	Db.Find(&hotels)

	log.Debug("Users: ", hotels)

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
