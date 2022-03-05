package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	userLogin := c.MustGet("userlogin").(user.UserMasuk)
	var history []model.Payment
	db.Where("user_id = ?", userLogin.ID).Find(&history)
	var printHistory []model.GetPayment
	for _, v := range history {
		printHistory = append(printHistory, model.GetPayment{
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
		"message": "See history payment success",
		"data":    printHistory,
	})

}
