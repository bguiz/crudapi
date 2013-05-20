package crudapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NoAuthApiMethods struct {
	s Storage
}

func NewNoAuthApiMethods(store Storage) NoAuthApiMethods {
	return NoAuthApiMethods{store}
}

func crudUnmarshall(resp http.ResponseWriter, req *http.Request) (vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	vars = mux.Vars(req)
	dec = json.NewDecoder(req.Body)
	return
}

func crudMarshall(resp http.ResponseWriter, respCode int, apiResp apiResponse, enc *json.Encoder) {
	resp.WriteHeader(respCode)
	enc = json.NewEncoder(resp)
	err := enc.Encode(apiResp)
	if err != nil {
		log.Println(err)
	}
	return
}

func (self NoAuthApiMethods) CreateOne(resp http.ResponseWriter, req *http.Request) {
	vars, enc, dec := crudUnmarshall(resp, req)
	kind := vars["kind"]

	// read body and parse into interface{}
	var resource map[string]interface{}
	err := dec.Decode(&resource)

	var respCode int
	var apiResp apiResponse
	if err != nil {
		log.Println(err)
		respCode = http.StatusBadRequest
		apiResp = apiResponse{"malformed json", "", nil}
	} else {
		// set in storage
		id, stoResp := self.s.Create(kind, resource)
		respCode = stoResp.StatusCode
		apiResp = apiResponse{stoResp.Err, id, nil}
	}

	// write response
	crudMarshall(resp, respCode, apiResp, enc)
	return
}

func (self NoAuthApiMethods) ReadOne(resp http.ResponseWriter, req *http.Request) {
	vars, enc, _ := crudUnmarshall(resp, req)
	kind := vars["kind"]
	id := vars["id"]

	// look for resource
	resource, stoResp := self.s.Get(kind, id)
	apiResp := apiResponse{stoResp.Err, "", resource}

	// write response
	crudMarshall(resp, stoResp.StatusCode, apiResp, enc)
	return
}

func (self NoAuthApiMethods) ReadAll(resp http.ResponseWriter, req *http.Request) {
	vars, enc, _ := crudUnmarshall(resp, req)
	kind := vars["kind"]

	// look for resources
	resources, stoResp := self.s.GetAll(kind)
	apiResp := apiResponse{stoResp.Err, "", resources}

	// write response
	crudMarshall(resp, stoResp.StatusCode, apiResp, enc)
	return
}

func (self NoAuthApiMethods) UpdateOne(resp http.ResponseWriter, req *http.Request) {
	vars, enc, dec := crudUnmarshall(resp, req)
	kind := vars["kind"]
	id := vars["id"]

	// read body and parse into interface{}
	var resource map[string]interface{}
	err := dec.Decode(&resource)

	var respCode int
	var apiResp apiResponse
	if err != nil {
		log.Println(err)
		respCode = http.StatusBadRequest
		apiResp = apiResponse{"malformed json", "", nil}
	} else {
		// update resource
		stoResp := self.s.Update(kind, id, resource)
		respCode = stoResp.StatusCode
		apiResp = apiResponse{stoResp.Err, "", nil}
	}

	// write response
	crudMarshall(resp, respCode, apiResp, enc)
	return
}

func (self NoAuthApiMethods) DeleteOne(resp http.ResponseWriter, req *http.Request) {
	vars, enc, _ := crudUnmarshall(resp, req)
	kind := vars["kind"]
	id := vars["id"]

	// delete resource
	stoResp := self.s.Delete(kind, id)
	apiResp := apiResponse{stoResp.Err, "", nil}

	// write response
	crudMarshall(resp, stoResp.StatusCode, apiResp, enc)
	return
}

func (self NoAuthApiMethods) DeleteAll(resp http.ResponseWriter, req *http.Request) {
	vars, enc, _ := crudUnmarshall(resp, req)
	kind := vars["kind"]

	// look for resources
	stoResp := self.s.DeleteAll(kind)
	apiResp := apiResponse{stoResp.Err, "", nil}

	// write response
	crudMarshall(resp, stoResp.StatusCode, apiResp, enc)
	return
}

func (self NoAuthApiMethods) OptionsOne(resp http.ResponseWriter, req *http.Request) {
	h := resp.Header()

	h.Add("Allow", "PUT")
	h.Add("Allow", "GET")
	h.Add("Allow", "DELETE")
	h.Add("Allow", "OPTIONS")

	resp.WriteHeader(http.StatusOK)
	return
}
func (self NoAuthApiMethods) OptionsAll(resp http.ResponseWriter, req *http.Request) {
	h := resp.Header()

	h.Add("Allow", "POST")
	h.Add("Allow", "GET")
	h.Add("Allow", "DELETE")
	h.Add("Allow", "OPTIONS")

	resp.WriteHeader(http.StatusOK)
	return
}
