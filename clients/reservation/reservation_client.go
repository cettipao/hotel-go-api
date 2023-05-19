package clients

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
)

var Db *gorm.DB

func GetReservationById(id int) model.Reservation {
	var reservation model.Reservation

	Db.Where("id = ?", id).Preload("Hotel").Preload("User").First(&reservation)
	//Db.Where("id = ?", id).First(&user)
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
	Db.Where("id_user = ?", idUser).Preload("Hotel").Preload("User").Find(&reservations)

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

func InsertReservation(reservation model.Reservation) model.Reservation {
	result := Db.Create(&reservation)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Reservation Created: ", reservation.Id)
	return reservation
}
