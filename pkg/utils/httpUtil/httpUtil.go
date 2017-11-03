package httpUtil

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/Ivandolchevic/goapis/pkg/models/util"
	configurationUtil "github.com/Ivandolchevic/goapis/pkg/utils/configurationUtil"
	logUtil "github.com/Ivandolchevic/goapis/pkg/utils/logUtil"
)

// APIHandler extends the basic http handler with a custom error handling
type APIHandler func(http.ResponseWriter, *http.Request) *APIError

func (fn APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	configuration := configurationUtil.Get()

	for _, v := range configuration.OptionsHeaders {
		w.Header().Set(v.Key, v.Value)
	}

	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		// write the error into a log file
		logUtil.WriteError(e)

		// raise an http error
		http.Error(w, e.Message, e.Code)
	} else {
		logUtil.WriteString("request executed")
	}
}

// ExtractJSON extract an array of bytes with a JSON format from the body of the request
func ExtractJSON(r *http.Request) ([]byte, error) {
	// Read the inputed data
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	// Handler error on body closing
	r.Body.Close()

	return body, err
}

// NewRouter instantiates a new router for the current API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.MethodNotAllowedHandler = methodNotAllowedHandler()

	return router
}

// methodNotAllowed handles OPTIONS or replies to the request with an HTTP status code 405.
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	// Check the method of the request
	if r.Method == "OPTIONS" {
		// If the app have to handle options, give back the options
		configuration := configurationUtil.Get()

		for _, v := range configuration.OptionsHeaders {
			w.Header().Set(v.Key, v.Value)
		}

		w.WriteHeader(http.StatusOK)
	} else {
		// Otherwise, the method is not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// methodNotAllowedHandler return the handler for the specials methods like OPTIONS, HEAD, PATCH, ...
func methodNotAllowedHandler() http.Handler { return http.HandlerFunc(methodNotAllowed) }
