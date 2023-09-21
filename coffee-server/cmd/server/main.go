package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agmguerra/coffee-api/db"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	// Model
}

func (app *Application) Serve() error {
	fmt.Println("API is listening in port: ", app.Config.Port)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", app.Config.Port),
		// TODO: Add router
	}
	return srv.ListenAndServe()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	// Database conncetion
	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Cannot get database connections: %s", dsn)
	}
	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
