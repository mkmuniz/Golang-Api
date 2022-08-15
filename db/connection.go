package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=user_api password=1234 dbname=api_go sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
