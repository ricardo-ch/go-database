package dbprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var connectPostgresTests = [4]struct {
	connectionBuilder PostgresConfig
	expectedResult    string
}{
	{
		connectionBuilder: PostgresConfig{
			Database: "database",
			Host:     &dbHostTest,
			Password: dbPwdTest,
			Port:     &dbPortTest,
			SslMode:  "disable",
			UserID:   dbUserTest,
		},
		expectedResult: "dbname=database user=TEST_USER password=TEST_PWD host=TEST_HOST port=12 sslmode=disable ",
	},
	{
		connectionBuilder: PostgresConfig{
			Database: "database",
			SslMode:  "disable",
			UserID:   dbUserTest,
		},
		expectedResult: "dbname=database user=TEST_USER sslmode=disable ",
	},
	{
		connectionBuilder: PostgresConfig{},
		expectedResult:    "",
	},
	{
		connectionBuilder: PostgresConfig{
			Database: "database",
			Host:     &dbHostTest,
			Port:     &dbPortTest,
			SslMode:  "disable",
			UserID:   dbUserTest,
		},
		expectedResult: "dbname=database user=TEST_USER host=TEST_HOST port=12 sslmode=disable ",
	},
}

func Test_PostgresConfig_ConnectionBuilder(t *testing.T) {
	for _, test := range connectPostgresTests {
		//Arrange
		builder := test.connectionBuilder
		//Act
		result := builder.Build()

		assert.Equal(t, test.expectedResult, result)
	}
}
