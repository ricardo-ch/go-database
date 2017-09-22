package database

import (
	"bytes"
	"fmt"

	// Import Microsoft SQL driver
	_ "github.com/denisenkom/go-mssqldb"
	//We need it to use postgres driver
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var parameter = []string{

	"password",

	//mssql
	"failoverpartner",
	"failoverport",
	"connection timeout",
	"app name",
	"log",

	//pq
	"connect_timeout",

	//sqlLite
}

type Enum interface {
	name() string
}

type StringParameter int

func (e StringParameter) name() string {
	return parameter[e]
}

// ConnectionParameter
type ConnectionParameter struct {
	StringParameter StringParameter
	Value           string
}

const (
	Password StringParameter = iota

	//mssql
	MsSQLFailoverPartner
	MsSQLFailoverPort
	MsSQLConnectionTimeout
	MsSQLAppName
	MsSQLLog

	//postgres
	PqConnectTimeout

	//sqlite
	// will be implemented
)

const (
	msSQLDriverName    = "mssql"
	postgresDriverName = "postgres"
	//sqliteDriverName   = "sqlite"
)

func buildConnectionString(connParameters map[StringParameter]string, separator string) string {
	var buffer bytes.Buffer

	for index, value := range connParameters {
		buffer.WriteString(fmt.Sprintf("%s=%s%s", index.name(), value, separator))
	}

	return buffer.String()
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

// ConnectToMsSQL - create connection for MsSql
func ConnectToMsSQL(dbServer, dbPort, dbName, dbUser string, maxIdleDBConnections int, connectionParameters map[StringParameter]string) (*sqlx.DB, error) {

	//build connection string
	connParams := buildConnectionString(connectionParameters, ";")
	connString := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;%s", dbServer, dbPort, dbName, dbUser, connParams)

	return connect(msSQLDriverName, connString, maxIdleDBConnections)
}

// ConnectToPostgres - create connection for Postgres
func ConnectToPostgres(dbHost, dbPort, dbName, dbUser, sslmode string, maxIdleDBConnections int, connectionParameters map[StringParameter]string) (*sqlx.DB, error) {

	connParams := buildConnectionString(connectionParameters, " ")
	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=%s %s", dbHost, dbPort, dbName, dbUser, sslmode, connParams)

	return connect(postgresDriverName, connString, maxIdleDBConnections)
}
