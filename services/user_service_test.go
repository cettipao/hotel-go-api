package services

import (
	"github.com/stretchr/testify/assert"
	userClient "mvc-go/clients/user"
	"mvc-go/dto/users_dto"
	"mvc-go/model"
	"testing"
)

type TestUsers struct {
}

func (testUsers TestUsers) GetUserByEmail(email string) model.User {
	user1 := model.User{
		Id:       1,
		Name:     "John",
		LastName: "Doe",
		Dni:      "123456789",
		Email:    email,
		Password: "$2a$14$0fuoXVH9zAPfWb8WtYQ.7.jDyF/lgaEaHdmTgXVAVs/s874cqWOKq",
		Admin:    0,
	}
	return user1
}

func TestUserLogin(t *testing.T) {
	//prepare
	userClient.MyClient = TestUsers{}
	userLogin := users_dto.UserLoginDto{
		Email:    "john.doe@example.com",
		Password: "user",
	}

	//act
	response, _ := UserService.UserLogin(userLogin)

	//assert
	assert.NotEmpty(t, response.Token)
}

func TestUserLoginFailed(t *testing.T) {
	//prepare
	userClient.MyClient = TestUsers{}
	userLogin := users_dto.UserLoginDto{
		Email:    "john.doe@example.com",
		Password: "x",
	}

	//act
	response, err := UserService.UserLogin(userLogin)

	//assert
	assert.Equal(t, response.Token, "")
	assert.Equal(t, err.Status(), 401)
}

func TestEmailTaken(t *testing.T) {
	//prepare
	userClient.MyClient = TestUsers{}

	//act
	response := UserService.IsEmailTaken("user@user.com")

	//assert
	assert.Equal(t, response, true)
}
