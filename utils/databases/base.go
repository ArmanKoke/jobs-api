package databases

import (
	"database/sql"
	"fmt"
	"jobs-api/utils/env"

	"github.com/pkg/errors"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

// DBEnv interface
type DBEnv interface {
	PostgresConnReader() string
	PostgresConnWriter() string
	PostgresDBReader() *sql.DB
	PostgresDBWriter() *sql.DB
	PostgresDriver() string
}

// NewDbEnv instance create
func NewDbEnv() (DBEnv, error) {
	driver := new(dbDriver)
	driver.postgresConnReader = env.AppPostgresConn(env.Reader)
	driver.postgresConnWriter = env.AppPostgresConn(env.Writer)
	driver.postgresDriver = env.AppPostgresDriver()

	msDbReader, err := Connect(driver.postgresDriver, driver.postgresConnReader)
	if err != nil {
		return nil, err
	}

	msDbWriter, err := Connect(driver.postgresDriver, driver.postgresConnWriter)
	if err != nil {
		return nil, err
	}

	driver.postgresDBReader = msDbReader
	driver.postgresDBWriter = msDbWriter

	return driver, err
}

// Connect to db
func Connect(driver, source string) (*sql.DB, error) {
	_db, err := apmsql.Open(driver, source)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Connection error (Driver: %s, Source: %s)", driver, source))
	}

	err = _db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Connection ping error (Driver: %s, Source: %s)", driver, source))
	}

	_db.SetMaxOpenConns(10)
	_db.SetConnMaxLifetime(0)

	return _db, nil
}
