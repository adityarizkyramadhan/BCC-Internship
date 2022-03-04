package middleware

import (
	"BCC-Internship/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckJwtUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		bearerToken := tokenString[7:]
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("petmate"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Error when parsing token",
				"data":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			idUser := int(claims["id"].(float64))
			fmt.Println(idUser)
			userIn := helper.SearchUserById(idUser)
			fmt.Println(userIn)
			c.Set("userlogin", userIn)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Invalid token",
				"data":    err.Error(),
			})
		}
	}
}

func CheckJwtClinic() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		bearerToken := tokenString[7:]
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("petmate"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Error when parsing token",
				"data":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			idUser := int(claims["id"].(float64))
			userIn := helper.SearchClinicById(idUser)
			c.Set("clinic", userIn)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": "Invalid token",
				"data":    err.Error(),
			})
		}
	}
}
