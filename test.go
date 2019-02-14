package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type salesRepository struct{ db *sqlx.DB }

// The person Type (more like an object)
type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
type ModelPerson struct {
	Id    int64  `db:"id"`
	Qeue  int64  `db:"qeue"`
	Code  string `db:"code"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

var people []Person
var mysql_np *sqlx.DB
var qeue int64 = 5000

// Display all from the people var

func ConnectMysqlNP(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	dsn := "root:1234@tcp(localhost:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	//fmt.Println(dsn,"DBName =", dbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	}
	fmt.Println("Connected")
	return db, err
}

func init() {

	mysql_nopadol, err := ConnectMysqlNP("carpark")
	if err != nil {
		fmt.Println(err.Error())
	}
	mysql_np = mysql_nopadol

}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	fmt.Println(person)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", healthCheckHandler)
	router.HandleFunc("/GET/", GetPerson).Methods("GET")
	router.HandleFunc("/post/", CreatePerson).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(2)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Success bool `json:"api success"`
	}{true})
}
