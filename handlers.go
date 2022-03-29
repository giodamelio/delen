package main

import (
	"net/http"

	"github.com/giodamelio/delen/models"
	"github.com/sanity-io/litter"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	items, err := models.Items().AllG(r.Context())
	if err != nil {
		panic(err)
	}

	renderIndex(w, items)
}

func handleGetUpload(w http.ResponseWriter, r *http.Request) {
	renderUpload(w)
}

func handlePostUploadText(w http.ResponseWriter, r *http.Request) {
	// Max memory before temporary files are used = 10MiB
	r.ParseMultipartForm(1024 * 1024 * 10)
	litter.Dump(r.Form)

	// Create a new item
	var newItem models.Item
	newItem.Name = r.FormValue("name")
	newItem.Contents = []byte(r.FormValue("contents"))
	newItem.MimeType = "text/plain"

	err := newItem.InsertG(r.Context(), boil.Infer())
	if err != nil {
		renderError(w, err)
		return
	}

	renderUploadResult(w)
}

func handlePostUploadFile(w http.ResponseWriter, r *http.Request) {
	renderUploadResult(w)
}
