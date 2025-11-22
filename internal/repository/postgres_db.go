package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

type PostgresDB struct {
	Conn *sql.DB
}

func (db *PostgresDB) ConnectToDatabase() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"))
	var err error
	db.Conn, err = sql.Open("pgx", dsn)
	if err == nil {
		log.Println("Successfully connected to database")
	}
	return err
}
