package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

var postgressql Postgres

func NewDB() (*Postgres, error) {
	configApp := Config()

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
	postgressql.DB = db
	return &postgressql, nil
}

func GetDB() *Postgres {
	return &postgressql
}
