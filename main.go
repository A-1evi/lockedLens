package main

import (
	"fmt"
	"lenslocked/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

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
	staticController := controllers.NewStatic()

	usersController := controllers.NewUsers()
	var h http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.Handle("/", staticController.Home).Methods("GET")
	r.Handle("/contact", staticController.Contact).Methods("GET")
	r.Handle("/faq", staticController.Faq).Methods("GET")
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
