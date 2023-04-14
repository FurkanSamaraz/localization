package db

import (
	"errors"
)

//Get shortcut for a sqlx Get query
func Get(dest interface{}, query string, args ...interface{}) error {
	if conn == nil {
		return errors.New("attempted to use db package without it initialized")
	}

	return conn.Get(dest, query, args...)
}

//Select shortcut for a sqlx Select query
func Select(dest interface{}, query string, args ...interface{}) error {
	if conn == nil {
		return errors.New("attempted to use db package without it initialized")
	}

	return conn.Select(dest, query, args...)
}

//Exec shortcut for a sqlx Exec query that drops the results
func Exec(query string, args ...interface{}) error {
	if conn == nil {
		return errors.New("attempted to use db package without it initialized")
	}

	_, e := conn.Exec(query, args...)
	return e
}
