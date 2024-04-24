package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDBConnection(driverDB string, database string, host string, user string, password string, port int) (*gorm.DB, error) {
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if driverDB == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			host,
			port,
			user,
			database,
			password,
			"disable",
		)

		dialect = postgres.Open(dsn)
	} else {
		return nil, fmt.Errorf("error to connect db")
	}

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// pool time
	tm := time.Minute * time.Duration(20)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(tm)

	return db, nil
}
