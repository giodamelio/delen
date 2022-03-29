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

// Helper to build render functions
func renderPage(filename string, w io.Writer, data any) error {
	template, err := template.ParseFS(templateFS, "templates/base.html", "templates/"+filename)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return err
	}

	err = template.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return err
	}

	return nil
}

func renderIndex(w io.Writer) error {
	return renderPage("index.html", w, nil)
}
