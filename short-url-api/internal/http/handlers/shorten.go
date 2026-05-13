// Package shorten is responsible for the shorten route
package shorten

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type URLBody struct {
	URL string `json:"url"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	var payload URLBody

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	fmt.Printf("Received: %s \n", payload.URL)

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("url: " + payload.URL))

}
