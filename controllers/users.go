package controllers

import (
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

type Users struct {
	NewView *views.View
}
