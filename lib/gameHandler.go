package lib

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/olahol/melody"
)

const separator = "\t"

// 機体情報
type TargetInfo struct {
	ID     string // 機体ID(セッションID)
	NAME   string // プレイヤー名
	CHARGE string // 黄色枠、赤枠、枠なしの状態
	X      int    // X座標
	Y      int    // Y座標
	LIFE   int    // 機体のLIFE
	SIZE   int    // 機体の描画サイズ
}

// 弾薬情報
type BulletInfo struct {
	ID        string // 機体ID(セッションID)
	X         int    // X座標
	Y         int    // Y座標
	LIFE      int    // 残り飛距離
	MAXLIFE   int    // 最大飛距離
	DIRECTION int    // 方向
	DAMAGE    int    // ダメージ
	SPEED     int    // 1LIFEで進むpx
	FIRERANGE int    // 安全地帯脱出距離
	SIZE      int    // 描画サイズ
	FIRE      bool   // 射出中(true)、待機中のフラグ(false)
	SPECIAL   bool   // ミサイル(true)、通常弾(false)
}

// ゲームのパラメータ情報
type Config struct {
	MaxLife    int `json:"maxLife"` // 機体の最大ライフ
	TargetSize int `json:"maxSize"` // 機体の最大サイズ(px)

	BombLife  int `json:"bombLife"`  // ボムの最大生存時間
	BombSpeed int `json:"bombSpeed"` // ボムの1LIFE当たりの移動距離(px)
	BombFire  int `json:"bombFire"`  // ボムの当たり判定発生距離
	BombSize  int `json:"bombSize"`  // ボムのサイズ(px)
	BombDmg   int `json:"bombDmg"`   // ボムのダメージ

	MissileLife  int `json:"missileLife"`  // ミサイルの最大生存時間
	MissileSpeed int `json:"missileSpeed"` // ミサイルの1LIFE当たりの移動距離(px)
	MissileFire  int `json:"missileFire"`  // ミサイルの当たり判定発生距離
	MissileSize  int `json:"missileSize"`  // ミサイルのサイズ(px)
	MissileDmg   int `json:"missileDmg"`   // ミサイルのダメージ
	DmgSize      int `json:"dmgSize"`      // ヒット時の機体の縮小サイズ
}

type GameHandler struct {
	Lock     *sync.Mutex
	M        *melody.Melody
	Counter  int
	Targets  map[*melody.Session]*TargetInfo
	Bombs    map[*melody.Session]*BulletInfo
	Missiles map[*melody.Session]*BulletInfo
}

func (g *GameHandler) handleconnect() {
	g.M.HandleConnect(func(s *melody.Session) {
		g.Lock.Lock()
		defer g.Lock.Unlock()

		id := strconv.Itoa(g.Counter)
		target := TargetInfo{ID: id, NAME: "", CHARGE: "none"}
		g.Targets[s] = &target
		bomb := BulletInfo{ID: id, SPECIAL: false}
		g.Bombs[s] = &bomb
		missile := BulletInfo{ID: id, SPECIAL: true}
		g.Missiles[s] = &missile

		message := fmt.Sprintf("My ID %s", id)
		s.Write([]byte(message))

		// TODO: フロントとの繋ぎ込みが必要
		// message = fmt.Sprintf("appear %s", g.targets[s].ID)
		// s.Write([]byte(message))

		g.Counter++
	})
}

func (g *GameHandler) handlemessage() {
	g.M.HandleMessage(func(s *melody.Session, msg []byte) {
		g.Lock.Lock()
		defer g.Lock.Unlock()

		params := strings.Split(string(msg), separator)
		for i, param := range params {
			fmt.Printf("[debug] params[%d] => %T %#v\n", i, param, param)
		}

		switch msgType := params[0]; msgType {
		case "init":
			config := Config{}
			// mapping JSON to Go
			err := json.Unmarshal([]byte(params[2]), &config)
			if err != nil {
				fmt.Printf("Failed to configure by json %s", params[2])
				panic(config)
			}

			target := g.Targets[s]
			target.NAME = params[1]
			target.LIFE = config.MaxLife
			target.SIZE = config.TargetSize
			bomb := g.Bombs[s]
			bomb.MAXLIFE = config.BombLife
			bomb.LIFE = config.BombLife
			bomb.FIRERANGE = config.BombFire
			bomb.SPEED = config.BombSpeed
			bomb.SIZE = config.BombSize
			bomb.DAMAGE = config.BombDmg
			missile := g.Missiles[s]
			missile.MAXLIFE = config.MissileLife
			missile.LIFE = config.MissileLife
			missile.FIRERANGE = config.MissileFire
			missile.SPEED = config.MissileSpeed
			missile.SIZE = config.MissileSize
			missile.DAMAGE = config.MissileDmg
		}
	})
}

func (g *GameHandler) New() {
	g.handleconnect()
	g.handlemessage()
}
