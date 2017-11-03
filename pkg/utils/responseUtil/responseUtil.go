package responseUtil

import (
	"encoding/json"
	"net/http"
)

// Answer fill the response that will be send back to the caller
func Answer(w http.ResponseWriter, object interface{}, err error) {
	// set the header of the response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// write an error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		panic(err)
	} else {
		// write a success
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(object)
	}

}
