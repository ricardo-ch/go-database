package main

import (
	"fmt"
	"log"

	//Import postgres driver
	_ "github.com/lib/pq"
	db "github.com/ricardo-ch/go-database"
	"github.com/ricardo-ch/go-database/dbprovider"
)

func main() {

	var dbHost = "YOUR_HOST"
	var dbPort = 0 //YOUR_PORT
	var dbName = "YOUR_DBNAME"
	var dbUser = "YOUR_USER"
	var dbPwd = "YOUR_PWD"
	var sslmode = "disable" // SSLMODE

	pqConf := dbprovider.PostgresConfig{
		Host:     &dbHost,
		Port:     &dbPort,
		Database: dbName,
		UserID:   dbUser,
		SslMode:  sslmode,
		Password: dbPwd,
	}

	dbase, err := db.Connect(pqConf, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// close connection after using
	defer func() {
		if dbase != nil {
			fmt.Println("Db closed")
			dbase.Close()
		}
	}()
}
