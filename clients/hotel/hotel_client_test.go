package clients

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"mvc-go/model"
	"testing"
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
	db.AutoMigrate(&model.Hotel{})

	// Asignar la base de datos inicializada al cliente
	Db = db

	// Asignar el cliente de producción al cliente actual
	MyClient = ProductionClient{}
}

func TestGetHotelById(t *testing.T) {
	// Obtener el hotel con ID 1
	hotel := MyClient.GetHotelById(1)

	// Verificar que se obtenga el hotel correcto
	assert.Equal(t, 1, hotel.Id)

	// Obtener un hotel inexistente con ID 5000
	hotel = MyClient.GetHotelById(5000)

	// Verificar que se obtenga un hotel vacío
	assert.Equal(t, model.Hotel{}, hotel)
}

func TestGetHotels(t *testing.T) {
	// Obtener todos los hoteles
	hotels := MyClient.GetHotels()

	// Verificar que se obtenga al menos un hotel
	assert.NotEmpty(t, hotels)

	// Obtener el número total de hoteles en la base de datos
	var totalHotels int
	Db.Model(&model.Hotel{}).Count(&totalHotels)

	// Verificar que se obtengan todos los hoteles
	assert.Equal(t, totalHotels, len(hotels))
}

func TestInsertHotel(t *testing.T) {
	// Crear un nuevo hotel
	hotel := model.Hotel{
		Name: "Hotel Example",
	}

	// Insertar el hotel en la base de datos
	insertedHotel := InsertHotel(hotel)

	// Verificar que el hotel tenga un ID asignado
	assert.Equal(t, "Hotel Example", insertedHotel.Name)
	assert.NotZero(t, insertedHotel.Id)

	//Eliminar el hotel
	DeleteHotelById(insertedHotel.Id)
}
