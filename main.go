package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type data struct {
	Id         int
	first_name string
	last_name  string
	interests  string
}

const (
	port     = "5432"
	userName = "postgres"
	password = "bOre7sue"
	host     = "localhost"
	dbname   = "postgres"
)

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatal(err)
	}
	createTable(db)
	createUser(db)
	getUser(db)
	postUser(db)
	deleteUser(db)
}

func newDB() (*sqlx.DB, error) {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New("file://schema-seq", connStr)
	if err != nil {
		fmt.Println("failed to make migrate", err)
		return nil, err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("failed to m.Up() in migrate")
		return nil, err
	}
	return db, nil
}

func createTable(db *sqlx.DB) {
	statement, err := db.Prepare("CREATE TABLE if NOT EXISTS data(Id SERIAL PRIMARY KEY, first_name VARCHAR(30) NOT NULL, last_name VARCHAR(30) NOT NULL, interests VARCHAR(30) NOT NULL);")
	if err != nil {
		log.Println(err)
	}
	statement.Exec()
}

func createUser(db *sqlx.DB) {
	statement, err := db.Prepare("INSERT INTO data VALUES (1, 'meru', 'temirbekova', 'golang, c++');")
	if err != nil {
		log.Println(err)
	}
	statement.Exec()
}
func postUser(db *sqlx.DB) {
	statement, err := db.Prepare("UPDATE data SET first_name = 'meruyert' WHERE last_name = 'temirbekova';")
	if err != nil {
		log.Println(err)
	}
	statement.Exec()
}
func getUser(db *sqlx.DB) {

	s := &data{}
	statement, err := db.Prepare("SELECT * from data")
	if err != nil {
		log.Println(err)
	}
	q := statement.QueryRow()
	q.Scan(&s.Id, &s.first_name, &s.last_name, &s.interests)
	fmt.Println(s.first_name)
}

func deleteUser(db *sqlx.DB) {
	statement, err := db.Prepare("DELETE FROM data WHERE last_name = 'temirbekova';")
	if err != nil {
		log.Println(err)
	}
	statement.Exec()
}
