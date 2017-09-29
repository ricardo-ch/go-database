# GO-DATABASE

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/ricardo-ch/go-database/blob/master/LICENSE)

Go-database is a library which allows to create connection object to database in easy way. it's using `sqlx` extensions on go's standard `database/sql` library.  

## install
```
go get github.com/ricardo-ch/go-database
```

## usage
Below is an example you can find MsSql implementation.

```golang
// you need to import sql driver
import _ "github.com/denisenkom/go-mssqldb"
 
// set your credentials for MsSql Database
msSqlConf := dbprovider.MsSqlConfig{
		Database: "YOUR_DBNAME",
		Port:     "YOUR_PORT",
		Server:   "YOUR_HOST",
		UserID:   "YOUR_USER",
		Password: "YOUR_PWD",
}

// in case if you need to use default parametes
dbase, _ := db.Connect(msSqlConf, nil)

//in case for additional configuration
// dbase, _ := db.Connect(msSqlConf, func(db *sqlx.DB){
//	db.SetMaxIdleConns(5)
// })

// close connection after using
defer func() {
	if dbase != nil {
		fmt.Println("Db closed")
		dbase.Close()
	}
}()
```

## License
go-database is licensed under the MIT license. (http://opensource.org/licenses/MIT)

## Contributing
Pull requests are the way to help us here. We will be really grateful.