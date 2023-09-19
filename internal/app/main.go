package app

import (
	"context"
	"crud-go/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Application struct {
	Port    int64
	app     *http.Server
	Pg      *sql.DB
	Context context.Context
}

func getIntEnv(key string) int {
	val := os.Getenv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("some error", err))
	}
	return ret
}

func Initialize() (*Application, error) {
	e := godotenv.Load(".env")
	if e != nil {
		log.Fatalf("Some error occured: %s", e)
	}
	config := config.Config{
		Port: int64(getIntEnv("PORT")),
		PostgresDatabase: config.PostgresDatabase{
			Username: os.Getenv("POSTGRESQL_USERNAME"),
			Password: os.Getenv("POSTGRESQL_PASSWORD"),
			Host:     os.Getenv("POSTGRESQL_HOST"),
			Port:     int64(getIntEnv("POSTGRESQL_PORT")),
			DBname:   os.Getenv("POSTGRESQL_NAME"),
		},
	}

	pg, e := ConnectToPostgreSQLDB(config.PostgresDatabase)
	if e != nil {
		log.Fatal("Connecting to postgres database failed with error ", e)
	}

	a := Application{
		Port:    config.Port,
		Pg:      pg,
		Context: context.Background(),
	}
	app := createApp(a)
	a.app = app

	return &a, nil
}

func (a Application) Run() error {
	e := a.app.ListenAndServe()
	if e != nil {
		log.Fatal("Server stops with error ", e)
	}
	log.Printf("Server starting")
	return nil
}

func ConnectToPostgreSQLDB(config config.PostgresDatabase) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.DBname)
	db, e := sql.Open("postgres", connectionString)
	if e != nil {
		return nil, e
	}

	return db, nil
}
