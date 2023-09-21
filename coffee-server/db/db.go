package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDBLifetime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDBLifetime)

	if err = testDB(d); err != nil {
		return nil, err
	}
	dbConn.DB = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("*** Pinged database successfuly! ***")
	return nil
}
