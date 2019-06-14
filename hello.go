package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

func checkDb(memberId int) string {
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
	row := db.QueryRow(sqlStatement, 2)
	switch err := row.Scan(&id, &name, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, name, email)
		return strings.Join([]string{strconv.Itoa(id), name, email}, " ")
	default:
		panic(err)
	}
	return ""
}

func main() {
	checkDb(2)
	http.HandleFunc("/", print)
	http.ListenAndServe(":8080", nil)
}
