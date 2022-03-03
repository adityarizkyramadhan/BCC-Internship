package helper

import (
	"BCC-Internship/config"
	"BCC-Internship/user"
)

func SearchUserById(id int) interface{} {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	var userSearch user.User
	db.Where("id = ?", id).Take(&userSearch)
	result := user.UserMasuk{
		ID:      userSearch.ID,
		Name:    userSearch.Name,
		Contact: userSearch.Contact,
		Address: userSearch.Address,
	}
	return result
}

func SearchClinicById(id int) interface{} {
	db, err := config.InitializeDatabases()
	if err != nil {
		panic(err)
	}
	var userSearch user.Clinic
	db.Where("id = ?", id).Take(&userSearch)
	result := user.ClinicMasuk{
		ID:          userSearch.ID,
		NameClinic:  userSearch.NameClinic,
		Contact:     userSearch.Contact,
		Address:     userSearch.Address,
		SpreadSheet: userSearch.SpreadSheet,
	}
	return result
}
