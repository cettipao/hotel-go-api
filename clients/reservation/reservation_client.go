package clients

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

var Db *gorm.DB

func GetReservationById(id int) model.Reservation {
	var reservation model.Reservation

	Db.Where("id = ?", id).Preload("Hotel").Preload("User").First(&reservation)
	//Db.Where("id = ?", id).First(&users_dto)
	log.Debug("Reservation: ", reservation)

	return reservation
}

func GetReservations() model.Reservations {
	var reservations model.Reservations
	//Db.Preload("Address").Find(&users)
	Db.Preload("Hotel").Preload("User").Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByUser(idUser int) model.Reservations {
	var reservations model.Reservations
	//Db.Preload("Address").Find(&users)
	Db.Where("user_id = ?", idUser).Preload("Hotel").Preload("User").Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByHotel(idHotel int) model.Reservations {
	var reservations model.Reservations
	//Db.Preload("Address").Find(&users)
	Db.Where("id_hotel = ?", idHotel).Preload("Hotel").Preload("User").Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByHotelAndDates(idHotel int, initialDate time.Time, finalDate time.Time) int {
	var count int
	Db.Model(&model.Reservation{}).Where("hotel_id = ? AND ? < final_date AND ? >= initial_date", idHotel, initialDate, finalDate).Preload("Hotel").Preload("User").Count(&count)

	log.Debug("Reservation Count: ", count)

	return count
}

func InsertReservation(reservation model.Reservation) model.Reservation {
	result := Db.Create(&reservation)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Reservation Created: ", reservation.Id)
	return reservation
}

func GetReservationsByUserAndHotel(idUser int, hotelID string) model.Reservations {
	var reservations model.Reservations
	Db.Where("user_id = ? AND hotel_id = ?", idUser, hotelID).Preload("Hotel").Preload("User").Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByUserAndDate(idUser int, date time.Time) model.Reservations {
	var reservations model.Reservations
	Db.Where("user_id = ? AND ? BETWEEN initial_date AND final_date", idUser, date).Preload("Hotel").Preload("User").Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByUserAndHotelAndDate(idUser int, hotelID string, date time.Time) model.Reservations {
	var reservations model.Reservations
	Db.Where("user_id = ? AND hotel_id = ? AND ? BETWEEN initial_date AND final_date", idUser, hotelID, date).Preload("Hotel").Preload("User").Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func DeleteReservationById(id int) e.ApiError {
	// Obtén el hotel por su ID antes de eliminarlo
	var reservation model.Reservation
	if err := Db.First(&reservation, id).Error; err != nil {
		// Maneja el error de búsqueda del hotel según sea necesario
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Reservation not found")
		}
		return e.NewBadRequestApiError("Failed to delete reservations")
	}

	if err := Db.Delete(&reservation).Error; err != nil {
		// Maneja el error de eliminación del hotel según sea necesario
		return e.NewBadRequestApiError("Failed to delete reservation")
	}

	return nil // Sin errores, se eliminó el hotel correctamente
}
