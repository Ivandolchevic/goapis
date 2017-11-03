// This is the package comment, a top-level piece of documentation
// used to explain things about the package (see json or exp/template)
// All godoc comments are in this form
// with no whitespace between them and what they accompany

package main

// That file is the main of the API

import (
	"log"
	"net/http"

	data "github.com/Ivandolchevic/goapis/pkg/data"
	router "github.com/Ivandolchevic/goapis/pkg/routers/authenticationRouter"
	configuration "github.com/Ivandolchevic/goapis/pkg/utils/configurationUtil"
)

func main() {
	// Read the configuration
	configuration := configuration.Get()

	// Initialise the database
	data.InitDatabase(configuration)

	// Initialise routing
	router := router.NewAuthenticationRouter()

	// Launch the server
	log.Fatal(http.ListenAndServe(configuration.GetServerURL(), router))

}
