package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	DefaultMaxConn         = 10
	DefaultMaxIdleConn     = 5
	DefaultMaxConnLifeTime = 60
)

// InitDB initializes the database connection
func InitDB(DbDsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DbDsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetConnMaxLifetime(time.Duration(DefaultMaxConnLifeTime) * time.Second)
	db.SetMaxOpenConns(DefaultMaxConn)
	db.SetMaxIdleConns(DefaultMaxIdleConn)

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// RunMigrations runs the database migrations
func RunMigrations(db *sql.DB, migrationsPath string) error {
	// Get database driver instance
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// Initialize migrator
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrator: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
