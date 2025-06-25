package controllers

import "lenslocked/views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("tailwindcss", "views/static/home.gohtml"),
		Contact: views.NewView("tailwindcss", "views/static/contact.gohtml"),
		Faq:     views.NewView("tailwindcss", "views/static/faq.gohtml"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
