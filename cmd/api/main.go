package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sql-injection-eafit/database"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	//db       *sql.DB
	models database.Models
}

func main() {

	var cfg config
	cfg.port = 9090

	var dsn string

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 1. Hardcoded database credentials in connection string
	// dsn := "host=localhost port=54325 user=postgres password=password dbname=sqli sslmode=disable timezone=UTC connect_timeout=5"

	// 2. Read enviroment variable
	if dsn = os.Getenv("DSN"); dsn == "" {
		fmt.Println("La variable de entorno DSN no está definida.")
	} else {
		fmt.Printf("El valor DSN es: %s\n", dsn)
	}

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		//db:       db,
		models: database.New(db.SQL),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
