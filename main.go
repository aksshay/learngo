package main

import (
	"net/http"

	"github.com/aksshay/learngo/controllers"
	"github.com/aksshay/learngo/services"
	"github.com/gorilla/mux"
)

const (
	host     = "192.168.1.10"
	port     = "5432"
	user     = "postgres"
	password = "mydb123"
	dbname   = "learngo"
)

func main() {

	//Creating model service
	dataSourceName := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	gormDB, err := services.DBConnection(dataSourceName)
	must(err)

	r := mux.NewRouter()
	userCtrl := controllers.NewUsers(gormDB)

	userAPI := "/api/user/"

	r.HandleFunc(userAPI+"create", userCtrl.Create).Methods("POST")
	r.HandleFunc(userAPI+"{id:[0-9]+}", userCtrl.Read).Methods("GET")
	r.HandleFunc(userAPI+"update/{id:[0-9]+}", userCtrl.Update).Methods("PUT")
	r.HandleFunc(userAPI+"delete/{id:[0-9]+}", userCtrl.Delete).Methods("DELETE")

	must(http.ListenAndServe(":8080", r))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
