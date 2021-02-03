package models

import (
	"errors"

	"../services"
	"gorm.io/gorm"
)

//User the properties related to the user details of our application
type User struct {
	gorm.Model
	UserName  string `gorm:"not null"`
	FirstName string
	LastName  string
	Email     string
	Password  string
}

//CreateUser will connect to the database and return whether the user is created
//successfully or not
func (u *User) CreateUser(db *gorm.DB) (uint, error) {
	pass, err := services.GenerateMD5(u.Password)
	if err != nil {
		return 0, errors.New("Unable to generate hash")
	}
	u.Password = pass
	result := db.Create(u)
	if result.RowsAffected == 1 {
		return u.ID, nil
	} else {
		return 0, errors.New("Failed to create an entry in Database")
	}
}

//ReadUser will get the user data based on id
func (u *User) ReadUser(db *gorm.DB) (interface{}, error) {
	result := db.First(&u, u.ID)
	if result.RowsAffected == 1 {
		return u, nil
	} else {
		return "", errors.New("Failed to retrieve user details")
	}
}

//DeleteUser will delete the user with specific id
func (u *User) DeleteUser(db *gorm.DB) error {
	result := db.Delete(&u)
	if result.RowsAffected == 1 {
		return nil
	}
	return errors.New("Failed to delete user")

}

//UpdateUser will update the user based on id
func (u *User) UpdateUser(db *gorm.DB) error {
	db.First(&u)
	result := db.Save(&u)
	if result.RowsAffected == 1 {
		return nil
	}
	return errors.New("Failed to update user")
}
