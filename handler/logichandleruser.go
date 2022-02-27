package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewUserHandler(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
		return
	}
	var body user.NewUser
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Bad Request",
			"Message": "Error when binding JSON",
			"Error":   err.Error(),
		})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when hashing password",
			"Error":   err.Error(),
		})
		return
	}
	user := user.User{
		Name:     body.Name,
		Contact:  body.Contact,
		Username: body.Username,
		Password: string(password),
		Address:  body.Address,
	}
	if db.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error creating user",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Status OK",
		"Message": "User created",
		"User":    user,
	})
}

func UserLogin(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
		return
	}
	var body user.UserLogin
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Bad Request",
			"Message": "Error when binding JSON",
			"Error":   err.Error(),
		})
		return
	}
	var user user.User
	if db.Where("username = ?", body.Username).Take(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Status Internal Server Error",
			"Message": "Error when querrying username",
			"Error":   err.Error(),
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":  "Unauthorized",
			"Message": "Error when comparing password",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Status OK",
		"Message": "User logged in",
		"User":    user,
	})
}
