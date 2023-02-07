package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

const separator = "\t"

// TargetInfo model
type TargetInfo struct {
	ID     string
	NAME   string
	CHARGE string
	X      int
	Y      int
	LIFE   int
	SIZE   int
}

// BulletInfo model
type BulletInfo struct {
	ID        string
	X         int
	Y         int
	LIFE      int
	MAXLIFE   int
	DIRECTION int
	DAMAGE    int
	SPEED     int
	FIRERANGE int
	SIZE      int
	FIRE      bool
	SPECIAL   bool
}

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

	counter := 0
	targets := make(map[*melody.Session]*TargetInfo)
	bombs := make(map[*melody.Session]*BulletInfo)
	missiles := make(map[*melody.Session]*BulletInfo)
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

	// セッション開始
	m.HandleConnect(func(s *melody.Session) {
		for _, target := range targets {
			message := fmt.Sprintf("show %s %d %d %d %s %s", target.ID, target.X, target.Y, target.LIFE, target.NAME, target.CHARGE)
			s.Write([]byte(message))
		}
		// append
		id := strconv.Itoa(counter)
		targets[s] = &TargetInfo{ID: id, NAME: "", CHARGE: "none"}
		bombs[s] = &BulletInfo{ID: id, SPECIAL: false}
		missiles[s] = &BulletInfo{ID: id, SPECIAL: true}
		message := fmt.Sprintf("appear %s", id)
		s.Write([]byte(message))
		counter++
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
