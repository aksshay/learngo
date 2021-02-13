package main

import (
	"net/http"
	"os"

	"./controllers"
	"./services"
	"github.com/gorilla/mux"
)

var (
	host     = readEnv("DBHOST")
	port     = readEnv("DBPORT")
	user     = readEnv("DBUSER")
	password = readEnv("DBPASSWORD")
	dbname   = readEnv("DBNAME")
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

func readEnv(env string) string {
	value := os.Getenv(env)
	if value == "" {
		panic("failed to read environment variable")
	}
	return value
}
