// Package shorten is responsible for the shorten route package shorten
package shorten

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"short-url-api/internal/http/handlers/encode"
)

type URLBody struct {
	URL string `json:"url"`
}

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var payload URLBody

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if payload.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	var id uint64

	err = h.DB.QueryRow(`
	INSERT INTO short_urls (original_url)
	VALUES ($1)
	RETURNING id
	`, payload.URL).Scan(&id)

	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	code := encode.Encode(id)

	_, err = h.DB.Exec(`
		UPDATE short_urls
		SET code = $1
		WHERE id = $2
	`, code, id)

	if err != nil {
		http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"original_url": payload.URL,
		"short_code":   code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (h *Handler) GetUnshorten(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	fmt.Println(code)

	var originalURL string

	err := h.DB.QueryRow(`
		SELECT original_url
		FROM short_urls
		WHERE code = $1
	`, code).Scan(&originalURL)

	fmt.Println(err)

	if err == sql.ErrNoRows {
		http.Error(w, "URL not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "Failed to get URL ", http.StatusInternalServerError)
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
