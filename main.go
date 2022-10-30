package main

import (
	"net/http"
	"ts/web/user"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", user.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", user.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", user.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/user/{id}", user.UpdateUserHandler).Methods(http.MethodPut)
	http.ListenAndServe(":8080", r)
}
