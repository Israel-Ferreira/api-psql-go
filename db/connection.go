package db

import (
	"database/sql"
	"fmt"
	"todo-list-go/configs"

	_ "github.com/lib/pq"
)

func OpenDbConnection() (*sql.DB, error) {
	confg := configs.GetDb()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		confg.Host,
		confg.Port,
		confg.Username,
		confg.Password,
		confg.DBName,
	)

	connection, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	if err := connection.Ping(); err != nil {
		return connection, err
	}

	return connection, nil
}
