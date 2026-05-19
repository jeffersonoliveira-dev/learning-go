package main

import (
	"log"
	"net/http"
	"short-url-api/internal/database/sql"
	"short-url-api/internal/http/handlers/shorten"
)

func main() {
	db, err := sql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shortenHandler := shorten.NewHandler(db)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", shortenHandler.Shorten)
	mux.HandleFunc("GET /{code}", shortenHandler.GetUnshorten)

	log.Println("server running on :8080")
	http.ListenAndServe(":8080", mux)
}
