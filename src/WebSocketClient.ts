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
      console.log("[WebSocket connected]")
    }
  }

  addOnMessage() {
    this.ws.onmessage = (event: MessageEvent<string>) => {
      console.log("[WebSocket receive message]: ", event.data)
    }
  }

  sendMessage(msg: Message) {
    if (!this.isOpened()) return

    switch (msg.type) {
      case 'init':
        const arr = Object.values(msg)
        this.ws.send(arr.join(SEPARATOR))
        console.log("[WebSocket send message] ", msg)
        break
    }
  }

  isOpened(): boolean {
    return this.opened
  }
}
