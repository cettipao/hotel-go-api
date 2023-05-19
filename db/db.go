package db

import (
	hotelClient "mvc-go/clients/hotel"
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
	DBUser := "cetti"
	DBPass := "123456"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
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

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Reservation{})

	log.Info("Finishing Migration Database Tables")
}
