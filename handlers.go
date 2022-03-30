package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/giodamelio/delen/models"
	"github.com/go-chi/chi/v5"
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

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	items, err := models.Items().AllG(r.Context())
	if err != nil {
		renderError(w, err)
		return
	}

	renderItems(w, items)
}

func handlePostItems(w http.ResponseWriter, r *http.Request) {
	// Max memory before temporary files are used = 10MiB
	r.ParseMultipartForm(1024 * 1024 * 10)

	switch r.FormValue("type") {
	case "text":
		// Create a new item
		var newItem models.Item
		newItem.Name = r.FormValue("name")
		newItem.Contents = []byte(r.FormValue("contents"))
		newItem.MimeType = "text/plain" // Default to text/plain since none can currently be specified

		err := newItem.InsertG(r.Context(), boil.Infer())
		if err != nil {
			renderError(w, err)
			return
		}

		renderUploadResult(w)
	case "file":
		renderError(w, fmt.Errorf("file upload currently not implemented"))
	default:
		renderError(w, fmt.Errorf("item type must be one of 'text' and 'file'"))
	}
}

func handleDeleteItems(w http.ResponseWriter, r *http.Request) {
	// Delete the item
	id_string := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_string, 10, 64)
	if err != nil {
		renderError(w, err)
		return
	}
	models.Items(models.ItemWhere.ID.EQ(id)).DeleteAllG(r.Context())

	// Rerender without that item
	items, err := models.Items().AllG(r.Context())
	if err != nil {
		panic(err)
	}

	renderItems(w, items)
}
