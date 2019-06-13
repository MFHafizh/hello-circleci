package main

import (
	"fmt"
	"net/http"
)

func print(w http.ResponseWriter, r *http.Request) {
	msg := message("hafizh")
	fmt.Fprintf(w, msg)
}

func message(name string) string {
	return "Hello " + name
}

func main() {
	http.HandleFunc("/", print)
	http.ListenAndServe(":8080", nil)
}
