package main

import (
	"fmt"
	"log"

	//Import postgres driver
	_ "github.com/lib/pq"
	db "github.com/ricardo-ch/go-database"
	"github.com/ricardo-ch/go-database/dbprovider"
)

var schema = `
CREATE TABLE person (
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
	var sslmode = "disable" // SSLMODE

	pqConf := dbprovider.PostgresConfig{
		Host:     &dbHost,
		Port:     &dbPort,
		Database: dbName,
		UserID:   dbUser,
		SslMode:  sslmode,
		Password: dbPwd,
	}

	dbase, err := db.Connect(pqConf, -1)
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

	// 	sqlx.
	// 		db.MustExec(schema)

	// 	tx := db.MustBegin()
	// 	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Andrii", "Lesch", "andriilesch@gmail.net")
	// 	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Yannick", "Devos", "yannickdevos@gmail.net")
	// 	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Francois", "Poinsot", "francoispoinsot@gmail.net")
	// 	tx.Commit()

	// 	persons := []Person{}
	// 	db.Select(&persons, "SELECT * FROM person ORDER BY first_name ASC")
	// 	andrii, yannick, francois := persons[0], persons[1], persons[2]

	// 	fmt.Printf("%#v\n%#v\n%#v\n", andrii, yannick, francois)

	// 	_, err = db.Exec("DROP TABLE IF EXISTS public.person;")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
}
