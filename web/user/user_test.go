package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"ts/pkg/users"
)

func TestCRUDUser(t *testing.T) {
	u1 := &users.Users{
		Id:   11,
		Data: "hello11",
	}
	tt, err := json.Marshal(u1)
	if err != nil {
		t.Error("failed test when marshal", "error: ", err)
	}
	resp, err := http.Post("http://localhost:8080/user/11", "application/json", bytes.NewBuffer(tt))
	if err != nil {
		t.Error("failed test while posting resp", "error: ", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Error("the user is not created")
	}
	u2 := forRead(u1.Id)
	if (u1.Id != u2.Id) || (u1.Data != u2.Data) {
		t.Error("the user cannot be read")
	}

	req3, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/user/%d", u1.Id), bytes.NewBuffer([]byte("newData")))
	if err != nil {
		t.Error("failed while putting", err)
	}
	resp3, err := http.DefaultClient.Do(req3)
	if err != nil {
		t.Error(err)
	}
	if resp3.StatusCode != http.StatusOK {
		t.Error("Error when testing delete", resp3.StatusCode)
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/user/%d", u1.Id), nil)
	if err != nil {
		t.Error(err)
	}
	resp2, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp2.StatusCode != http.StatusNoContent {
		t.Error("Error when testing delete", resp2.StatusCode)
	}
}

func forRead(id int) *users.Users {
	myURL := fmt.Sprintf("http://localhost:8080/user/%d", id)
	u2 := &users.Users{}
	resp, err := http.Get(myURL)
	if err != nil {
		log.Println("failed while getting the resp", "error: ", err)
		return nil
	}
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error on read body from %s", myURL)
		return nil
	}
	err = json.Unmarshal(bb, u2)
	if err != nil {
		fmt.Println("error on unmarshall body", "error: ", err)
		return nil
	}
	return u2
}
