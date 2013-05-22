/*
Package crudapi implements a RESTful JSON API exposing CRUD functionality relying on a custom storage.

See http://en.wikipedia.org/wiki/RESTful and http://en.wikipedia.org/wiki/Create,_read,_update_and_delete for more information.

An example can be found at: https://github.com/sauerbraten/crudapi#example
*/
package crudapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type apiResponse struct {
	Error  string      `json:"error,omitempty"`
	Id     string      `json:"id,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func crudReadRequest(resp http.ResponseWriter, req *http.Request) (vars map[string]string, dec *json.Decoder) {
	vars = mux.Vars(req)
	dec = json.NewDecoder(req.Body)
	return
}

func crudWriteResponse(resp http.ResponseWriter, respCode int, apiResp apiResponse) {
	resp.WriteHeader(respCode)
	enc := json.NewEncoder(resp)
	err := enc.Encode(apiResp)
	if err != nil {
		log.Println(err)
	}
	return
}

func crudCall(crudMethod func(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)) (httpMethod func(resp http.ResponseWriter, req *http.Request)) {
	httpMethod = func(resp http.ResponseWriter, req *http.Request) {
		vars, dec := crudReadRequest(resp, req)

		//perform the specific CRUD action requested
		respCode, apiResp := crudMethod(vars, dec)

		crudWriteResponse(resp, respCode, apiResp)
	}
	return
}

// Adds CRUD and OPTIONS routes to the router, which rely on the given Storage.
func MountAPI(router *mux.Router, api ApiMethods) {

	// Create
	router.HandleFunc("/{kind}", crudCall(api.CreateOne)).Methods("POST")

	// Read
	router.HandleFunc("/{kind}", crudCall(api.ReadAll)).Methods("GET")
	router.HandleFunc("/{kind}/{id}", crudCall(api.ReadOne)).Methods("GET")

	// Update
	router.HandleFunc("/{kind}/{id}", crudCall(api.UpdateOne)).Methods("PUT")

	// Delete
	router.HandleFunc("/{kind}", crudCall(api.DeleteAll)).Methods("DELETE")
	router.HandleFunc("/{kind}/{id}", crudCall(api.DeleteOne)).Methods("DELETE")

	// Options routes for API discovery
	router.HandleFunc("/{kind}", api.OptionsAll).Methods("OPTIONS")
	router.HandleFunc("/{kind}/{id}", api.OptionsOne).Methods("OPTIONS")
}
