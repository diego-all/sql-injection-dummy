package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 5               // Número maximo de conexiones permitidas abiertas
const maxIdleDbConn = 5               // Número maximo de conexiones inactivas (ociosas) abiertas y disponibles para reutilización
const maxDbLifeTime = 5 * time.Minute // tiempo antes de que se considere inactiva una conexión

func ConnectPostgres(dsn string) (*DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}
	fmt.Println("Database ping successful!")
	return nil
}
