package handler

import (
	"BCC-Internship/config"
	"BCC-Internship/model"
	"BCC-Internship/tokengenerator"
	"BCC-Internship/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UploadClinicImage(c *gin.Context) {
	// upload file by gonic/gin
	file, err := c.FormFile("fotoklinik")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	var idTrx model.InputUriClinic
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
	path := fmt.Sprintf("imageClinics/foto%d-%s", idTrx.IdClinic, file.Filename)
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
	image := user.ImageClinic{
		ClinicID: uint(idTrx.IdClinic),
		Path:     path,
	}
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	db.Create(&image)
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "File uploaded",
		"data":    image,
	})
}

func SearchClinic(c *gin.Context) {
	querry, isKotaExist := c.GetQuery("kota")
	if !isKotaExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Kota is required",
			"data":    nil,
		})
		return
	}
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	var body []user.Clinic
	if db.Preload("ImageClinic").Where("address LIKE ?", "%"+querry+"%").Find(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when querrying database",
			"data":    err.Error(),
		})
		return
	}
	var printClinic []user.PrintClinic
	for i := range body {
		printClinic = append(printClinic, user.PrintClinic{
			ID:          body[i].ID,
			NameClinic:  body[i].NameClinic,
			Address:     body[i].Address,
			Contact:     body[i].Contact,
			SpreadSheet: body[i].SpreadSheet,
			NoRekening:  body[i].NoRekening,
			AtasNama:    body[i].AtasNama,
			ImageClinic: body[i].ImageClinic,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Status OK",
		"message": "Clinic found",
		"data":    printClinic,
	})
}

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
	if err := db.Preload("ImageClinic").Find(&clinic).Error; err != nil {
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
			NoRekening:  clinic[i].NoRekening,
			AtasNama:    clinic[i].AtasNama,
			ImageClinic: clinic[i].ImageClinic,
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
	err = c.BindJSON(&body)
	if err != nil {
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
		NoRekening:     body.NoRekening,
		AtasNama:       body.AtasNama,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input",
			"data":    err.Error(),
		})
		return
	}
	var clinicSama user.Clinic
	if err := db.Where("username_clinic = ?", body.UsernameClinic).First(&clinicSama).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Error when querrying clinic",
			"data":    err.Error(),
		})
		return
	}
	if clinicSama.UsernameClinic == clinic.UsernameClinic {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "Conflict",
			"message": "Username already exist",
			"data":    "Error because username already exists",
		})
		return
	}
	if err := db.Create(&clinic).Error; err != nil {
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
		NoRekening:  clinic.NoRekening,
		Token:       token,
		AtasNama:    clinic.AtasNama,
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
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Bad Request",
			"message": "Error when binding JSON",
			"data":    err.Error(),
		})
	}
	var clinic user.Clinic
	if err := db.Where("username_clinic = ?", body.Username).Take(&clinic).Error; err != nil {
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
		NoRekening:  clinic.NoRekening,
		AtasNama:    clinic.AtasNama,
		Token:       token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Status OK",
		"message": "Clinic logged in",
		"data":    clinicReturn,
	})
}
