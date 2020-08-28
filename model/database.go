package model

import "database/sql"

func GetDbConnection() (db *sql.DB) {
	/**
	create a db in mysql and change it with your info
	*/
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "your pass"
	dbName := "Mafia"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
