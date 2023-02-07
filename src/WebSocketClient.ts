// よく使われる文字だとサーバー側でパラメータ分割する際に意図しない箇所で分割されるため使われくいタブ文字にしている
const SEPARATOR = '\t'

export type InitMsg = {
  type: 'init',
  userName: string,
  config: string,
}

type Message = InitMsg

export class WebSocketClient {
  ws: WebSocket
  opened: boolean

  constructor(ws: WebSocket) {
    this.ws = ws
    this.opened = false
  }

  addOnOpen() {
    this.ws.onopen = () => {
      this.opened = true
      console.log("WebSocket connected")
    }
  }

  addOnMessage() {
    this.ws.onmessage = () => {
      console.log("WebSocket receive message")
    }
  }

  sendMessage(msg: Message) {
    if (!this.isOpened()) return

    console.log("WebSocket send message")

    switch (msg.type) {
      case 'init':
        const arr = Object.values(msg)
        this.ws.send(arr.join(SEPARATOR))
        break
    }
  }

  isOpened(): boolean {
    return this.opened
  }
}
