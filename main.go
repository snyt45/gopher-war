package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

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

func initRouter(r *gin.Engine, m *melody.Melody) {
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
}

func initHandleRequest(m *melody.Melody) {
	lock := new(sync.Mutex)
	counter := 0
	targets := make(map[*melody.Session]*TargetInfo)
	bombs := make(map[*melody.Session]*BulletInfo)
	missiles := make(map[*melody.Session]*BulletInfo)

	// セッション開始
	m.HandleConnect(func(s *melody.Session) {
		lock.Lock()

		// セッション開始した旨を通知する
		session_start_message := "Session Start! Welcome to gopaher-war"
		s.Write([]byte(session_start_message))

		// 他の機体が参加している場合に他の機体情報を通知する
		for _, target := range targets {
			other_client_info_message := fmt.Sprintf(
				"Other Client Info: ID %s, X %d, Y %d, LIFE %d, NAME %s, CHARGE %s",
				target.ID, target.X, target.Y, target.LIFE, target.NAME, target.CHARGE,
			)
			s.Write([]byte(other_client_info_message))
		}

		// 自機の情報、ボム、ミサイルの情報を詰める
		id := strconv.Itoa(counter)
		targets[s] = &TargetInfo{ID: id, NAME: "", CHARGE: "none"}
		bombs[s] = &BulletInfo{ID: id, SPECIAL: false}
		missiles[s] = &BulletInfo{ID: id, SPECIAL: true}

		// 自機の情報を通知する
		my_client_info_message := fmt.Sprintf(
			"My Client Info: ID %s, X %d, Y %d, LIFE %d, NAME %s, CHARGE %s",
			id, targets[s].X, targets[s].Y, targets[s].LIFE, targets[s].NAME, targets[s].CHARGE,
		)
		s.Write([]byte(my_client_info_message))

		// IDをカウントアップ
		counter++

		lock.Unlock()
	})

	// メッセージ受信
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		lock.Lock()

		// パラメータ分割
		params := strings.Split(string(msg), separator)
		for i, param := range params {
			fmt.Printf("[debug] params[%d] => %T %#v\n", i, param, param)
		}

		switch msgType := params[0]; msgType {
		case "init":
			initConfig(params[1], params[2], s)
		}

		lock.Unlock()
	})
}

func initConfig(userName string, configJson string, s *melody.Session) {
	config := Config{}
	// mapping JSON to Go
	err := json.Unmarshal([]byte(configJson), &config)
	if err != nil {
		fmt.Printf("Failed to configure by json %s", configJson)
		panic(config)
	}
	// Configを設定した旨を通知する
	message := fmt.Sprintf("initConfig: %#v", config)
	s.Write([]byte(message))
}

func main() {
	r := gin.Default()
	m := melody.New()

	initRouter(r, m)
	initHandleRequest(m)

	// run server port 8080
	r.Run()
}
