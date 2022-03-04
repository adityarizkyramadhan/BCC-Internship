package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/tokengenerator"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ReadClinic(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
		return
	}
	var clinic []user.Clinic
	if db.Find(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when finding clinic",
			"data":    err.Error(),
		})
		return
	}
	var printClinic []user.PrintClinic
	for i := range clinic {
		printClinic = append(printClinic, user.PrintClinic{
			ID:          clinic[i].ID,
			NameClinic:  clinic[i].NameClinic,
			Address:     clinic[i].Address,
			Contact:     clinic[i].Contact,
			SpreadSheet: clinic[i].SpreadSheet,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Status OK",
		"message": "Clinic found",
		"data":    printClinic,
	})
}

func NewClinicalHandler(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
		return
	}
	var body user.NewClinic
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Bad Request",
			"message": "Error when binding JSON",
			"data":    err.Error(),
		})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.PasswordClinic), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when hashing password",
			"data":    err.Error(),
		})
		return
	}
	clinic := user.Clinic{
		NameClinic:     body.NameClinic,
		Address:        body.Address,
		Contact:        body.Contact,
		UsernameClinic: body.UsernameClinic,
		SpreadSheet:    body.SpreadSheet,
		PasswordClinic: string(password),
	}
	if db.Create(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error creating clinic",
			"data":    err.Error(),
		})
		return
	}
	token, err := tokengenerator.GenerateTokenClinic(&clinic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when generating token",
			"data":    err.Error(),
		})
		return
	}
	clinicReturn := user.GetClinic{
		ID:          clinic.ID,
		NameClinic:  clinic.NameClinic,
		Address:     clinic.Address,
		Contact:     clinic.Contact,
		SpreadSheet: clinic.SpreadSheet,
		Token:       token,
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Status OK",
		"message": "Clinic created",
		"data":    clinicReturn,
	})
}

func ClinicLogin(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
	}
	var body user.ClinicLogin
	if c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Bad Request",
			"message": "Error when binding JSON",
			"data":    err.Error(),
		})
	}
	var clinic user.Clinic
	if db.Where("username_clinic = ?", body.Username).Take(&clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Erorr when querrying database",
			"data":    err.Error(),
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(clinic.PasswordClinic), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "Unauthorized",
			"message": "Invalid username or password",
			"data":    err.Error(),
		})
		return
	}
	token, err := tokengenerator.GenerateTokenClinic(&clinic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when generating token",
			"data":    err.Error(),
		})
		return
	}
	clinicReturn := user.GetClinic{
		ID:          clinic.ID,
		NameClinic:  clinic.NameClinic,
		Address:     clinic.Address,
		Contact:     clinic.Contact,
		SpreadSheet: clinic.SpreadSheet,
		Token:       token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Status OK",
		"message": "Clinic logged in",
		"data":    clinicReturn,
	})
}
