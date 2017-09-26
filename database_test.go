package database

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/ricardo-ch/go-database/dbprovider"
	"github.com/stretchr/testify/assert"
)

var expectedErrorMessage = errors.New("fail DB connection")

var connectTests = [2]struct {
	description   string
	connect       func(driverName, dataSourceName string) (*sqlx.DB, error)
	expectedError error
}{
	{
		description: "Connection succesfully Done",
		connect: func(driverName, dataSourceName string) (*sqlx.DB, error) {
			return &sqlx.DB{}, nil
		},
		expectedError: nil,
	},
	{
		description: "Connection failed",
		connect: func(driverName, dataSourceName string) (*sqlx.DB, error) {
			return nil, expectedErrorMessage
		},
		expectedError: expectedErrorMessage,
	},
}

func TestConnect(t *testing.T) {

	setMaxIdleConns = func(db *sqlx.DB, maxIdleDBConnections int) {
	}

	var dbHost = "YOUR_HOST"
	var dbPort = 0
	var dbName = "YOUR_DBNAME"
	var dbUser = "YOUR_USER"
	var dbPwd = "YOUR_PWD"

	mssqlConf := dbprovider.MsSqlConfig{
		Database: dbName,
		Password: dbPwd,
		Port:     &dbPort,
		Server:   &dbHost,
		UserID:   &dbUser,
	}
	for _, test := range connectTests {
		//Arrange
		sqlConnect = test.connect

		//Act
		_, err := Connect(mssqlConf, -1)

		assert.Equal(t, test.expectedError, err, test.description)
	}
}
