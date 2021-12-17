package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Trust no proxies
	r.SetTrustedProxies(nil)

	// Load the templates
	r.LoadHTMLGlob("templates/*")

	// Setup API
	ApiRouter(r.Group("/api"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run()
}

func ApiRouter(r *gin.RouterGroup) {
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
			return
		}
		files := form.File["files[]"]

		for _, file := range files {
			log.Println(file.Filename)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!\n", len(files)))
	})
}
