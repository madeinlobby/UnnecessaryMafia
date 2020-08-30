package controller

import (
	"UnnecessaryMafia-Backend/model"
	"database/sql"
	"log"
)

type GameUser struct {
	Username, password, phoneNumber, email, fname, lname, status string
}

func InsertUser(username, password, phoneNumber, email, fname, lname, status string) { //GameUser is the username of the table //todo hashing
	db := model.GetDbConnection()
	insForm, err := db.Prepare("INSERT INTO mafia.gameuser(username, password, `phone number`, email, fname, lname, status) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, _ = insForm.Exec(username, password, phoneNumber, email, fname, lname, status)
	log.Println("INSERT User: " + username)

	defer db.Close()
	return
}

func GetUser(username, password string) GameUser {
	db := model.GetDbConnection()
	var g GameUser
	err := db.QueryRow("SELECT mafia.gameuser.username, password, `phone number`, email, fname, lname, status FROM mafia.gameuser WHERE username = ? AND password = ?", username, password).
		Scan(&g.Username, &g.password, &g.phoneNumber, &g.email, &g.fname, &g.lname, &g.status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User Doesn't Exist Or Pass Is Wrong")
		} else {
			panic(err.Error())
		}
		defer db.Close()
		return GameUser{}
	}
	log.Println("SELECT User: " + username)
	log.Println(g)

	defer db.Close()
	return g
}
