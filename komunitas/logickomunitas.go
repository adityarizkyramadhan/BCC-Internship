package komunitas

import (
	"BCC-Internship/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type komunitas struct {
	gorm.Model
	NamaKomunitas string
	JenisHewan    string
	Deskripsi     string
	AsalKota      string
	LinkKomunitas string
}

type komunitasInput struct {
	NamaKomunitas string `json:"namaKomunitas" binding:"required"`
	JenisHewan    string `json:"jenisHewan" binding:"required"`
	Deskripsi     string `json:"deskripsi" binding:"required"`
	AsalKota      string `json:"asalKota" binding:"required"`
	LinkKomunitas string `json:"linkKomunitas" binding:"required"`
}

type returnKomunitas struct {
	ID            uint
	NamaKomunitas string
	JenisHewan    string
	Deskripsi     string
	AsalKota      string
	LinkKomunitas string
}

func AddKomunitas(c *gin.Context) {
	var komunitasIn komunitasInput
	if err := c.BindJSON(&komunitasIn); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Status Unprocessable Entity",
			"message": "Error when bind JSON",
			"data":    err.Error(),
		})
		return
	}
	komunitas := komunitas{
		NamaKomunitas: komunitasIn.NamaKomunitas,
		JenisHewan:    komunitasIn.JenisHewan,
		Deskripsi:     komunitasIn.Deskripsi,
		AsalKota:      komunitasIn.AsalKota,
		LinkKomunitas: komunitasIn.LinkKomunitas,
	}
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	db.Create(&komunitas)
	retKom := returnKomunitas{
		ID:            komunitas.ID,
		NamaKomunitas: komunitas.NamaKomunitas,
		JenisHewan:    komunitas.JenisHewan,
		Deskripsi:     komunitas.Deskripsi,
		AsalKota:      komunitas.AsalKota,
		LinkKomunitas: komunitas.LinkKomunitas,
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Created",
		"message": "Community created successfully",
		"data":    retKom,
	})
}
func GetKomunitas(c *gin.Context) {
	var komunitas []komunitas
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	db.Find(&komunitas)
	var retKomunitas returnKomunitas
	for _, kom := range komunitas {
		retKomunitas = returnKomunitas{
			ID:            kom.ID,
			NamaKomunitas: kom.NamaKomunitas,
			JenisHewan:    kom.JenisHewan,
			Deskripsi:     kom.Deskripsi,
			AsalKota:      kom.AsalKota,
			LinkKomunitas: kom.LinkKomunitas,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Sucessfully get all community",
		"data":    retKomunitas,
	})
}

func SearchKomunitas(c *gin.Context) {
	var komunitas []komunitas
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	jenisHewan, isHewanExist := c.GetQuery("jenisHewan")
	asalKota, isKotaExist := c.GetQuery("asalKota")

	if !isHewanExist && !isKotaExist {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Status Unprocessable Entity",
			"message": "Error when GetQuery URI",
			"data":    "URI unprocessable entity",
		})
		return
	}
	if isHewanExist {
		db.Where("jenis_hewan LIKE ?", "%"+jenisHewan+"%")
	}
	if isKotaExist {
		db.Where("asal_kota LIKE ?", "%"+asalKota+"%")
	}
	if db.Find(&komunitas).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Community not found",
			"data":    "Not found",
		})
	}
	var retKom returnKomunitas
	for _, kom := range komunitas {
		retKom = returnKomunitas{
			ID:            kom.ID,
			NamaKomunitas: kom.NamaKomunitas,
			JenisHewan:    kom.JenisHewan,
			Deskripsi:     kom.Deskripsi,
			AsalKota:      kom.AsalKota,
			LinkKomunitas: kom.LinkKomunitas,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Comunity found",
		"data":    retKom,
	})
}
