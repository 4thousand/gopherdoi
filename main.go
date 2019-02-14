package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"math/rand"

	"time"
)

// Person The person Type
type Person struct {
	Queue int64  `json:"queue"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var queue int64 = 5000
var queuelist int64 = 0
var people []Person

//create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var person Person
	fmt.Println("Body:", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&person)
	fmt.Println("parmeter", person)

	if queuelist <= 5000 {
		fmt.Println(queuelist)
		queuelist += 1
		person.Queue = queuelist
		//q := len(people)
		queue -= 1

		c := rand.Intn(1000)
		time.Sleep(time.Duration(c) * time.Millisecond)
		people = append(people, person)
		json.NewEncoder(w).Encode((person.Queue))
	} else {
		json.NewEncoder(w).Encode("full")
	}
}

// ListPerson create a new item
func ListPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(people)
}

// CountPerson
func CountPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(len(people))
}

// main function to boot up
func main() {
	router := mux.NewRouter()

	fmt.Println("Start Serve Port:9000 ")
	router.HandleFunc("/post/", CreatePerson).Methods("POST")
	router.HandleFunc("/list/", ListPerson).Methods("GET")
	router.HandleFunc("/count/", CountPerson).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", router))
}
