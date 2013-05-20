package crudapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type NoAuthApiMethods struct {
	s Storage
}

func NewNoAuthApiMethods(store Storage) NoAuthApiMethods {
	return NoAuthApiMethods{store}
}

func (self NoAuthApiMethods) CreateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]

	// read body and parse into interface{}
	var resource map[string]interface{}
	err := dec.Decode(&resource)

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
	return
}

func (self NoAuthApiMethods) ReadOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]
	id := vars["id"]

	// look for resource
	resource, stoResp := self.s.Get(kind, id)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", resource}
	return
}

func (self NoAuthApiMethods) ReadAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]

	// look for resources
	resources, stoResp := self.s.GetAll(kind)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", resources}
	return
}

func (self NoAuthApiMethods) UpdateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]
	id := vars["id"]

	// read body and parse into interface{}
	var resource map[string]interface{}
	err := dec.Decode(&resource)

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
	return
}

func (self NoAuthApiMethods) DeleteOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]
	id := vars["id"]

	// delete resource
	stoResp := self.s.Delete(kind, id)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", nil}
	return
}

func (self NoAuthApiMethods) DeleteAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]

	// look for resources
	stoResp := self.s.DeleteAll(kind)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", nil}
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
