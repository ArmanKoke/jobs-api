package env

import (
	"fmt"
	"os"
)

const (
	// DebugMode env var name
	DebugMode = "DEBUG"
	// DefaultDebugMode value
	DefaultDebugMode = ""

	// IPAddress of the app
	IPAddress = "IP_ADDRESS"
	// DefaultIPAddress is default variable for docker
	// DefaultIPAddress = "0.0.0.0"
	DefaultIPAddress = "127.0.0.1"

	// Port of the app
	Port = "PORT"
	// DefaultPort is default variable
	DefaultPort = "8080"

	// PostgresDriver env var name
	PostgresDriver = "POSTGRES_DRIVER"
	// DefaultPostgresDriver value
	DefaultPostgresDriver = "postgres"

	// PostgresIP env var name
	PostgresIP = "POSTGRES_IP"
	// DefaultPostgresIP value
	DefaultPostgresIP = "127.0.0.1"

	PostgresPort        = "POSTGRES_PORT"
	DefaultPostgresPort = "5432"

	PostgresDatabase        = "POSTGRES_DATABASE"
	DefaultPostgresDatabase = "dev"

	PostgresReaderUser        = "POSTGRES_READER_USER"
	DefaultPostgresReaderUser = "postgres"

	PostgresWriterUser        = "POSTGRES_WRITER_USER"
	DefaultPostgresWriterUser = "postgres"

	PostgresReaderPassword        = "POSTGRES_READER_PASSWORD"
	DefaultPostgresReaderPassword = "dumb123!"

	PostgresWriterPassword        = "POSTGRES_WRITER_PASSWORD"
	DefaultPostgresWriterPassword = "dumb123!"

	Reader = "ReadOnly"
	Writer = "ReadWrite"
)

// AppDebugMode env value getter
func AppDebugMode() (debugMode string) {
	if debugMode = os.Getenv(DebugMode); "" == debugMode {
		debugMode = DefaultDebugMode
	}

	return
}

// AppIPAddress env value getter
func AppIPAddress() (ipAddress string) {
	ipAddress = DefaultIPAddress
	return
}

// AppPort env value getter
func AppPort() (port string) {
	port = DefaultPort
	return
}

// AppPostgresDriver env value getter
func AppPostgresDriver() (driver string) {
	if driver = os.Getenv(PostgresDriver); "" == driver {
		driver = DefaultPostgresDriver
	}

	return
}

// AppPostgresIPAddress env value getter
func AppPostgresIPAddress() (ipAddress string) {
	if ipAddress = os.Getenv(PostgresIP); "" == ipAddress {
		ipAddress = DefaultPostgresIP
	}

	return
}

// AppPostgresPort env value getter
func AppPostgresPort() (port string) {
	if port = os.Getenv(PostgresPort); "" == port {
		port = DefaultPostgresPort
	}

	return
}

// AppPostgresDatabase env value getter
func AppPostgresDatabase() (database string) {
	if database = os.Getenv(PostgresDatabase); "" == database {
		database = DefaultPostgresDatabase
	}

	return
}

// AppPostgresUser env value getter
func AppPostgresUser(applicationIntent string) (user string) {
	if Reader == applicationIntent {
		if user = os.Getenv(PostgresReaderUser); "" == user {
			user = DefaultPostgresReaderUser
		}
		return
	}

	if Writer == applicationIntent {
		if user = os.Getenv(PostgresWriterUser); "" == user {
			user = DefaultPostgresWriterUser
		}
	}

	return
}

// AppPostgresPassword env value getter
func AppPostgresPassword(applicationIntent string) (password string) {
	if Reader == applicationIntent {
		if password = os.Getenv(PostgresReaderPassword); "" == password {
			password = DefaultPostgresReaderPassword
		}
		return
	}

	if Writer == applicationIntent {
		if password = os.Getenv(PostgresWriterPassword); "" == password {
			password = DefaultPostgresWriterPassword
		}
	}

	return
}

// AppPostgresConn connection source
// applicationIntent reader or writer
func AppPostgresConn(applicationIntent string) (conn string) {
	ipAddress := AppPostgresIPAddress()
	port := AppPostgresPort()
	database := AppPostgresDatabase()
	user := AppPostgresUser(applicationIntent)
	password := AppPostgresPassword(applicationIntent)

	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", // todo change in cluster
		ipAddress, port, user, password, database)

	return
}
