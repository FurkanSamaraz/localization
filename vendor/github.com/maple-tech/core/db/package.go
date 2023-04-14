// Package db maintains the connection to PostgreSQL database. Also provides helpful types and functions for
// interacting with the pgsql servers. The types supplied are purely database wrappers. Most useful types in general
// are in the core.types package
package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Binds the postgresql drivers

	"github.com/maple-tech/core/config"
	"github.com/maple-tech/core/control"
	"github.com/maple-tech/core/log"
)

var conn *sqlx.DB

// Initialize sets up and connects to the database given the config options provided.
// It will bind to the core.control package for auto-shutdown when the service closes
func Initialize(cfg *config.OptionsSQL) error {
	//Grab the configuration if we need it
	if cfg == nil {
		if !config.IsLoaded() {
			return errors.New("attempted to initialize core.db package but no configuration has been loaded yet")
		}

		cfg = &(config.Get().SQL)
	}

	//Protect if we already opened a connection
	if conn != nil {
		err := conn.Close()
		if err != nil {
			return fmt.Errorf("failed to close existing database connection on db.Initialize(); error = %s", err.Error())
		}
	}

	//Should we use SSL mode for the connection?
	ssl := "disable"
	if cfg.SSL {
		ssl = "verify-full"
	}

	//Build the connection string
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password, ssl)

	//Open the connection
	var err error
	conn, err = sqlx.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection in db.Initialize(); error = %s", err.Error())
	}

	log.Infof("opened database connection to postgres: host=%s port=%d dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Database, ssl)

	//Test the connection is valid for sanity
	err = conn.Ping()
	if err != nil {
		Close()
		return fmt.Errorf("failed to test database connection; error = %s", err.Error())
	}
	log.Debug("tested the database connection via ping")

	//Bind the shutdown callback into the control package so we can cleanup after ourselves later
	control.AddShutdownCallback("db", Close)
	conn.DB.SetMaxIdleConns(5)
	conn.DB.SetMaxOpenConns(3)
	conn.DB.SetConnMaxLifetime(5 * time.Minute)

	//Everything seems ok
	return nil
}

// Close terminates the database connection
func Close() {
	if conn == nil {
		return
	}

	conn.Close()
	conn = nil

	log.Info("closed connection to database")
}

// GetConnection returns the current database connection
func GetConnection() *sqlx.DB {
	if conn == nil {
		log.Error("attempted to use db package without it initialized")
	}
	return conn
}
