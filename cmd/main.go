package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Post("/webhook", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("got event"))
	})

	http.ListenAndServe(":8080", router)
}
