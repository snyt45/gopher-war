define(["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.WebSocketClient = void 0;
    // よく使われる文字だとサーバー側でパラメータ分割する際に意図しない箇所で分割されるため使われくいタブ文字にしている
    const SEPARATOR = '\t';
    class WebSocketClient {
        constructor(ws) {
            this.ws = ws;
            this.opened = false;
        }
        addOnOpen() {
            this.ws.onopen = () => {
                this.opened = true;
                console.log("WebSocket connected");
            };
        }
        addOnMessage() {
            this.ws.onmessage = () => {
                console.log("WebSocket receive message");
            };
        }
        sendMessage(msg) {
            if (!this.isOpened())
                return;
            console.log("WebSocket send message");
            switch (msg.type) {
                case 'init':
                    const arr = Object.values(msg);
                    this.ws.send(arr.join(SEPARATOR));
                    break;
            }
        }
        isOpened() {
            return this.opened;
        }
    }
    exports.WebSocketClient = WebSocketClient;
});
