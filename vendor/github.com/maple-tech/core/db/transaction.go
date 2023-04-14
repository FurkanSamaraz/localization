package db

import (
	"github.com/jmoiron/sqlx"

	"github.com/maple-tech/core/log"
)

//WrapTX takes a callback function, that will be provided a sqlx.Tx transaction pointer.
//It will then wrap the callback function in a pgsql transaction.
//Should catch any panics and errors thrown during the transaction and return it back immediately.
func WrapTX(fn func(*sqlx.Tx) error) (err error) {
	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Errorf("panic in database transaction; panic = %s", p)
			/*
				err = tx.Rollback()
				if err != nil {
					log.Errorf("rollback after panic failed; error = %s", err.Error())
				}
				log.Debug("transaction rolled-back")
			*/
			panic(p)
		} else if err != nil {
			log.Errorf("error in database transaction; error = %s", err)
			re := tx.Rollback()
			if re != nil {
				log.Errorf("transaction rollback failed; error = %s", re.Error())
			}
			log.Debug("transaction rolled-back")
		} else {
			err = tx.Commit()
			if err != nil {
				log.Errorf("error commiting transaction; error = %s", err)
			}
		}
	}()

	err = fn(tx)
	return
}
