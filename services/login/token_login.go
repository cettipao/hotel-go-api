package login

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userId int, userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"user_email": userEmail,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte("your_jwt_secret_key") // Cambia esto por tu propia clave secreta

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
