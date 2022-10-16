package main

import (
	"fmt"
	"log"

	"ts/pkg/users"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	u := users.NewUser(1, "Meruyert")
	err := u.Create()
	if err != nil {
		log.Println(err)
		return
	}
	err = u.Update("Ardak")
	if err != nil {
		log.Println(err)
		return
	}

	u1, err := u.Read()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(u1)
}
