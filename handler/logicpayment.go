package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	userLogin := c.MustGet("user")
	var PayIn model.PaymentInput
	err := c.BindJSON(&PayIn)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	var IdUri model.InputUriClinic
	err = c.BindUri(&IdUri)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}

	paymentUser := model.Payment{
		UserId:       userLogin.(user.UserMasuk).ID,
		ClinicId:     uint(IdUri.IdClinic),
		Status:       false,
		JenisHewan:   PayIn.JenisHewan,
		Keluhan:      PayIn.Keluhan,
		Ras:          PayIn.Ras,
		JenisKelamin: PayIn.JenisKelamin,
		Umur:         PayIn.Umur,
		Tanggal:      PayIn.Tanggal,
	}
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	db.Create(&paymentUser)
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Payment created",
		"data":    paymentUser,
	})
}
