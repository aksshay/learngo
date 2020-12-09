package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aksshay/learngo/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//Users struct that is created using the models
type Users struct {
	db *gorm.DB
}

var user models.User

//NewUsers To return a user service, to call the functions associated with it
func NewUsers(db *gorm.DB) *Users {
	return &Users{db}
}

//Create This will accept the user value and create the entry in database
//POST /api/user/create [curl -X POST  http://localhost:8080/api/users/create]
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "Unable to parse the request post data")
	} else {
		id, err := user.CreateUser(u.db)
		if err != nil {
			panic(err)
		}
		msg := fmt.Sprintf("User with id: %d created successfully!\n", id)
		fmt.Fprintln(w, msg)
	}
	// user := models.User{
	// 	UserName 	: "johndoe",
	// 	FirstName 	: "john",
	// 	LastName  	: "doe",
	// 	Email     	: "johndoe@gmail.com",
	// 	Password 	: "john@123",
	// }
}

//Read It will receive id as an input and generate the user details as response
//GET /api/user/id
func (u *Users) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user.ID = uint(id)
	result, err := user.ReadUser(u.db)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, result)
	}
}

//Update It will update the user with the id details provided in request
//PUT /api/user/update
func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user.ID = uint(id)
	err := user.UpdateUser(u.db)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "User updated successfully")
	}
}

//Delete It will delete the user from the record. It is not deleting, just
//setting the deleted_at for the user
//DELETE /api/user/delete/id
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user.ID = uint(id)
	err := user.DeleteUser(u.db)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "User deleted successfully")
	}
}
