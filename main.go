package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		availableMediaTypes := []contenttype.MediaType{
			contenttype.NewMediaType("text/html"),
			contenttype.NewMediaType("application/json"),
		}

		accepted, _, _ := contenttype.GetAcceptableMediaType(c.Request, availableMediaTypes)
		log.Println("Accepted media type:", accepted.String())
		c.Set("accept", accepted.String())

		c.Next()
	}
}

func main() {
	queries, err := setupDb()
	if err != nil {
		log.Fatalf("error: %d", err)
	}

	r := gin.Default()

	// Trust no proxies
	r.SetTrustedProxies(nil)

	// Handle the accept header
	r.Use(AcceptHeader())

	// Load the templates
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		accept := c.GetString("accept")
		if accept == "application/json" {
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

		accept := c.GetString("accept")
		if accept == "application/json" {
			c.JSON(http.StatusOK, gin.H{"todo": "Not implemented in JSON"})
		} else {
			c.String(http.StatusOK, fmt.Sprintf("%d files uploaded\n", len(files)))
		}
	})

	r.Run()
}
