package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/tokengenerator"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewUserHandler(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
		return
	}
	var body user.NewUser
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Error when binding JSON",
			"data":    err.Error(),
		})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when hashing password",
			"data":    err.Error(),
		})
		return
	}
	userPrivate := user.User{
		Name:     body.Name,
		Contact:  body.Contact,
		Username: body.Username,
		Password: string(password),
		Address:  body.Address,
	}
	if db.Create(&userPrivate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error creating user",
			"data":    err.Error(),
		})
		return
	}
	token, err := tokengenerator.GenerateTokenUser(&userPrivate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when generating token",
			"data":    err.Error(),
		})
		return
	}
	getUser := user.ReturnUser{
		ID:      userPrivate.ID,
		Name:    userPrivate.Name,
		Contact: userPrivate.Contact,
		Address: userPrivate.Address,
		Token:   token,
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Status OK",
		"message": "User created",
		"data":    getUser,
	})
}

func UserLogin(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
		return
	}
	var body user.UserLogin
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Error when binding JSON",
			"data":    err.Error(),
		})
		return
	}
	var userPrivate user.User
	if db.Where("username = ?", body.Username).Take(&userPrivate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Status Internal Server Error",
			"message": "Error when querrying username",
			"data":    err.Error(),
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userPrivate.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "Unauthorized",
			"message": "Error when comparing password",
			"data":    err.Error(),
		})
		return
	}
	token, err := tokengenerator.GenerateTokenUser(&userPrivate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when generating token",
			"data":    err.Error(),
		})
		return
	}
	getUser := user.ReturnUser{
		ID:      userPrivate.ID,
		Name:    userPrivate.Name,
		Contact: userPrivate.Contact,
		Address: userPrivate.Address,
		Token:   token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Status OK",
		"message": "User logged in",
		"data":    getUser,
		"token":   token,
	})
}
