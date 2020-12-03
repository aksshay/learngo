package main

import (
	"net/http"

	"github.com/aksshay/learngo/controllers"
	"github.com/gorilla/mux"
)

const (
	host     = "192.168.1.10"
	port     = 5432
	user     = "postgres"
	password = "mydb123"
	dbname   = "learngo"
)

func main() {

	r := mux.NewRouter()
	userCtrl := controllers.NewUsers()

	userAPI := "/api/users/"

	r.HandleFunc(userAPI+"create", userCtrl.Create).Methods("POST")
	r.HandleFunc(userAPI+"get", userCtrl.Read).Methods("GET")
	r.HandleFunc(userAPI+"update", userCtrl.Update).Methods("PUT")
	r.HandleFunc(userAPI+"delete", userCtrl.Delete).Methods("DELETE")

	must(http.ListenAndServe(":8080", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
