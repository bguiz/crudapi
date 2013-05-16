package crudapi

import (
  "net/http"
)

type ApiMethods interface {

  CreateOne(resp http.ResponseWriter, req *http.Request)

  ReadOne(resp http.ResponseWriter, req *http.Request)
  ReadAll(resp http.ResponseWriter, req *http.Request)

  UpdateOne(resp http.ResponseWriter, req *http.Request)

  DeleteOne(resp http.ResponseWriter, req *http.Request)
  DeleteAll(resp http.ResponseWriter, req *http.Request)

  OptionsOne(resp http.ResponseWriter, req *http.Request)
  OptionsAll(resp http.ResponseWriter, req *http.Request)
}