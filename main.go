package main

import (
	"fmt"
	"lenslocked/controllers"
	"lenslocked/views"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))

}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> The page you looking for doesnt exits</h1>")
}

// By changing the last argument in the call to HandleFunc we have instructed our
// router that we now want it to use the New method we defined to handle any web requests
// for the page /signup. Because this method matches the definition of a HandlerFunc
// our program will accept it happily. It doesn’t matter to our router that this is a method attached to the users controller.
// All that matters is that the New method will accept two arguments of the type ResponseWriter and Request.
func main() {
	homeView = views.NewView("tailwindcss", "views/home.gohtml")
	contactView = views.NewView("tailwindcss", "views/contact.gohtml")
	faqView = views.NewView("tailwindcss", "views/faq.gohtml")
	usersController := controllers.NewUsers()
	var h http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
