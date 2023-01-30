package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

func Connect() *pgxpool.Pool {
	/*
		This is a function that is used to connect to a Postgres Connection Pool.
		We use pgxpool to connect as it automatically manages the Connection Pool
		and we can simply use the Connect() function to return an open connection.
	*/

	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	DATABASE_URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	conn, err := pgxpool.Connect(context.Background(), DATABASE_URL)

	if err != nil {
		log.Println("Failed to connect to pgoltp!", err)
	}
	return conn
}
