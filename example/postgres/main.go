package main

import (
	"fmt"
	"log"

	dbase "github.com/ricardo-ch/go-database"
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

	db, err := dbase.ConnectToPostgres("DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "sslmode", -1, map[dbase.StringParameter]string{
		dbase.Password: "DB_PWD",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// close connection after using
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Andrii", "Lesch", "andriilesch@gmail.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Yannick", "Devos", "yannickdevos@gmail.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Francois", "Poinsot", "francoispoinsot@gmail.net")
	tx.Commit()

	persons := []Person{}
	db.Select(&persons, "SELECT * FROM person ORDER BY first_name ASC")
	andrii, yannick, francois := persons[0], persons[1], persons[2]

	fmt.Printf("%#v\n%#v\n%#v\n", andrii, yannick, francois)

	_, err = db.Exec("DROP TABLE IF EXISTS public.person;")
	if err != nil {
		log.Fatalln(err)
	}

}
