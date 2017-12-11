package manager

import (
	"database/sql"
	"log"
)

var db *sql.DB

// InitDB initializes postgres db connection
func InitDB(connStr string) {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
