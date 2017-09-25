package main

import (
	"fmt"

	// Import Microsoft SQL driver
	_ "github.com/denisenkom/go-mssqldb"
	db "github.com/ricardo-ch/go-database"
	"github.com/ricardo-ch/go-database/dbprovider"
)

func main() {

	var dbHost = "YOUR_HOST"
	var dbPort = 0 //YOUR_PORT
	var dbName = "YOUR_DBNAME"
	var dbUser = "YOUR_USER"
	var dbPwd = "YOUR_PWD"
	var Log byte = 8

	msSqlConf := dbprovider.MsSqlConfig{
		Database: dbName,
		Password: dbPwd,
		Port:     &dbPort,
		Server:   &dbHost,
		UserID:   &dbUser,
		Log:      &Log,
	}

	dbase, err := db.Connect(msSqlConf, 1)
	if err != nil {
		fmt.Println(err)
	}

	// close connection after using
	defer func() {
		if dbase != nil {
			fmt.Println("Db closed")
			dbase.Close()
		}
	}()

}
