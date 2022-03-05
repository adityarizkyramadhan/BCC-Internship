package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SeeValidatePayment(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
		return
	}
	clinicLogin := c.MustGet("cliniclogin").(user.ClinicMasuk)
	var payment []model.Payment
	if err := db.Preload("SaveImage").Find(&payment, "clinic_id = ?", clinicLogin.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when getting payment",
			"data":    err.Error(),
		})
		return
	}
	var getPayment []model.GetPayment
	for _, v := range payment {
		if v.Status {
			getPayment = append(getPayment, model.GetPayment{
				IDTransaction: v.ID,
				UserId:        v.UserId,
				ClinicId:      v.ClinicId,
				Status:        v.Status,
				JenisHewan:    v.JenisHewan,
				Keluhan:       v.Keluhan,
				Ras:           v.Ras,
				JenisKelamin:  v.JenisKelamin,
				Umur:          v.Umur,
				Tanggal:       v.Tanggal,
				Layanan:       v.Layanan,
				Harga:         v.Harga,
				SaveImage:     v.SaveImage,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "See all payment success",
		"data":    getPayment,
	})
}

func ShowInvoices(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when initializing databases",
			"data":    err.Error(),
		})
		return
	}
	clinicLogin := c.MustGet("cliniclogin").(user.ClinicMasuk)
	var payment []model.Payment
	if err := db.Preload("SaveImage").Find(&payment, "clinic_id = ?", clinicLogin.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Payment not found",
			"data":    err.Error(),
		})
		return
	}
	var getPayment []model.DataInvoice
	for _, v := range payment {
		if v.Status {
			getPayment = append(getPayment, model.DataInvoice{
				IDTransaction: v.ID,
				UserId:        v.UserId,
				ClinicId:      v.ClinicId,
				Status:        v.Status,
			})
		}
	}
	invoice := model.ReturnInvoice{
		BanyakPayment: len(getPayment),
		TotalBayar:    len(getPayment) * 4000,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get invoices success",
		"data":    invoice,
	})
}

func GetAllPaymentClinic(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	clinicLogin := c.MustGet("cliniclogin").(user.ClinicMasuk)
	var payment []model.Payment
	if err := db.Preload("SaveImage").Find(&payment, "clinic_id = ?", clinicLogin.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Internal server error",
			"data":    err.Error(),
		})
		return
	}
	var getPayment []model.GetPayment
	for _, v := range payment {
		getPayment = append(getPayment, model.GetPayment{
			IDTransaction: v.ID,
			UserId:        v.UserId,
			ClinicId:      v.ClinicId,
			Status:        v.Status,
			JenisHewan:    v.JenisHewan,
			Keluhan:       v.Keluhan,
			Ras:           v.Ras,
			JenisKelamin:  v.JenisKelamin,
			Umur:          v.Umur,
			Tanggal:       v.Tanggal,
			Layanan:       v.Layanan,
			Harga:         v.Harga,
			SaveImage:     v.SaveImage,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "See all payment success",
		"data":    getPayment,
	})
}

func UpdatePayment(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	clinicLogin := c.MustGet("cliniclogin").(user.ClinicMasuk)
	var payment model.Payment
	var uri model.InputUriPayment
	if err := c.BindUri(&uri); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	if err := db.Where("id = ?", uri.IdPayment).Take(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Payment not found",
			"data":    err.Error(),
		})
		return
	}
	if payment.ClinicId != clinicLogin.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"data":    "You are not authorized to update this payment",
		})
	}
	var input model.InputStatusPayment
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	payment.Status = input.Status
	if err := db.Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Internal server error",
			"data":    err.Error(),
		})
		return
	}
	sendData := model.GetPayment{
		IDTransaction: payment.ID,
		UserId:        payment.UserId,
		ClinicId:      payment.ClinicId,
		Status:        payment.Status,
		JenisHewan:    payment.JenisHewan,
		Keluhan:       payment.Keluhan,
		Ras:           payment.Ras,
		JenisKelamin:  payment.JenisKelamin,
		Umur:          payment.Umur,
		Tanggal:       payment.Tanggal,
		Layanan:       payment.Layanan,
		Harga:         payment.Harga,
		SaveImage:     payment.SaveImage,
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Validate payment success",
		"data":    sendData,
	})
}
