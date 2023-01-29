package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "static/index.html")
	})
	r.Run()
}
