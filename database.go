package database

import (
	"github.com/jmoiron/sqlx"
)

type ConnectionBuilder interface {
	DriverName() string
	Build() string
}

var sqlConnect = sqlx.Connect

func connect(driverName, connectionString string, advFunc func(*sqlx.DB)) (*sqlx.DB, error) {
	db, err := sqlConnect(driverName, connectionString)

	if err != nil {
		return nil, err
	}

	// execute function
	if advFunc != nil {
		advFunc(db)
	}

	return db.Unsafe(), nil
}

// Connect to DB
func Connect(builder ConnectionBuilder, advFunc func(*sqlx.DB)) (*sqlx.DB, error) {
	return connect(builder.DriverName(), builder.Build(), advFunc)
}
