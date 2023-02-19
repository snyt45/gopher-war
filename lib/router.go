package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type Router struct {
	R *gin.Engine
	M *melody.Melody
}

func (r *Router) static() {
	r.R.Static("/assets", "./dist/assets")
	r.R.LoadHTMLFiles("dist/index.html")
}

func (r *Router) router() {
	r.R.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.R.GET("/ws", func(c *gin.Context) {
		r.M.HandleRequest(c.Writer, c.Request)
	})
}

func (r *Router) New() {
	r.static()
	r.router()
}
