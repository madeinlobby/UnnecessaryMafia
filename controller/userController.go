package controller

import (
	"UnnecessaryMafia-Backend/model"
	"log"
)

func InsertUser(name, password string) { //GameUser is the name of the table //todo hashing
	db := model.GetDbConnection()
	insForm, err := db.Prepare("INSERT INTO GameUser(name, password) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, _ = insForm.Exec(name, password)
	log.Println("INSERT: Name: " + name + " | Password: " + password)

	defer db.Close()
	return
}
