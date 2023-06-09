package clients

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

var Db *gorm.DB

type HotelClientInterface interface {
	GetHotels() model.Hotels
	GetHotelById(id int) model.Hotel
}

var (
	MyClient HotelClientInterface
)

type ProductionClient struct{}

func (HotelClientInterface ProductionClient) GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	//Db.Where("id = ?", id).Preload("Address").Preload("Telephones").First(&users_dto)
	Db.Where("id = ?", id).Preload("Amenities").Preload("Images").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func (HotelClientInterface ProductionClient) GetHotels() model.Hotels {
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

func DeleteHotelById(id int) e.ApiError {
	// Obtén el hotel por su ID antes de eliminarlo
	var hotel model.Hotel
	if err := Db.First(&hotel, id).Error; err != nil {
		// Maneja el error de búsqueda del hotel según sea necesario
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Hotel not found")
		}
		return e.NewBadRequestApiError("Failed to delete hotel")
	}

	// Elimina el hotel por su ID
	if err := Db.Delete(&hotel).Error; err != nil {
		// Maneja el error de eliminación del hotel según sea necesario
		return e.NewBadRequestApiError("Failed to delete hotel")
	}

	return nil // Sin errores, se eliminó el hotel correctamente
}

func UpdateHotel(hotel model.Hotel) e.ApiError {
	err := Db.Save(&hotel)
	if err != nil {
		//TODO Manage Errors
		log.Error(err)
		return e.NewBadRequestApiError("Failed to delete hotel amenities")
	}
	log.Debug("Hotel Updated: ", hotel.Id)
	return nil
}

func IsAmenitieAlreadyLinked(hotelID, amenitieID int) bool {
	var count int
	err := Db.Table("hotel_amenities").
		Where("hotel_id = ? AND amenitie_id = ?", hotelID, amenitieID).
		Count(&count).
		Error

	return err == nil && count > 0
}

func DeleteLinkAmenitieHotel(hotelID int, amenitieID int) bool {
	// Eliminar la fila que vincula el hotel y la amenidad en "hotel_amenities"
	result := Db.Table("hotel_amenities").
		Where("hotel_id = ? AND amenitie_id = ?", hotelID, amenitieID).
		Delete(nil)
	if result.Error != nil {
		// Manejar el error en caso de que ocurra
		return false
	}
	return true
}
