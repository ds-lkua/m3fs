package model

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/open3fs/m3fs/pkg/errors"
)

// ConnectionArgs holds the connection arguments for the database.
type ConnectionArgs struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// DSN generates a Data Source Name (DSN) for the PostgreSQL connection.
func (arg *ConnectionArgs) DSN() string {
	newPart := func(k string, v any) string {
		return fmt.Sprintf("%s=%v ", k, v)
	}
	parts := []string{
		newPart("host", arg.Host),
		newPart("port", arg.Port),
		newPart("user", arg.User),
		newPart("password", arg.Password),
		newPart("dbname", arg.DBName),
		newPart("sslmode", "disable"),
	}

	return strings.Join(parts, " ")
}

// NewDB creates a new database connection using the provided connection arguments.
func NewDB(connArg *ConnectionArgs) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connArg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, errors.Annotate(err, "open database")
	}

	return db, nil
}
