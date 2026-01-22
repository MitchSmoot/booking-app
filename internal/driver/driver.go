package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds Database Connection Pool
type DB struct {
	SQL *sql.DB
}

var dbConnection = &DB{}

const maxOpenDbConnections = 10
const maxIdleDbConnections = 5
const maxDbLifetime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConnections)
	d.SetMaxIdleConns(maxIdleDbConnections)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConnection.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConnection, nil

}

// tries to ping the database7
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates new DB for application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
