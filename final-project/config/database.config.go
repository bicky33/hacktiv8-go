package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewDB() (*Postgres, error) {
	configApp, err := Config()
	if err != nil {
		log.Fatalln(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		configApp.DbHost, configApp.DbUser, configApp.DbPassword, configApp.DbName, configApp.DbPort, configApp.DbSslMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)
	return &Postgres{DB: db}, nil
}
