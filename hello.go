package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "demotable"
)

func print(w http.ResponseWriter, r *http.Request) {
	msg := message("hafizh")
	fmt.Fprintf(w, msg)
}

func message(name string) string {
	return "Hello " + name
}

func checkDb(memberId int64) (int, string, string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	log.Println("connString", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT id, name, email FROM member WHERE id=$1;`
	var email string
	var id int
	var name string
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, memberId)
	switch err := row.Scan(&id, &name, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, name, email)
		return id, name, email
	default:
		panic(err)
	}
	return 0, "", ""
}

func getMemeberData(w http.ResponseWriter, r *http.Request) {
	var email string
	var id int
	var name string
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	i2, err := strconv.ParseInt(memberId, 10, 64)
	if err == nil {
		id, name, email = checkDb(i2)
		if id != 0 && name != "" && email != "" {
			member := Member{Id: id, Name: name, Email: email}
			json.NewEncoder(w).Encode(member)
		}

	} else {
		fmt.Println("err")
	}
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", print)
	router.HandleFunc("/member/{memberId}", getMemeberData)
	log.Fatal(http.ListenAndServe(":8080", router))
}
