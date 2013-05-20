package crudapi

import (
	"encoding/json"
	"net/http"
)

type ApiMethods interface {
	CreateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	ReadOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)
	ReadAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	UpdateOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	DeleteOne(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)
	DeleteAll(vars map[string]string, dec *json.Decoder) (respCode int, apiResp apiResponse)

	OptionsOne(resp http.ResponseWriter, req *http.Request)
	OptionsAll(resp http.ResponseWriter, req *http.Request)
}
