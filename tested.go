package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Person The person Type (more like an object)
type Person struct {
	Qeue  int64  `json:"qeue"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var people []Person
var qeue int64 = 5000
var qeuelist int64 = 0

// CreatePerson create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	fmt.Println(qeuelist)
	if qeuelist <= 5000 {
		qeuelist += 1
		//q := len(people)
		qeue -= 1
		person.Qeue = qeuelist
		c := rand.Intn(1000)
		time.Sleep(time.Duration(c) * time.Millisecond)
		people = append(people, person)
		json.NewEncoder(w).Encode((people))
	} else {
		json.NewEncoder(w).Encode("full")
	}
}

// ListPerson create a new item
func ListPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(people)
}

// CountPerson create a new item
func CountPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(len(people))
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/post/", CreatePerson).Methods("POST")
	router.HandleFunc("/list/", ListPerson).Methods("GET")
	router.HandleFunc("/count/", CountPerson).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
