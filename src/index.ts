import { WebSocketClient, InitMsg } from "./WebSocketClient";
import { config } from "./config";

(function main(){
  const url = "ws://" + window.location.host + "/ws"
  const ws = new WebSocket(url) // WebSocket接続開始
  const client = new WebSocketClient(ws)
  client.addOnOpen()
  client.addOnMessage()

  // WebSocket接続が確立されるまで試行する
  const timerId = setInterval(() => {
    if (client.isOpened()) {
      const initMsg: InitMsg = {
        type: 'init',
        userName: 'testUser',
        config: JSON.stringify(config)
      }
      client.sendMessage(initMsg)
      clearInterval(timerId)
    }
  }, 25)
})();
