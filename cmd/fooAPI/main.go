package main

import (
	"log"
	"net/http"

	data "github.com/Ivandolchevic/goapis/pkg/data"
	router "github.com/Ivandolchevic/goapis/pkg/routers/fooRouter"
	configuration "github.com/Ivandolchevic/goapis/pkg/utils/configurationUtil"
)

func main() {
	// Read the configuration
	configuration := configuration.Get()

	// Initialise the database
	data.InitDatabase(configuration)

	// Initialise routing
	router := router.NewFooRouter()

	// Launch the server
	log.Fatal(http.ListenAndServe(configuration.GetServerURL(), router))

}
