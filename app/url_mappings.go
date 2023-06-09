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
	router.DELETE("/user/:id", userController.DeleteUserById)
	router.PUT("/user/:id", userController.UpdateUserById)
	router.GET("/user", userController.GetUsers)
	router.GET("/myuser", userController.GetMyUser)
	router.POST("/user", userController.UserInsert)
	router.POST("/login", userController.UserLogin)
	// Hotels Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.PUT("/hotel/:id", hotelController.UpdateHotelById)
	router.DELETE("/hotel/:id", hotelController.DeleteHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)
	router.PUT("/hotel/:id/add-amenitie/:id_amenitie", hotelController.AddAmenitieToHotel)
	router.DELETE("/hotel/:id/remove-amenitie/:id_amenitie", hotelController.RemoveAmenitieToHotel)
	router.POST("/hotel/:id/add-image", imageController.InsertImage)
	// Amenities Mapping
	router.GET("/amenitie/:id", amenitieController.GetAmenitieById)
	router.PUT("/amenitie/:id", amenitieController.UpdateAmenitie)
	router.DELETE("/amenitie/:id", amenitieController.DeleteAmenitieById)
	router.GET("/amenitie", amenitieController.GetAmenities)
	router.GET("/hotel_amenitie", amenitieController.GetHotelAmenities)
	router.POST("/amenitie", amenitieController.InsertAmenitie)
	// Images Mapping
	//router.GET("/images/:id", imageController.GetImageById)
	router.GET("/image/hotel/:id", imageController.GetImagesByHotelId)
	router.GET("/image", imageController.GetImages)
	router.DELETE("/image/:id", imageController.DeleteImageById)
	// Reservations Mapping
	router.GET("/reservation/:id", reservationController.GetReservationById)
	router.DELETE("/reservation/:id", reservationController.DeleteReservationById)
	router.GET("/reservation", reservationController.GetReservations)
	router.GET("/my-reservations", reservationController.GetReservationsByUser)
	router.POST("/reservation", reservationController.InsertReservation)

	router.GET("/rooms-available", reservationController.RoomsAvailable)

	log.Info("Finishing mappings configurations")
}
