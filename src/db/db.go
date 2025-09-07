package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DB struct {
	DB *gorm.DB
}

// NewDB initializes a Postgres connection. It supports either:
// 1. DATABASE_URL (full URL) OR
// 2. Individual DB_* variables: DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
func NewDB() *DB {
	var dsn string

	if full := os.Getenv("DATABASE_URL"); full != "" {
		dsn = full
	} else {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		if host == "" || port == "" || user == "" || dbname == "" {
			panic("missing database configuration: set DATABASE_URL or DB_HOST, DB_PORT, DB_USER, DB_NAME")
		}
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &DB{DB: db}
}
