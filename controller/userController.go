package controller

import (
	"UnnecessaryMafia-Backend/model"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func InsertUser(username, password, phoneNumber, email, fname, lname, status string) { //GameUser is the username of the table
	//TODO authorize bayad she ba email
	db := model.GetDbConnection()
	hashedPass := hashPassword(password)
	insForm, err := db.Prepare("INSERT INTO mafia.gameuser(username, password, `phone number`, email, fname, lname, status) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, _ = insForm.Exec(username, hashedPass, phoneNumber, email, fname, lname, status)
	log.Println("INSERT User: " + username)

	defer db.Close()
	return
}

func GetUser(username, password string) model.GameUser {
	db := model.GetDbConnection()
	var g model.GameUser
	//err := db.QueryRow("SELECT mafia.gameuser.username, password, `phone number`, email, fname, lname, status FROM mafia.gameuser WHERE username = ? AND password = ?", username, password).
	err := db.QueryRow("SELECT mafia.gameuser.username, password, `phone number`, email, fname, lname, status FROM mafia.gameuser WHERE username = ?", username).
		Scan(&g.Username, &g.Password, &g.PhoneNumber, &g.Email, &g.Fname, &g.Lname, &g.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User Doesn't Exist")
		} else {
			panic(err.Error())
		}
		defer db.Close()
		return model.GameUser{}
	}
	if !checkPassword(password, g.Password) {
		log.Println("Pass is not correct.")
		defer db.Close()
		return model.GameUser{}
	}

	log.Println("SELECT User: " + username)
	log.Println(g)

	defer db.Close()
	return g
}

/**
Hashes password to save in db
*/
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	return string(hash)
}

/**
Is (password -> hash) same with hashed one in db
*/
func checkPassword(password, hashFromDatabase string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(password)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	return true
}
