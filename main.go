package main

import (
	"UnnecessaryMafia-Backend/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ser := server.Server{}
	ser.Run()
}
