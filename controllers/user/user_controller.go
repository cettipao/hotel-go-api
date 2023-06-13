package userController

import (
	"mvc-go/controllers"
	"mvc-go/dto/users_dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")
	//Verificar si es admin
	if !controllers.IsAdmin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Debes tener permisos de administrador para realizar esta accion"})
		return
	}

	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto users_dto.UserDetailDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func DeleteUserById(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")
	//Verificar si es admin
	if !controllers.IsAdmin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Debes tener permisos de administrador para realizar esta accion"})
		return
	}

	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto users_dto.UserDetailDto

	err := service.UserService.DeleteUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")
	//Verificar si es admin
	if !controllers.IsAdmin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Debes tener permisos de administrador para realizar esta accion"})
		return
	}

	var usersDto users_dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UpdateUserById(c *gin.Context) {
	var userDto users_dto.UserDtoRegister
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Verificar si alguno de los campos está vacío
	if controllers.IsEmptyField(userDto.Name) || controllers.IsEmptyField(userDto.LastName) ||
		controllers.IsEmptyField(userDto.Email) || controllers.IsEmptyField(userDto.Dni) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios esta vacio o no se envio"})
		return
	}

	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	er := service.UserService.UpdateUser(userDto, id)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, "Usuario Modificado con exito")
}

func UserInsert(c *gin.Context) {
	var userDto users_dto.UserDtoRegister
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Verificar si alguno de los campos está vacío
	if controllers.IsEmptyField(userDto.Name) || controllers.IsEmptyField(userDto.LastName) ||
		controllers.IsEmptyField(userDto.Email) || controllers.IsEmptyField(userDto.Dni) ||
		controllers.IsEmptyField(userDto.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios esta vacio o no se envio"})
		return
	}

	// Verificar si el email ya existe
	if service.UserService.IsEmailTaken(userDto.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El Email ingresado ya se encuentra registrado"})
		return
	}

	var userDetailDto users_dto.UserDetailDto
	userDetailDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDetailDto)
}

func UserLogin(c *gin.Context) {
	var userDto users_dto.UserLoginDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Verificar si alguno de los campos está vacío
	if controllers.IsEmptyField(userDto.Email) ||
		controllers.IsEmptyField(userDto.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios esta vacio o no se envio"})
		return
	}

	var loginResponse users_dto.UserLoginResponseDto
	loginResponse, er := service.UserService.UserLogin(userDto)
	// Error del Login
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

func GetMyUser(c *gin.Context) {
	controllers.TokenVerification()(c)
	// Verificar si ocurrió un error durante la verificación del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Obtener el ID del usuario del contexto
	userID := c.GetInt("user_id")

	var userDto users_dto.UserDetailDto

	userDto, err := service.UserService.GetUserById(userID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}
