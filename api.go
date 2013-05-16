/*
Package crudapi implements a RESTful JSON API exposing CRUD functionality relying on a custom storage.

See http://en.wikipedia.org/wiki/RESTful and http://en.wikipedia.org/wiki/Create,_read,_update_and_delete for more information.

An example can be found at: https://github.com/sauerbraten/crudapi#example
*/
package crudapi

import (
	"github.com/gorilla/mux"
)

type apiResponse struct {
	Error  string      `json:"error,omitempty"`
	Id     string      `json:"id,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// Adds CRUD and OPTIONS routes to the router, which rely on the given Storage.
func MountAPI(router *mux.Router, api ApiMethods) {

	// Create
	router.HandleFunc("/{kind}", api.CreateOne).Methods("POST")

	// Read
	router.HandleFunc("/{kind}", api.ReadAll).Methods("GET")
	router.HandleFunc("/{kind}/{id}", api.ReadOne).Methods("GET")

	// Update
	router.HandleFunc("/{kind}/{id}", api.UpdateOne).Methods("PUT")

	// Delete
	router.HandleFunc("/{kind}", api.DeleteAll).Methods("DELETE")
	router.HandleFunc("/{kind}/{id}", api.DeleteOne).Methods("DELETE")

	// Options routes for API discovery
	router.HandleFunc("/{kind}", api.OptionsAll).Methods("OPTIONS")
	router.HandleFunc("/{kind}/{id}", api.OptionsOne).Methods("OPTIONS")
}
