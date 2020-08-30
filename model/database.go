package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDbConnection() (db *sql.DB) {
	/**
	create a db in mysql and change it with your info
	*/
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "shabanreza"
	dbName := "Mafia"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
