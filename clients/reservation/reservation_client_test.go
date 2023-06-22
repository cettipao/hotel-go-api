package clients

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"mvc-go/model"
	"testing"
	"time"
)

func init() {
	// DB Connections Paramters
	DBName := "hotel"
	DBUser := "cetti"
	DBPass := "123456"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
	// ------------------------

	db, err := gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Realizar las migraciones necesarias para crear las tablas en la base de datos
	db.AutoMigrate(&model.Reservation{})

	// Asignar la base de datos inicializada al cliente
	Db = db

	// Asignar el cliente de producción al cliente actual
	MyClient = ProductionClient{}
}

func TestInsertReservation(t *testing.T) {
	// Crear una nueva reserva
	reservation := model.Reservation{
		UserId:      1,
		HotelId:     1,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 0, 3),
	}

	// Insertar la reserva en la base de datos
	insertedReservation := MyClient.InsertReservation(reservation)

	// Verificar que la reserva tenga un ID asignado
	assert.NotZero(t, insertedReservation.Id)

	// Eliminar la reserva
	DeleteReservationById(insertedReservation.Id)
}

func TestGetReservationById(t *testing.T) {
	// Obtener la reserva con ID 1
	reservation := MyClient.GetReservationById(30)

	// Verificar que se obtenga la reserva correcta
	assert.Equal(t, 30, reservation.Id)

	// Obtener una reserva inexistente con ID 5000
	reservation = MyClient.GetReservationById(5000)

	// Verificar que se obtenga una reserva vacía
	assert.Equal(t, model.Reservation{}, reservation)
}

func TestGetReservationsByUser(t *testing.T) {

	// Obtener las reservas de un usuario inexistente con ID 5000
	reservations := MyClient.GetReservationsByUser(5000)

	// Verificar que se obtengan reservas vacías
	assert.Empty(t, reservations)
}

func TestGetReservationsByHotelAndDate(t *testing.T) {
	// Obtener el número de reservas del hotel con ID 1 en una fecha específica
	count := MyClient.GetReservationsByHotelAndDate(1, time.Now())

	// Verificar que se obtenga el número correcto de reservas
	assert.NotZero(t, count)
}

func TestDeleteReservationById(t *testing.T) {
	// Insertar una nueva reserva para eliminarla posteriormente
	reservation := model.Reservation{
		UserId:      1,
		HotelId:     1,
		InitialDate: time.Now(),
		FinalDate:   time.Now().AddDate(0, 0, 3),
	}
	insertedReservation := MyClient.InsertReservation(reservation)

	// Eliminar la reserva por su ID
	err := DeleteReservationById(insertedReservation.Id)

	// Verificar que no haya errores al eliminar la reserva
	assert.Nil(t, err)

	// Obtener la reserva eliminada por su ID
	deletedReservation := MyClient.GetReservationById(insertedReservation.Id)

	// Verificar que se obtenga una reserva vacía
	assert.Equal(t, model.Reservation{}, deletedReservation)
}
