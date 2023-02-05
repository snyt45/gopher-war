package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.Static("/assets", "./dist/assets")
	r.LoadHTMLFiles("dist/index.html")

	// トップページ
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// WebSocket接続を受け取る
	r.GET("/ws", func(c *gin.Context) {
		// WebSocket接続をmelodyインスタンスが処理する
		m.HandleRequest(c.Writer, c.Request)
	})

	// メッセージ受信
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// 全てのセッションにメッセージをブロードキャスト
		m.Broadcast(msg)
	})

	r.Run()
}
