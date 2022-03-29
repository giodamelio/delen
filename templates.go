package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

// Embed all the templates into the binary
//go:embed templates/*
var templateFS embed.FS
var templates map[string]*template.Template

// Parse all the templates and store them in a map
func init() {
	templates = make(map[string]*template.Template)

	files, _ := templateFS.ReadDir("templates")
	for _, file := range files {
		if file.Name() == "base.index" {
			continue
		}

		template, err := template.ParseFS(templateFS, "templates/base.html", "templates/"+file.Name())
		// Panic at startup if any of the templates are malformed
		if err != nil {
			panic(err)
		}

		templates[file.Name()] = template
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

func renderIndex(w io.Writer) error {
	return renderPage("index.html", w, nil)
}
