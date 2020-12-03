package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "192.168.1.10"
	port     = 5432
	user     = "postgres"
	password = "mydb123"
	dbname   = "learngo"
)

func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello user using Gorilla!")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", helloUser)

	must(http.ListenAndServe(":8080", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
