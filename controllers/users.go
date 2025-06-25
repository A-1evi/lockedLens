package controllers

import (
	"fmt"
	"lenslocked/views"
	"net/http"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("tailwindcss", "views/users/new.gohtml"),
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
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema: "email"`
	Password string `schema: "password"`
}
