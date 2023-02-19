package main

import (
	"sync"

	"gopher-war/lib"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()

	router := lib.Router{
		R: r,
		M: m,
	}
	router.New()

	gameHandler := lib.GameHandler{
		Lock:     new(sync.Mutex),
		M:        m,
		Counter:  0,
		Targets:  make(map[*melody.Session]*lib.TargetInfo),
		Bombs:    make(map[*melody.Session]*lib.BulletInfo),
		Missiles: make(map[*melody.Session]*lib.BulletInfo),
	}
	gameHandler.New()

	r.Run() // run server default port 8080
}
