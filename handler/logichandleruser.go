package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
	}
	var body user.NewUser
	c.BindJSON(&body)
	user := user.NewUser{
		Name:     body.Name,
		Contact:  body.Contact,
		Username: body.Username,
		Password: body.Password,
		Address:  body.Address,
	}
	if db.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error creating user",
			"Error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
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
	}
	var body user.UserLogin
	c.BindJSON(&body)
	if 

}
