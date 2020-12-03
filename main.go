package main

import (
	"fmt"
	"net/http"
)

func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello user!")
}

func main() {

	http.HandleFunc("/hello", helloUser)

	http.ListenAndServe(":8080", nil)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
