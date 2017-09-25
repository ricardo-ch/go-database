package dbprovider

import "bytes"
import "fmt"
import "strconv"

type MsSqlConfig struct {
	Server            *string
	Port              *int
	FailoverPartner   *string
	FailoverPort      *int
	UserID            *string
	Password          string
	Database          string
	ConnectionTimeout *int
	Log               *byte
	AppName           *string
}

// DriverName
func (conf MsSqlConfig) DriverName() string {
	return "mssql"
}

//Build method for MsSql
func (conf MsSqlConfig) Build() string {
	const formatParam = "%s=%s;"
	var buffer bytes.Buffer

	if conf.Server != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "server", *conf.Server))
	}

	if conf.Port != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "port", strconv.Itoa(*conf.Port)))
	}

	if conf.FailoverPartner != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "failoverpartner", *conf.FailoverPartner))
	}

	if conf.FailoverPort != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "failoverport", strconv.Itoa(*conf.FailoverPort)))
	}

	if conf.UserID != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "user id", *conf.UserID))
	}

	if conf.Password != "" {
		buffer.WriteString(fmt.Sprintf(formatParam, "password", conf.Password))
	}

	if conf.Database != "" {
		buffer.WriteString(fmt.Sprintf(formatParam, "database", conf.Database))
	}

	if conf.ConnectionTimeout != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "connection timeout", strconv.Itoa(*conf.ConnectionTimeout)))
	}

	if conf.Log != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "log", strconv.Itoa(int(*conf.Log))))
	}

	if conf.AppName != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "app name", *conf.AppName))
	}

	return buffer.String()
}
