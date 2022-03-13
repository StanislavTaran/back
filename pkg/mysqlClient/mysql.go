package mysqlClient

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Config provides options to establish connection to MySQL db
type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	MaxRetries    int    `json:"maxRetries"`
	WaitRetry     int    `json:"waitRetry"`
}

// MySQLClient a way to work with MySQL database
type MySQLClient struct {
	config *Config
	db     *sql.DB
}

// New - initialize new MySQL struct with config
func New(config *Config) *MySQLClient {
	return &MySQLClient{
		config: config,
	}
}

// Open new MySQL connection using passed to New func config
func (m *MySQLClient) Open() error {
	db, err := sql.Open("mysql", m.config.ConnectionURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	m.db = db

	return nil
}

// Close current MySQL connection
func (m *MySQLClient) Close() error {
	return m.db.Close()
}
