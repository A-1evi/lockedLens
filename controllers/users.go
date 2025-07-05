package controllers

import (
	"fmt"
	"lenslocked/models"
	"lenslocked/views"
	"net/http"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("tailwindcss", "users/new"),
		us:      us,
	}
}

// New will used to render the signup page to new user to signup

//GET /signup

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}

}

// Create is used to process the signup form whena user
// tries  ot create a new user account

// POST /signup
// Our PostForm field is really just a map behind the scenes,
// which means that we can access fields stored in the PostForm field in
// the same way we would access fields in a map - by using the ["key"] syntax
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "user is ", user)
}

type Users struct {
	NewView *views.View
	us      *models.UserService
}

type SignupForm struct {
	Email    string `schema: "email"`
	Password string `schema: "password"`
	Name     string `schema: "name"`
}
