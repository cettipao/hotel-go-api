package reservationController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mvc-go/dto/reservations_dto"
	service "mvc-go/services"
	login "mvc-go/services/login"
	e "mvc-go/utils/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInsertReservation(t *testing.T) {
	service.ReservationService = &TestReservations{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/reservation", InsertReservation)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
		"hotel_id": 2,
		"initial_date": "30/05/2023",
		"final_date": "05/06/2023"
	}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/reservation", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	token, _ := login.GenerateToken(0, "test@test.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, resp.Code)

	jsonResponse := resp.Body.Bytes()

	// Convertir los datos en formato JSON a un objeto de tipo RoomsResponse
	var roomsResponse reservations_dto.RoomsResponse
	err := json.Unmarshal(jsonResponse, &roomsResponse)
	if err != nil {
		fmt.Println("Error al convertir los datos JSON:", err)
		return
	}

}

func TestInsertReservationBadDateFormat(t *testing.T) {
	service.ReservationService = &TestReservations{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/reservation", InsertReservation)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
		"hotel_id": 2,
		"initial_date": "30-05-2023",
		"final_date": "05-06-2023"
	}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/reservation", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	token, _ := login.GenerateToken(0, "test@test.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, 400, resp.Code)

}

func TestInsertReservationBadDateOrder(t *testing.T) {
	service.ReservationService = &TestReservations{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.POST("/reservation", InsertReservation)

	// Crear una solicitud HTTP de tipo POST al endpoint /reservation con el cuerpo JSON
	jsonStr := `{
		"hotel_id": 2,
		"initial_date": "08/05/2023",
		"final_date": "05/05/2023"
	}`
	reqBody := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", "/reservation", reqBody)

	// Establecer el encabezado de autenticación con un token válido
	token, _ := login.GenerateToken(0, "test@test.com")
	req.Header.Set("Authorization", "Bearer "+token)

	// Crear un registrador de respuestas HTTP simulado
	resp := httptest.NewRecorder()

	// Enviar la solicitud al enrutador Gin y capturar la respuesta
	router.ServeHTTP(resp, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, 400, resp.Code)

}

func TestRoomsAvailable(t *testing.T) {
	service.ReservationService = &TestReservations{}
	// Crear un enrutador Gin
	router := gin.Default()

	// Ruta de ejemplo que ejecuta la función RoomsAvailable
	router.GET("/rooms", RoomsAvailable)

	req, _ := http.NewRequest("GET", "/rooms", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	// Obtener los datos en formato JSON de la respuesta
	jsonResponse := resp.Body.Bytes()

	// Convertir los datos en formato JSON a un objeto de tipo RoomsResponse
	var roomsResponse reservations_dto.RoomsResponse
	err := json.Unmarshal(jsonResponse, &roomsResponse)
	if err != nil {
		fmt.Println("Error al convertir los datos JSON:", err)
		return
	}

	// Utilizar el objeto de tipo RoomsResponse
	assert.Equal(t, "Hotel A", roomsResponse.Rooms[0].Name)
}

type TestReservations struct {
}

func (testReservations TestReservations) RoomsAvailable(initialDate string, finalDate string) (reservations_dto.RoomsResponse, e.ApiError) {
	rooms := []reservations_dto.RoomInfo{
		{
			Name:           "Hotel A",
			Id:             123,
			RoomsAvailable: 5,
		},
		{
			Name:           "Hotel B",
			Id:             456,
			RoomsAvailable: 10,
		},
	}

	response := reservations_dto.RoomsResponse{
		Rooms: rooms,
	}
	return response, nil
}

func (testReservations TestReservations) GetReservationById(id int) (reservations_dto.ReservationDetailDto, e.ApiError) {
	return reservations_dto.ReservationDetailDto{}, nil
}
func (testReservations TestReservations) DeleteReservationById(id int) e.ApiError {
	return nil
}
func (testReservations TestReservations) GetReservations() (reservations_dto.ReservationsDetailDto, e.ApiError) {
	return reservations_dto.ReservationsDetailDto{}, nil
}
func (testReservations TestReservations) InsertReservation(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.ReservationDetailDto, e.ApiError) {
	return reservations_dto.ReservationDetailDto{}, nil
}
func (testReservations TestReservations) RoomsAvailableHotel(reservationDto reservations_dto.ReservationCreateDto) (reservations_dto.RoomsAvailable, e.ApiError) {
	roomsAvailable := reservations_dto.RoomsAvailable{Rooms: 5}
	return roomsAvailable, nil
}
func (testReservations TestReservations) GetFilteredReservations(hotelID int, userID int, startDate string, endDate string) (reservations_dto.ReservationsDetailDto, error) {
	return reservations_dto.ReservationsDetailDto{}, nil
}
