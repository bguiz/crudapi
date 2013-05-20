package crudapi

import (
	"encoding/json"
	"net/http"
)

type ApiMethods interface {
	// CreateOne(resp http.ResponseWriter, req *http.Request)

	// ReadOne(resp http.ResponseWriter, req *http.Request)
	// ReadAll(resp http.ResponseWriter, req *http.Request)

	// UpdateOne(resp http.ResponseWriter, req *http.Request)

	// DeleteOne(resp http.ResponseWriter, req *http.Request)
	// DeleteAll(resp http.ResponseWriter, req *http.Request)

	// OptionsOne(resp http.ResponseWriter, req *http.Request)
	// OptionsAll(resp http.ResponseWriter, req *http.Request)

	CrudCall(crudMethod func(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)) (httpMethod func(resp http.ResponseWriter, req *http.Request))

	CreateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	ReadOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)
	ReadAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	UpdateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	DeleteOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)
	DeleteAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	OptionsOne(resp http.ResponseWriter, req *http.Request)
	OptionsAll(resp http.ResponseWriter, req *http.Request)
}
