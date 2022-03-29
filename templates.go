package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/giodamelio/delen/models"
)

// Embed all the templates into the binary
//go:embed templates/*
var templateFS embed.FS
var templates map[string]*template.Template

// Parse all the templates and store them in a map
func init() {
	templates = make(map[string]*template.Template)

	templateDefinitions := map[string][]string{
		"error":         {"error.html"},
		"index":         {"base.html", "index.html", "section/upload.html", "section/items.html"},
		"upload":        {"section/upload.html"},
		"upload-result": {"section/upload-result.html"},
		"items":         {"section/items.html"},
	}

	for name, templateParts := range templateDefinitions {
		// Put templates/ on the front of each part
		partsWithPrefix := make([]string, len(templateParts))
		for i, part := range templateParts {
			partsWithPrefix[i] = "templates/" + part
		}

		template, err := template.ParseFS(templateFS, partsWithPrefix...)
		// Panic at startup if any of the templates are malformed
		if err != nil {
			panic(err)
		}

		templates[name] = template
	}
}

// Helper to build render functions
func renderPage(filename string, w io.Writer, data any) error {
	err := templates[filename].Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return err
	}

	return nil
}

func renderError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "text/html")
	return renderPage("error", w, err)
}

func renderIndex(w io.Writer, items models.ItemSlice) error {
	return renderPage("index", w, struct {
		Items models.ItemSlice
	}{
		Items: items,
	})
}

func renderUpload(w io.Writer) error {
	return renderPage("upload", w, nil)
}

func renderUploadResult(w io.Writer) error {
	return renderPage("upload-result", w, nil)
}

func renderItems(w io.Writer, items models.ItemSlice) error {
	return renderPage("items", w, items)
}
