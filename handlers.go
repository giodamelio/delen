package main

import (
	"net/http"

	"github.com/giodamelio/delen/models"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	items, err := models.Items().AllG(r.Context())
	if err != nil {
		panic(err)
	}

	renderIndex(w, items)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	renderUpload(w)
}
