package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// HandleGetUrlOriginal handles getting original urls from shortened
func HandleGetURLOriginal(w http.ResponseWriter, r *http.Request) {
	shortener := context.Get(r, "shortener").(*urlShortener)

	params := mux.Vars(r)
	short, found := params["short"]

	if !found {
		http.Error(w, "wrong request", http.StatusBadRequest)
		return
	}

	original, err := shortener.Original(short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, original, 302)
}
