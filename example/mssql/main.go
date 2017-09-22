package main

import (
	"fmt"
	"log"

	db "github.com/ricardo-ch/go-database"
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

	dbase, err := db.ConnectToMsSQL("DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", 1, map[db.StringParameter]string{
		db.Password: "DB_PWD",
		//db.MsSQLLog: "1",
	})

	if err != nil {
		fmt.Println(err)
	}

	// close connection after using
	defer func() {
		if dbase != nil {
			dbase.Close()
		}
	}()

	dbase.MustExec(schema)

	tx := dbase.MustBegin()
	tx.MustExec("INSERT INTO dbo.person (first_name, last_name, email) VALUES ($1, $2, $3)", "Andrii", "Lesch", "andriilesch@gmail.net")
	tx.MustExec("INSERT INTO dbo.person (first_name, last_name, email) VALUES ($1, $2, $3)", "Yannick", "Devos", "yannickdevos@gmail.net")
	tx.MustExec("INSERT INTO dbo.person (first_name, last_name, email) VALUES ($1, $2, $3)", "Francois", "Poinsot", "francoispoinsot@gmail.net")
	tx.Commit()

	users := []Person{}
	dbase.Select(&users, "SELECT * FROM dbo.person ORDER BY first_name ASC")
	andrii, yannick, francois := users[0], users[1], users[2]

	fmt.Printf("%#v\n%#v\n%#v\n", andrii, yannick, francois)

	_, err = dbase.Exec("DROP TABLE dbo.person;")
	if err != nil {
		log.Fatalln(err)
	}
}
