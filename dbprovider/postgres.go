package dbprovider

import (
	"bytes"
	"fmt"
	"strconv"
)

type PostgresConfig struct {
	Database          string
	UserID            string
	Password          string
	Host              *string
	Port              *int
	SslMode           string
	ConnectionTimeout *int
}

// DriverName return driver db
func (conf PostgresConfig) DriverName() string {
	return "postgres"
}

//Build method for Postgres
func (conf PostgresConfig) Build() string {
	const formatParam = "%s=%s "
	var buffer bytes.Buffer

	if conf.Database != "" {
		buffer.WriteString(fmt.Sprintf(formatParam, "dbname", conf.Database))
	}

	if conf.UserID != "" {
		buffer.WriteString(fmt.Sprintf(formatParam, "user", conf.UserID))
	}

	if conf.Password != "" {
		buffer.WriteString(fmt.Sprintf(formatParam, "password", conf.Password))
	}

	if conf.Host != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "host", *conf.Host))
	}

	if conf.Port != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "port", strconv.Itoa(*conf.Port)))
	}

	if conf.SslMode != "" {
		buffer.WriteString(fmt.Sprintf(formatParam, "sslmode", conf.SslMode))
	}

	if conf.ConnectionTimeout != nil {
		buffer.WriteString(fmt.Sprintf(formatParam, "connect_timeout", strconv.Itoa(*conf.ConnectionTimeout)))
	}

	return buffer.String()
}
