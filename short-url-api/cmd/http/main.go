package main

import (
	"fmt"
	"net/http"
	shorten "short-url-api/internal/http/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", shorten.Shorten)
	fmt.Println("listen on port :3000")
	http.ListenAndServe(":3000", r)
}
