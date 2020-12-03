package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
