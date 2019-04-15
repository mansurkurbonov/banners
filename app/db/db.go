package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	//postgres database driver
	_ "github.com/lib/pq"
)

// Errors
var (
	ErrNotFound       = errors.New("not found")
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrDuplicate      = errors.New("already exist")
	ErrInvalidStatus  = errors.New("invalid value for status")
)

var conn *sql.DB

// type rowScanner interface {
// 	Scan(dest ...interface{}) error
// }

// Connect establish connection with database.
func Connect(usr, pswd, db, host, port string) {
	var (
		dns = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable connect_timeout=10",
			usr, pswd, db, host, port)

		err error
	)

	conn, err = sql.Open("postgres", dns)
	if err != nil {
		log.Panicln("database: " + err.Error())
	}
	err = conn.Ping()
	if err != nil {
		log.Panicln("database: " + err.Error())
	}
	log.Printf("database: established with %s:%s/%s\n", host, port, db)
}

// Disconnect closes connection with database.
func Disconnect() {
	conn.Close()
}

// // exec makes query without returns.
// func exec(query string, args ...interface{}) (err error) {
// 	var (
// 		res      sql.Result
// 		affected int64
// 	)

// 	res, err = conn.Exec(query, args...)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "invalid input value for enum card_status") {
// 			err = ErrInvalidStatus
// 		}
// 		return
// 	}

// 	affected, err = res.RowsAffected()
// 	if affected == 0 {
// 		err = ErrNoRowsAffected
// 		return
// 	}
// 	return
// }
