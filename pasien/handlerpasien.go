package pasien

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pasien struct {
	NamaUser        string
	Alamat          string
	Contact         string
	JenisHewan      string
	Keluhan         string
	Ras             string
	JenisKelamin    string
	Umur            string
	Tanggal         string
	Jam             string
	Layanan         string
	StatusPelayanan bool
}

func GetPasien(c *gin.Context) {
	var pasien []Pasien
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	clinicLogin := c.MustGet("cliniclogin").(user.ClinicMasuk)
	var payment []model.Payment
	db.Where("clinic_id = ?", clinicLogin.ID).Find(&payment)
	for _, v := range payment {
		if v.Status {
			var user user.User
			db.Where("id = ?", v.UserId).Take(&user)
			pasien = append(pasien, Pasien{
				NamaUser:     user.Name,
				Alamat:       user.Address,
				Contact:      user.Contact,
				JenisHewan:   v.JenisHewan,
				Keluhan:      v.Keluhan,
				Ras:          v.Ras,
				JenisKelamin: v.JenisKelamin,
				Umur:         v.Umur,
				Jam:          v.Jam,
				Tanggal:      v.Tanggal,
				Layanan:      v.Layanan,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "See all patient success",
		"data":    pasien,
	})
}

type inputanUpdate struct {
	Status bool `json:"status" binding:"required"`
}

type returnStatus struct {
	IdPayment uint
	Status    bool
}

func UpdatePatient(c *gin.Context) {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	var uriPayment model.InputUriPayment
	c.BindUri(&uriPayment)
	var statusPasien user.StatusPasien
	if err := db.Where("id = ?", uint(uriPayment.IdPayment)).Take(&statusPasien).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Payment not found",
			"data":    "nil",
		})
	}
	var inputan inputanUpdate
	c.BindJSON(&inputan)
	statusPasien.Status = inputan.Status
	db.Save(&statusPasien)
	returnStatus := returnStatus{
		IdPayment: uint(uriPayment.IdPayment),
		Status:    inputan.Status,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Update status pasien success",
		"data":    returnStatus,
	})

}
