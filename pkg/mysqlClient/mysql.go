package mysqlClient

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Config provides options to establish connection to MySQL db
type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	MaxRetries    int    `json:"maxRetries"`
	WaitRetry     int    `json:"waitRetry"`
}

// MySQLClient a way to work with MySQL database
type MySQLClient struct {
	Config *Config
	Db     *sqlx.DB
}

// NewMySQLClient - initialize new MySQL struct with config
func NewMySQLClient(config *Config) *MySQLClient {
	return &MySQLClient{
		Config: config,
	}
}

// Open new MySQL connection using passed to New func config
func (m *MySQLClient) Open() error {
	db, err := sqlx.Connect("mysql", m.Config.ConnectionURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	m.Db = db

	return nil
}

// Close current MySQL connection
func (m *MySQLClient) Close() error {
	return m.Db.Close()
}
