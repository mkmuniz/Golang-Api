package main

import (
	"go/api/configs"
	userhandlers "go/api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", userhandlers.Create)
	r.Get("/", userhandlers.List)

	log.Printf("Server started on port %v", configs.GetServerPort())
	log.Fatal(http.ListenAndServe(configs.GetServerPort(), r))
}
