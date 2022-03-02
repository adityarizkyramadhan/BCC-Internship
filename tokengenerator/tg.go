package tokengenerator

import (
	"BCC-Internship/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenUser(user *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	})
	tokenString, err := token.SignedString([]byte(("petmate")))
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

func GenerateTokenClinic(clinic *user.Clinic) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       clinic.ID,
		"password": clinic.PasswordClinic,
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	})
	tokenString, err := token.SignedString([]byte(("petmate")))
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
