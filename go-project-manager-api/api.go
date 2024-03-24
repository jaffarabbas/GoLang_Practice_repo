package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr  string
	store Store
}

func NewApiServer(addr string, store Store) *ApiServer {
	return &ApiServer{addr: addr, store: store}
}

func (s *ApiServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	log.Printf("Starting server on %s", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
