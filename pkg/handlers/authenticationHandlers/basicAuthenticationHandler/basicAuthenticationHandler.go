package basicAuthenticationHandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Ivandolchevic/goapis/pkg/utils/responseUtil"

	database "github.com/Ivandolchevic/goapis/pkg/data"
	. "github.com/Ivandolchevic/goapis/pkg/models/business"
	. "github.com/Ivandolchevic/goapis/pkg/models/util"
	authenticationUtil "github.com/Ivandolchevic/goapis/pkg/utils/authenticationUtil"
	httpUtil "github.com/Ivandolchevic/goapis/pkg/utils/httpUtil"
)

var collectionname = "basicauthentications"

// GetAll return all the elements of a the basic authentication collection
func GetAll(w http.ResponseWriter, r *http.Request) *APIError {
	fmt.Print("Executing GETALL")
	// read the filter
	filter := r.Header.Get("filter")

	// Select the objects from the database
	elements, err := database.Find(collectionname, filter)
	if err != nil {
		return &APIError{Code: http.StatusInternalServerError, Error: err, Message: "The data could not be selected"}
	}

	// Send back the OK response
	responseUtil.Answer(w, elements, nil)

	return nil
}

// Get return an element of a collection using its identifier
func Get(w http.ResponseWriter, r *http.Request) *APIError {
	authenticationUtil.CreateToken()

	// Retrieve the identifier from the url
	vars := mux.Vars(r)
	id := vars["id"]

	// Select the object from the database
	element, err := database.FindByID(collectionname, id)
	if err != nil {
		return &APIError{Code: http.StatusInternalServerError, Error: err, Message: "The data could not be selected"}
	}

	// Send back the OK response
	responseUtil.Answer(w, element, nil)

	return nil
}

// Put create a new basic authentication element into the database
func Put(w http.ResponseWriter, r *http.Request) *APIError {
	var basic BasicAuthentication
	var oJSON []byte
	var element interface{}

	// Extract the json object as an array of bytes
	oJSON, err := httpUtil.ExtractJSON(r)
	if err != nil {
		return &APIError{Code: http.StatusBadRequest, Error: err, Message: "The input format is not a JSON"}
	}

	// Try to map the json to the object
	err = basic.Mapper(oJSON)
	if err != nil {
		return &APIError{Code: http.StatusUnprocessableEntity, Error: err, Message: "The JSON transmitted is incomplete or inconsistent"}
	}

	// Securize the data
	err = basic.Secure()
	if err != nil {
		return &APIError{Code: http.StatusInternalServerError, Error: err, Message: "The data could not be securized"}
	}

	// Insert the object into the database
	element, err = database.Upsert(collectionname, basic)
	if err != nil {
		return &APIError{Code: http.StatusInternalServerError, Error: err, Message: "The data could not be persisted"}
	}

	// Send back the OK response
	responseUtil.Answer(w, element, nil)

	return nil
}
