package controllers

import (
	"fmt"
	"net/http"
)

//Users struct that is created using the models
type Users struct {
}

//NewUsers To return a user service, to call the functions associated with it
func NewUsers() *Users {
	return &Users{}
}

//Create This will accept the user value and create the entry in database
//POST /api/users/create
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is post request")
}

//Read It will receive id as an input and generate the user details as response
//GET /api/users/get
func (u *Users) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is get request")
}

//Update It will update the user with the id details provided in request
//PUT /api/users/update
func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is put request")
}

//Delete It will delete the user from the record. It is not deleting, just
//setting the deleted_at for the user
//DELETE /api/users/delete
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is delete request")
}
