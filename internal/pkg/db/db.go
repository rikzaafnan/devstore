package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection(DBDriver string, DBConn string) (*sqlx.DB, error) {

	db, err := sqlx.Connect(DBDriver, DBConn)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established....")

	return db, nil

}
