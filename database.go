package database

import "github.com/jmoiron/sqlx"

type ConnectionBuilder interface {
	DriverName() string
	Build() string
}

// var sqlConnect *sqlx.DB
var sqlConnect = sqlx.Connect
var setMaxIdleConns = func(db *sqlx.DB, maxIdleDBConnections int) {
	db.SetMaxIdleConns(maxIdleDBConnections)
}

func connect(driverName, connectionString string, maxIdleDBConnections int) (*sqlx.DB, error) {
	db, err := sqlConnect(driverName, connectionString)
	if err != nil {
		return nil, err
	}

	setMaxIdleConns(db, maxIdleDBConnections)
	return db.Unsafe(), nil
}

// Connect to DB
func Connect(builder ConnectionBuilder, maxIdleDBConnections int) (*sqlx.DB, error) {
	return connect(builder.DriverName(), builder.Build(), maxIdleDBConnections)
}
