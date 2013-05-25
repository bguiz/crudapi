package crudapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type DefaultApiMethods struct {
	s Storage
	g Guard
}

func NewDefaultApiMethods(store Storage, guard Guard) DefaultApiMethods {
	if guard == nil {
		guard = noopGuard{}
	}
	return DefaultApiMethods{store, guard}
}

func (self DefaultApiMethods) CreateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
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

func (self DefaultApiMethods) ReadOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]
	id := vars["id"]

	// look for resource
	resource, stoResp := self.s.Get(kind, id)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", resource}
	return
}

func (self DefaultApiMethods) ReadAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]

	// look for resources
	resources, stoResp := self.s.GetAll(kind)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", resources}
	return
}

func (self DefaultApiMethods) UpdateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
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

func (self DefaultApiMethods) DeleteOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]
	id := vars["id"]

	// delete resource
	stoResp := self.s.Delete(kind, id)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", nil}
	return
}

func (self DefaultApiMethods) DeleteAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse) {
	kind := vars["kind"]

	// look for resources
	stoResp := self.s.DeleteAll(kind)
	respCode = stoResp.StatusCode
	apiResp = apiResponse{stoResp.Err, "", nil}
	return
}

func (self DefaultApiMethods) OptionsOne(resp http.ResponseWriter, req *http.Request) {
	h := resp.Header()

	h.Add("Allow", "PUT")
	h.Add("Allow", "GET")
	h.Add("Allow", "DELETE")
	h.Add("Allow", "OPTIONS")

	resp.WriteHeader(http.StatusOK)
	return
}
func (self DefaultApiMethods) OptionsAll(resp http.ResponseWriter, req *http.Request) {
	h := resp.Header()

	h.Add("Allow", "POST")
	h.Add("Allow", "GET")
	h.Add("Allow", "DELETE")
	h.Add("Allow", "OPTIONS")

	resp.WriteHeader(http.StatusOK)
	return
}
