package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPaymentClinic(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	clinicLogin := c.MustGet("cliniclogin").(user.ClinicMasuk)
	var payment []model.Payment
	if err := db.Find(&payment, "clinic_id = ?", clinicLogin.ID).Error; err != nil {
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
			SaveImage:     v.SaveImage,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "See all payment success",
		"data":    getPayment,
	})
}
