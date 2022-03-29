package main

import "net/http"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	renderIndex(w)
}
