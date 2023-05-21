package controllers

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	service "mvc-go/services"
)

func TokenVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(errors.New("No Token Provided"))
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Aquí debes proporcionar la clave secreta utilizada para firmar los tokens
			// Puedes ajustar esta lógica según tu implementación
			return []byte("your_jwt_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.Error(errors.New("Token not valid"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Error(errors.New("Unauthorized"))
			c.Abort()
			return
		}

		// Si el token es válido, puedes extraer el ID del usuario
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.Error(errors.New("Unauthorized"))
			c.Abort()
			return
		}

		c.Set("user_id", int(userID))
		c.Next()
	}
}

func IsAdmin(userId int) bool {
	userDto, err := service.UserService.GetUserById(userId)
	if err != nil {
		log.Error("Error en el get de usuario del token")
		return false
	}
	if userDto.Admin == 1 {
		return true
	}
	return false

}
