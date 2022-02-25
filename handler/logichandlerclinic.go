package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewClinicalHandler(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
	}
	var body user.NewClinic
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Bad Request",
			"Message": "Error when binding JSON",
			"Error":   err.Error(),
		})
	}
	clinic := user.Clinic{
		NameClinic:     body.NameClinic,
		Address:        body.Address,
		Contact:        body.Contact,
		UsernameClinic: body.UsernameClinic,
		PasswordClinic: body.PasswordClinic,
	}
	if db.Create(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error creating clinic",
			"Error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Status OK",
		"Message": "Clinic created",
		"Clinic":  clinic,
	})
}
