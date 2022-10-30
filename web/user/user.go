package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"ts/pkg/users"

	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	newID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}
	u1 := &users.Users{
		Id: newID,
	}

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bb, u1)
	if err != nil {
		log.Println("error while unmarshalling ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = u1.Create()
	if err != nil {
		log.Println("error on create user ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("user was created", u1)

	w.WriteHeader(http.StatusCreated)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	newID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}
	u1 := &users.Users{
		Id: newID,
	}
	u2, err := u1.Read()
	if err != nil {
		fmt.Println("error when reading", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(u2)
	if err != nil {
		fmt.Println("error when encoding", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(u2)
	w.WriteHeader(http.StatusOK)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	newID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}
	u1 := &users.Users{
		Id: newID,
	}
	err = u1.Delete()
	if err != nil {
		fmt.Println("error when deleting", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	newID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}
	u1 := &users.Users{
		Id: newID,
	}
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = u1.Update(string(bb))
	if err != nil {
		fmt.Println("error when updating", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
