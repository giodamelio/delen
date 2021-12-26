package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/elnormous/contenttype"
	"github.com/gin-gonic/gin"
	"github.com/giodamelio/delen/db"
	_ "github.com/lib/pq"
)

func setupDb() (*db.Queries, error) {
	dbConn, err := sql.Open("postgres", "user=postgres dbname=app sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db.New(dbConn), nil
}

// Add the Accepted header to the request context
func AcceptHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Format: %s, length: %d\n", c.GetString("format"), len(c.GetString("format")))
		if c.GetString("format") != "" {
			log.Println("AcceptHeader middleware already run")
			c.Next()
			return
		}

		// Set the format based on the Accept header
		availableMediaTypes := []contenttype.MediaType{
			contenttype.NewMediaType("text/html"),
			contenttype.NewMediaType("application/json"),
		}
		accepted, _, _ := contenttype.GetAcceptableMediaType(c.Request, availableMediaTypes)
		c.Set("format", accepted.String())

		c.Next()
	}
}

func main2() {
	queries, err := setupDb()
	if err != nil {
		log.Fatalf("error: %d", err)
	}

	r := gin.Default()

	// Trust no proxies
	r.SetTrustedProxies(nil)

	// Handle the accept header
	r.Use(AcceptHeader())
	r.Use(AcceptHeader())

	// Load the templates
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		format := c.GetString("format")
		log.Printf("Format: %s\n", format)
		if format == "application/json" {
			c.JSON(http.StatusOK, gin.H{"todo": "Not implemented in JSON"})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		}
	})

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
			return
		}
		files := form.File["files[]"]

		ctx := context.Background()
		for _, file := range files {
			log.Println(file.Filename)

			f, err := file.Open()
			if err != nil {
				// TODO: handle this error
			}

			bytes, err := ioutil.ReadAll(f)
			if err != nil {
				// TODO: handle this error
			}

			queries.CreateItem(ctx, db.CreateItemParams{
				Name:     "test",
				Type:     sql.NullString{String: "text/plain", Valid: true},
				Contents: bytes,
			})
		}

		log.Printf("%d files uploaded\n", len(files))

		format := c.GetString("format")
		if format == "application/json" {
			c.JSON(http.StatusOK, gin.H{"todo": "Not implemented in JSON"})
		} else {
			c.String(http.StatusOK, fmt.Sprintf("%d files uploaded\n", len(files)))
		}
	})

	// If the route is not found, check if there is a .json or .html on the end and redirect
	r.NoRoute(func(c *gin.Context) {
		log.Println("NoRoute hit")

		// Potentially override based on the file path extension
		extension := filepath.Ext(c.Request.URL.Path)
		log.Printf("Extension: %s\n", extension)
		if extension == ".json" {
			c.Set("format", "application/json")
			a := strings.TrimSuffix(c.Request.URL.Path, extension)
			log.Printf("New path: %s\n", a)
			c.Request.URL.Path = a
			r.HandleContext(c)
			return
		} else if extension == ".html" {
			c.Set("format", "text/html")
			c.Request.URL.Path = strings.TrimSuffix(c.Request.URL.Path, extension)
			r.HandleContext(c)
			return
		}
	})

	r.Run()
}
