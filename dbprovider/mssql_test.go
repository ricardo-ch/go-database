package dbprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbHostTest = "TEST_HOST"
var dbPortTest = 12
var dbNameTest = "TEST_DBNAME"
var dbUserTest = "TEST_USER"
var dbPwdTest = "TEST_PWD"
var appNameTest = "TEST-APP-NAME"
var logTest byte = 8

var connectMsSqlTests = [4]struct {
	//description       string
	connectionBuilder MsSqlConfig
	expectedResult    string
}{
	{
		connectionBuilder: MsSqlConfig{
			Server:   &dbHostTest,
			Port:     &dbPortTest,
			Database: dbNameTest,
			UserID:   &dbUserTest,
			Password: dbPwdTest,
			AppName:  &appNameTest,
			Log:      &logTest,
		},
		expectedResult: "server=TEST_HOST;port=12;user id=TEST_USER;password=TEST_PWD;database=TEST_DBNAME;log=8;app name=TEST-APP-NAME;",
	},
	{
		connectionBuilder: MsSqlConfig{
			Server: &dbHostTest,
			Log:    &logTest,
		},
		expectedResult: "server=TEST_HOST;log=8;",
	},
	{
		connectionBuilder: MsSqlConfig{},
		expectedResult:    "",
	},
	{
		connectionBuilder: MsSqlConfig{
			Database: dbNameTest,
			Password: dbPwdTest,
		},
		expectedResult: "password=TEST_PWD;database=TEST_DBNAME;",
	},
}

func Test_MsSqlConfig_ConnectionBuilder(t *testing.T) {
	for _, test := range connectMsSqlTests {
		//Arrange
		builder := test.connectionBuilder
		//Act
		result := builder.Build()

		assert.Equal(t, test.expectedResult, result)
	}
}
