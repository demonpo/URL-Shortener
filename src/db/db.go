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

// new database
func NewDB() *DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: conn,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//println("Connected to database")
	//err = db.Table("users").AutoMigrate(&entitiesInfra.User{})
	//if err != nil {
	//	panic("failed to auto migrate models")
	//}
	return &DB{
		DB: db,
	}
}
