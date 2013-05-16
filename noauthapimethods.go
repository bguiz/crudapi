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

func (self NoAuthApiMethods) CreateOne(resp http.ResponseWriter, req *http.Request) {
}

func (self NoAuthApiMethods) ReadOne(resp http.ResponseWriter, req *http.Request) {
  return
}
func (self NoAuthApiMethods) ReadAll(resp http.ResponseWriter, req *http.Request) {
  return
}

func (self NoAuthApiMethods) UpdateOne(resp http.ResponseWriter, req *http.Request) {
  return
}

func (self NoAuthApiMethods) DeleteOne(resp http.ResponseWriter, req *http.Request) {
  return
}
func (self NoAuthApiMethods) DeleteAll(resp http.ResponseWriter, req *http.Request) {
  return
}

func (self NoAuthApiMethods) OptionsOne(resp http.ResponseWriter, req *http.Request) {
  return
}
func (self NoAuthApiMethods) OptionsAll(resp http.ResponseWriter, req *http.Request) {
  return
}
