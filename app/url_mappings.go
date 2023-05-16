package app

import (
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.GET("/hotel/:id", userController.GetUserById)
	router.GET("/hotel", userController.GetUsers)
	router.GET("/rooms-available", userController.GetUsers)
	/*
		router.GET("/user/:id", userController.GetUserById)
		router.GET("/user", userController.GetUsers)
		router.POST("/user", userController.UserInsert)
		router.POST("/user/:id/telephone", userController.AddUserTelephone)

		// Alumnos Mapping
		router.GET("/alumno/:id", alumnoController.GetAlumnoById)
		router.GET("/alumno", alumnoController.GetAlumnos)
		router.POST("/alumno", alumnoController.AlumnoInsert)
		router.POST("/alumno/:id/materia", alumnoController.AddAlumnoMateria)*/

	log.Info("Finishing mappings configurations")
}
