# GO-DATABASE

<!-- [![wercker status](https://app.wercker.com/status/2bf9dccb9a12513dde0f54316c59a6b9/s/master "wercker status")](https://app.wercker.com/project/byKey/2bf9dccb9a12513dde0f54316c59a6b9)
[![Coverage Status](https://coveralls.io/repos/github/ricardo-ch/go-tracing/badge.svg?branch=master)](https://coveralls.io/github/ricardo-ch/go-tracing?branch=master) -->

Go-database allows to create connection object to database in easy way.

## Quick start

```golang

// connect to MsSql Database
msSqlConf := dbprovider.MsSqlConfig{
		Database: "YOUR_DBNAME",
		Port:     "YOUR_PORT",
		Server:   "YOUR_HOST",
		UserID:   "YOUR_USER",
		Password: "YOUR_PWD",
}

dbase, _ := db.Connect(msSqlConf, 1)

// close connection after using
defer func() {
	if dbase != nil {
		fmt.Println("Db closed")
		dbase.Close()
	}
}()


// connect to Postgres Database
pqConf := dbprovider.PostgresConfig{
		Host:     "YOUR_HOST",
		Port:     "YOUR_PORT",
		Database: "YOUR_DBNAME",
		UserID:   "YOUR_USER",
		SslMode:  "YOUR_SSLMODE", // 
		Password: "YOUR_PWD",
}

dbase, _ := db.Connect(pqConf, 1)

// close connection after using
defer func() {
	if dbase != nil {
		fmt.Println("Db closed")
		dbase.Close()
	}
}()

```
## Features


## License
go-database is licensed under the MIT license. (http://opensource.org/licenses/MIT)

## Contributing
Pull requests are the way to help us here. We will be really grateful.