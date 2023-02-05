import { testFunc } from './config'

(function main(){
  testFunc();
  const url = "ws://" + window.location.host + "/ws"
  // WebSocket接続を作成(/wsにリクエストが飛ぶ)
  const ws = new WebSocket(url)

  // 接続通知
  ws.onopen = function() {
    console.log("WebSocket connected")
    // サーバーにメッセージ送信
    ws.send("test")
  }

  // メッセージ受信
  ws.onmessage = function(event) {
    console.log("WebSocket receive message")
    console.log(event)
  }
})();
