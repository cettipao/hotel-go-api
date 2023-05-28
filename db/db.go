package db

import (
	amenitieClient "mvc-go/clients/amenitie"
	hotelClient "mvc-go/clients/hotel"
	imageClient "mvc-go/clients/image"
	reservationClient "mvc-go/clients/reservation"
	userClient "mvc-go/clients/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "hotel"
	DBUser := "root"
	DBPass := "123456"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "mysql"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build

	userClient.Db = db
	hotelClient.Db = db
	reservationClient.Db = db
	amenitieClient.Db = db
	imageClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Reservation{})
	db.AutoMigrate(&model.Amenitie{})
	db.AutoMigrate(&model.Image{})

	log.Info("Finishing Migration Database Tables")
}
