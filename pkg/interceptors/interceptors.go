package interceptors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
	})
}

func Authentifier(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Check rights ...")

		//err := "Unauthorized: the user is not allowed to access to this method"
		err := ""
		if err != "" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		} else {
			h.ServeHTTP(w, r) // call original
		}
	})
}

/*
func ControlAllow(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Control allow")

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}

		// Stop here for a Preflighted OPTIONS request.
		if r.Method == "OPTIONS" {
			return
		}

		// Lets Gorilla work
		h.ServeHTTP(w, r)
	})
}
*/
func ControlAllow(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Control allow")

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
}
