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
	db.AutoMigrate(&model.User{}, &model.Reservation{})

	// Asignar la base de datos inicializada al cliente
	Db = db
}

func TestGetUserById(t *testing.T) {

	// Obtener el usuario con ID 1
	user := GetUserById(1)

	// Verificar que se obtenga el usuario correcto
	assert.Equal(t, 1, user.Id)

	// Obtener un usuario inexistente con ID 5000
	user = GetUserById(5000)

	// Verificar que se obtenga un usuario vacío
	assert.Equal(t, model.User{}, user)
}

func TestGetUserByEmail(t *testing.T) {
	MyClient = ProductionClient{}
	// Obtener el usuario con el correo electrónico "john@example.com"
	user := MyClient.GetUserByEmail("test@test.com")

	// Verificar que se obtenga el usuario correcto
	assert.Equal(t, "test@test.com", user.Email)

	// Obtener un usuario inexistente con otro correo electrónico
	user = MyClient.GetUserByEmail("non@non.com")

	// Verificar que se obtenga un usuario vacío
	assert.Equal(t, model.User{}, user)
}

func TestInsertUser(t *testing.T) {

	// Crear un nuevo usuario
	user := model.User{
		Name:  "Jane Smith",
		Email: "example@example.com",
	}

	// Insertar el usuario en la base de datos
	insertedUser := InsertUser(user)

	// Verificar que el usuario tenga un ID asignado
	assert.Equal(t, "Jane Smith", insertedUser.Name)
	assert.Equal(t, "example@example.com", insertedUser.Email)

	//Eliminar el usuario
	DeleteUserById(insertedUser.Id)
}
