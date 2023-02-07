package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

const separator = "\t"

// Config model
type Config struct {
	MaxLife      int `json:"maxLife"`
	TargetSize   int `json:"maxSize"`
	BombLife     int `json:"bombLife"`
	BombSpeed    int `json:"bombSpeed"`
	BombFire     int `json:"bombFire"`
	BombSize     int `json:"bombSize"`
	BombDmg      int `json:"bombDmg"`
	MissileLife  int `json:"missileLife"`
	MissileSpeed int `json:"missileSpeed"`
	MissileFire  int `json:"missileFire"`
	MissileSize  int `json:"missileSize"`
	MissileDmg   int `json:"missileDmg"`
	DmgSize      int `json:"dmgSize"`
}

func main() {
	r := gin.Default()
	m := melody.New()

	config := Config{}

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
		params := strings.Split(string(msg), separator)
		fmt.Printf("%#v\n", params)
		err := json.Unmarshal([]byte(string(params[2])), &config)
		if err != nil {
			message := fmt.Sprintf("Failed to configure by json [%s]", string(msg))
			fmt.Println(message)
			panic(message)
		}
		fmt.Printf("%#v\n", config)
	})

	r.Run()
}
