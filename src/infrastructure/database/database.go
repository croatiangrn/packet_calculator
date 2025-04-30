package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"time"
)

const (
	// DefaultMaxConn is the default maximum number of open connections to the database.
	DefaultMaxConn = 10
	// DefaultMaxIdleConn is the default maximum number of idle connections in the pool.
	DefaultMaxIdleConn = 5
	// DefaultMaxConnLifeTime is the default maximum connection lifetime in seconds.
	DefaultMaxConnLifeTime = 60
)

func InitDB(DbDsn string) (*sql.DB, error) {
	db, e := sql.Open("mysql", DbDsn)
	if e != nil {
		return nil, e
	}
	db.SetConnMaxLifetime(time.Duration(DefaultMaxConnLifeTime) * time.Second)
	db.SetMaxOpenConns(DefaultMaxConn)
	db.SetMaxIdleConns(DefaultMaxIdleConn)
	return db, nil
}
