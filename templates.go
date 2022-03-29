package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"

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
		"index":  {"base.html", "index.html"},
		"upload": {"upload.html"},
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
