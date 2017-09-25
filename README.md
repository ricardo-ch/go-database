# GO-DATABASE

<!-- [![wercker status](https://app.wercker.com/status/2bf9dccb9a12513dde0f54316c59a6b9/s/master "wercker status")](https://app.wercker.com/project/byKey/2bf9dccb9a12513dde0f54316c59a6b9)
[![Coverage Status](https://coveralls.io/repos/github/ricardo-ch/go-tracing/badge.svg?branch=master)](https://coveralls.io/github/ricardo-ch/go-tracing?branch=master) -->

Go-database allows to create connection object to database in easy way.

## Quick start


### MsSql implementation
```golang
// you need to import sql driver
import _ "github.com/denisenkom/go-mssqldb"
 
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
```

### Postgres implementation
```golang
// you need to import postgres driver
import _ "github.com/lib/pq"

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
Developer can implement in easy way his new dbprovider as he needed for his db. 
As example let's do it for sqlite
```golang

// you need to import sqlite driver
import _ "github.com/mattn/go-sqlite3"

// create your own dbprovide and override two methods
type SqlLiteConfig struct {
	DataSource string
	Version *int 
}

// DriverName
func (conf SqlLiteConfig) DriverName() string {
	return "sqlite3"
}

//Build method for MsSql
func (conf SqlLiteConfig) Build() string {
	var buffer bytes.Buffer

	// your code

	return buffer.String()
}
``` 

## License
go-database is licensed under the MIT license. (http://opensource.org/licenses/MIT)

## Contributing
Pull requests are the way to help us here. We will be really grateful.