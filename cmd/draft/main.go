package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Route gives the struct of a route toward a method
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// GetAll return all the elements of a the basic authentication collection
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the RESTful api")
}

// Routes is an array of Route
type Routes []Route

var testroutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Test,
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range testroutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// print env
	log.Println("Running api server ...")

	http.ListenAndServe(":8080", router)
}
