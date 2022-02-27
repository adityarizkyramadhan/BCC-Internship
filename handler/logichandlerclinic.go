package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ReadClinic(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
		return
	}
	var clinic []user.Clinic
	if db.Find(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when finding clinic",
			"Error":   err.Error(),
		})
		return
	}
	var printClinic []user.GetClinic
	for i := range clinic {
		printClinic = append(printClinic, user.GetClinic{
			ID:         clinic[i].ID,
			NameClinic: clinic[i].NameClinic,
			Address:    clinic[i].Address,
			Contact:    clinic[i].Contact,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Status OK",
		"Message": "Clinic found",
		"Clinic":  printClinic,
	})

}

func NewClinicalHandler(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
		return
	}
	var body user.NewClinic
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Bad Request",
			"Message": "Error when binding JSON",
			"Error":   err.Error(),
		})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.PasswordClinic), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when hashing password",
			"Error":   err.Error(),
		})
		return
	}
	clinic := user.Clinic{
		NameClinic:     body.NameClinic,
		Address:        body.Address,
		Contact:        body.Contact,
		UsernameClinic: body.UsernameClinic,
		PasswordClinic: string(password),
	}
	if db.Create(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error creating clinic",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Status OK",
		"Message": "Clinic created",
		"Clinic":  clinic,
	})
}

func ClinicLogin(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Error when initializing databases",
			"Error":   err.Error(),
		})
	}
	var body user.ClinicLogin
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Bad Request",
			"Message": "Error when binding JSON",
			"Error":   err.Error(),
		})
	}
	var clinic user.Clinic
	if db.Where("usernameclinic = ?", body.Username).Take(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Internal Server Error",
			"Message": "Erorr when querrying database",
			"Error":   err.Error(),
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(clinic.PasswordClinic), []byte(body.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":  "Unauthorized",
			"Message": "Invalid username or password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Status OK",
		"Message": "Clinic logged in",
		"Clinic":  clinic,
	})
}
