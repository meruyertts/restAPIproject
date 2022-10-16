package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
)

const (
	port     = "5432"
	userName = "postgres"
	password = "qwerty"
	host     = "localhost"
	dbname   = "postgres"
)

var DB *sqlx.DB

func init() {
	DB = NewDB()
}

func NewDB() *sqlx.DB {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New("file://schema-seq", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	return db
}
