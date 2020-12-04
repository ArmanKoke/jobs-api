package databases

import (
	"database/sql"
	"jobs-api/utils/env"

	"go.uber.org/zap"
)

type dbDriver struct {
	postgresConnReader string
	postgresConnWriter string
	postgresDBReader   *sql.DB
	postgresDBWriter   *sql.DB
	postgresDriver     string
}

func (driver *dbDriver) PostgresConnReader() string {
	return driver.postgresConnReader
}

func (driver *dbDriver) PostgresConnWriter() string {
	return driver.postgresConnWriter
}

func (driver *dbDriver) PostgresDBReader() *sql.DB {
	return driver.postgresDBReader
}

func (driver *dbDriver) PostgresDBWriter() *sql.DB {
	return driver.postgresDBWriter
}

func (driver *dbDriver) PostgresDriver() string {
	return driver.postgresDriver
}

func (driver *dbDriver) Close(applicationIntent string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if applicationIntent == env.Reader {
		if driver.postgresDBReader != nil {
			// if err := driver.PostgresDBReader.Close(); err != nil {
			// 	logger.Fatal("Postgres DB reader close error", zap.Error(err))
			// }
		}
	}

	if applicationIntent == env.Writer {
		if driver.postgresDBWriter != nil {
			// if err := driver.PostgresDBWriter.Close(); err != nil {
			// 	logger.Fatal("Postgres DB writer close error", zap.Error(err))
			// }
		}
	}
}
