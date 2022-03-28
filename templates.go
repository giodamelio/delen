package main

import "embed"

// Embed all the templates into the binary
//go:embed templates/*
var templates embed.FS
