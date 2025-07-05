package controllers

import "lenslocked/views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("tailwindcss", "static/home"),
		Contact: views.NewView("tailwindcss", "static/contact"),
		Faq:     views.NewView("tailwindcss", "static/faq"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
