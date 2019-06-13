package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "circleci_demo"
)

func print(w http.ResponseWriter, r *http.Request) {
	msg := message("hafizh")
	fmt.Fprintf(w, msg)
}

func message(name string) string {
	return "Hello " + name
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", print)
	http.ListenAndServe(":8080", nil)
}
