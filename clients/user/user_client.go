package clients

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

var Db *gorm.DB

type UserClientInterface interface {
	GetUserByEmail(email string) model.User
}

var (
	MyClient UserClientInterface
)

type ProductionClient struct{}

func GetUserById(id int) model.User {
	var user model.User

	//Db.Where("id = ?", id).Preload("Address").Preload("Telephones").First(&users_dto)
	Db.Where("id = ?", id).Preload("Reservations").First(&user)
	log.Debug("User: ", user)

	return user
}

func (UserClientInterface ProductionClient) GetUserByEmail(email string) model.User {
	var user model.User

	Db.Where("email = ?", email).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users
	//Db.Preload("Address").Find(&users)
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}

func DeleteUserById(id int) e.ApiError {
	// Obtén el hotel por su ID antes de eliminarlo
	var user model.User
	if err := Db.First(&user, id).Error; err != nil {
		// Maneja el error de búsqueda del hotel según sea necesario
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Reservation not user")
		}
		return e.NewBadRequestApiError("Failed to delete user")
	}

	if err := Db.Delete(&user).Error; err != nil {
		// Maneja el error de eliminación del hotel según sea necesario
		return e.NewBadRequestApiError("Failed to delete user")
	}

	return nil // Sin errores, se eliminó el hotel correctamente
}

func UpdateUser(user model.User) e.ApiError {
	log.Debug(user)
	Db.Save(&user)
	log.Debug("user Updated: ", user.Id)
	return nil
}
