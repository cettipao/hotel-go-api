package app

import (
	log "github.com/sirupsen/logrus"
	hotelController "mvc-go/controllers/hotel"
	reservationController "mvc-go/controllers/reservation"
	userController "mvc-go/controllers/user"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/users", userController.UserInsert)
	router.POST("/login", userController.UserLogin)
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)
	router.GET("/reservation/:id", reservationController.GetReservationById)
	router.GET("/reservation", reservationController.GetReservations)
	router.POST("/reservation", reservationController.InsertReservation)
	router.GET("/rooms-available", reservationController.RoomsAvailable)

	log.Info("Finishing mappings configurations")
}
