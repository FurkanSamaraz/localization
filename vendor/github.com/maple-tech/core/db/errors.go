package db

import (
	"database/sql"
	"errors"

	pq "github.com/lib/pq"
)

//Built from https://github.com/lib/pq/blob/master/error.go
const (
	//PGErrForeignKey foreign key didn't match
	PGErrForeignKey = "23503"
	//PGErrUnique unique constraint didn't match
	PGErrUnique = "23505"
)

//ErrUnique error when the unique constraint fails
var ErrUnique = errors.New("unique")

//ErrForeign error when the foreign key is invalid
var ErrForeign = errors.New("foreign")

//ErrExists error when a database existance check returned true
var ErrExists = errors.New("exists")

//ErrNotExists error when a check for existance returns false
var ErrNotExists = errors.New("not-exists")

//ErrDiscrepancy error when the results of a select or search did not return the same amount of information requested
var ErrDiscrepancy = errors.New("discrepancy")

//CheckError briefly checks the PGError type against a constraint number
func CheckError(err error, code string) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		if pgErr.Code == pq.ErrorCode(code) {
			return true
		}
	}
	return false
}

//IsEmptyError returns true if the error is because there are no rows
func IsEmptyError(err error) bool {
	return err == sql.ErrNoRows
}
