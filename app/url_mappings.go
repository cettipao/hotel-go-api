package app

import (
	log "github.com/sirupsen/logrus"
	amenitieController "mvc-go/controllers/amenitie"
	hotelController "mvc-go/controllers/hotel"
	imageController "mvc-go/controllers/image"
	reservationController "mvc-go/controllers/reservation"
	userController "mvc-go/controllers/user"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.GET("/myuser", userController.GetMyUser)
	router.POST("/user", userController.UserInsert)
	router.POST("/login", userController.UserLogin)
	// Hotels Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)
	router.PUT("/hotel/:id/add-amenitie/:id_amenitie", hotelController.AddAmenitieToHotel)
	router.POST("/hotel/:id/add-image", imageController.InsertImage)
	// Amenities Mapping
	router.GET("/amenitie/:id", amenitieController.GetAmenitieById)
	router.GET("/amenitie", amenitieController.GetAmenities)
	router.POST("/amenitie", amenitieController.InsertAmenitie)
	// Images Mapping
	//router.GET("/images/:id", imageController.GetImageById)
	router.GET("/image/hotel/:id", imageController.GetImagesByHotelId)
	router.GET("/image", imageController.GetImages)
	// Reservations Mapping
	router.GET("/reservation/:id", reservationController.GetReservationById)
	router.GET("/reservation", reservationController.GetReservations)
	router.GET("/my-reservations", reservationController.GetReservationsByUser)
	router.POST("/reservation", reservationController.InsertReservation)

	router.GET("/rooms-available", reservationController.RoomsAvailable)

	log.Info("Finishing mappings configurations")
}
