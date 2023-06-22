package userController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mvc-go/dto/users_dto"
	service "mvc-go/services"
	"mvc-go/services/login"
	e "mvc-go/utils/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	service.UserService = &TestUser{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/login", UserLogin)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
    "email":"test@test.com",
    "password":"test"
}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/login", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	//token, _ := login.GenerateToken(0, "test@test.com")
	//req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, resp.Code)

	jsonResponse := resp.Body.Bytes()

	// Convertir los datos en formato JSON a un objeto de tipo RoomsResponse
	var loginResponse users_dto.UserLoginResponseDto
	err := json.Unmarshal(jsonResponse, &loginResponse)
	if err != nil {
		fmt.Println("Error al convertir los datos JSON:", err)
		return
	}
	assert.Equal(t, "token", loginResponse.Token)

}

func TestLoginNoPassword(t *testing.T) {
	service.UserService = &TestUser{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/login", UserLogin)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
    "email":"test@test.com"
}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/login", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	//token, _ := login.GenerateToken(0, "test@test.com")
	//req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, 400, resp.Code)

	// Convertir los datos en formato JSON
	var loginResponse users_dto.UserLoginResponseDto
	jsonResponse := resp.Body.Bytes()
	err := json.Unmarshal(jsonResponse, &loginResponse)
	if err != nil {
		fmt.Println("Error al convertir los datos JSON:", err)
		return
	}
	assert.Equal(t, 400, resp.Code)

}

func TestGetMyUser(t *testing.T) {
	service.UserService = &TestUser{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/my-user", GetMyUser)

	req, _ := http.NewRequest("POST", "/my-user", nil)

	// Establecer el encabezado de autenticación con un token válido
	token, _ := login.GenerateToken(0, "test@test.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Convertir los datos en formato JSON
	jsonResponse := resp.Body.Bytes()
	var userDetail users_dto.UserDetailDto
	err := json.Unmarshal(jsonResponse, &userDetail)
	if err != nil {
		fmt.Println("Error al convertir los datos JSON:", err)
		return
	}
	assert.Equal(t, "test@test.com", userDetail.Email)

}

func TestRegister(t *testing.T) {
	service.UserService = &TestUser{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/register", UserInsert)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
    "name":"test",
    "last_name":"test",
    "email":"test@test.com",
    "dni":"test",
    "password":"test"
}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/register", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	//token, _ := login.GenerateToken(0, "test@test.com")
	//req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, resp.Code)

	jsonResponse := resp.Body.Bytes()

	// Convertir los datos en formato JSON a un objeto de tipo RoomsResponse
	var registerResponse users_dto.UserDetailDto
	err := json.Unmarshal(jsonResponse, &registerResponse)
	if err != nil {
		fmt.Println("Error al convertir los datos JSON:", err)
		return
	}
	assert.Equal(t, "test", registerResponse.Name)

}

func TestRegisterNoPassword(t *testing.T) {
	service.UserService = &TestUser{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/register", UserInsert)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
    "name":"test",
    "last_name":"test",
    "email":"test@test.com",
    "dni":"test"
}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/register", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	//token, _ := login.GenerateToken(0, "test@test.com")
	//req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, 400, resp.Code)

}

type TestUser struct {
}

func (testUser TestUser) GetUserById(id int) (users_dto.UserDetailDto, e.ApiError) {
	return users_dto.UserDetailDto{
		Name:     "test",
		LastName: "test",
		Dni:      "test",
		Email:    "test@test.com",
		Admin:    0,
	}, nil
}
func (testUser TestUser) GetUsers() (users_dto.UsersDto, e.ApiError) {
	return users_dto.UsersDto{}, nil
}
func (testUser TestUser) InsertUser(userDto users_dto.UserDtoRegister) (users_dto.UserDetailDto, e.ApiError) {
	return users_dto.UserDetailDto{
		Name:     "test",
		LastName: "test",
		Dni:      "test",
		Email:    "test@test.com",
		Admin:    0,
	}, nil
}
func (testUser TestUser) UserLogin(userDto users_dto.UserLoginDto) (users_dto.UserLoginResponseDto, e.ApiError) {
	return users_dto.UserLoginResponseDto{
		Token: "token",
	}, nil
}
func (testUser TestUser) IsEmailTaken(email string) bool {
	return false
}
func (testUser TestUser) DeleteUserById(id int) e.ApiError {
	return nil
}
func (testUser TestUser) UpdateUser(userDto users_dto.UserDtoRegister, id int) e.ApiError {
	return nil
}
