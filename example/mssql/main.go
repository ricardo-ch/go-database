package main

import (
	"fmt"

	// Import Microsoft SQL driver
	_ "github.com/denisenkom/go-mssqldb"
	db "github.com/ricardo-ch/go-database"
	"github.com/ricardo-ch/go-database/dbprovider"
)

var schema = `
CREATE TABLE dbo.person (
	first_name text,
	last_name text,
	email text
); `

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

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

	// 	dbase.MustExec(schema)

	// 	tx := dbase.MustBegin()
	// 	tx.MustExec("INSERT INTO dbo.person (first_name, last_name, email) VALUES ($1, $2, $3)", "Andrii", "Lesch", "andriilesch@gmail.net")
	// 	tx.MustExec("INSERT INTO dbo.person (first_name, last_name, email) VALUES ($1, $2, $3)", "Yannick", "Devos", "yannickdevos@gmail.net")
	// 	tx.MustExec("INSERT INTO dbo.person (first_name, last_name, email) VALUES ($1, $2, $3)", "Francois", "Poinsot", "francoispoinsot@gmail.net")
	// 	tx.Commit()

	// 	users := []Person{}
	// 	dbase.Select(&users, "SELECT * FROM dbo.person ORDER BY first_name ASC")
	// 	andrii, yannick, francois := users[0], users[1], users[2]

	// 	fmt.Printf("%#v\n%#v\n%#v\n", andrii, yannick, francois)

	// 	_, err = dbase.Exec("DROP TABLE dbo.person;")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
}
