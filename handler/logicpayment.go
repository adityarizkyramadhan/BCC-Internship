package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	userLogin := c.MustGet("userlogin").(user.UserMasuk)
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
	fmt.Println(userLogin)
	paymentUser := model.Payment{
		UserId:       userLogin.ID,
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
	sendPayment := model.PaymentReturn{
		IDTransaction: paymentUser.ID,
		UserId:        paymentUser.UserId,
		ClinicId:      paymentUser.ClinicId,
		Status:        paymentUser.Status,
		JenisHewan:    paymentUser.JenisHewan,
		Keluhan:       paymentUser.Keluhan,
		Ras:           paymentUser.Ras,
		JenisKelamin:  paymentUser.JenisKelamin,
		Umur:          paymentUser.Umur,
		Tanggal:       paymentUser.Tanggal,
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Payment created",
		"data":    sendPayment,
	})
}

func UploadStructPayment(c *gin.Context) {
	// upload file by gonic/gin
	file, err := c.FormFile("transaction")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	var idTrx model.InputUriTransaction
	err = c.BindUri(&idTrx)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	// get file name
	path := fmt.Sprintf("image/foto%d-%s", idTrx.IdTransaction, file.Filename)
	// save file to folder
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	image := model.SaveImage{
		TransactionID: uint(idTrx.IdTransaction),
		Path:          path,
	}
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	db.Create(&image)
	returnImage := model.ReturnImage{
		TransactionID: image.TransactionID,
		Path:          image.Path,
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "File uploaded",
		"data":    returnImage,
	})
}
