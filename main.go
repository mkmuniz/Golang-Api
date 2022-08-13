package main

import (
	"fmt"
	"go/api/configs"
	"go/api/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)

	http.ListenAndServe(fmt.Sprintf("%s", configs.GetServerPort()), r)
}
