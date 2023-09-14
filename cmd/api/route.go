package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
func routes()
Using chi/v5 Router mux to handle requests
*/
func (server *GQServer) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", serve.HomeHandler)
	mux.Post("/graphql", serve.GQHandler)

	return mux
}
