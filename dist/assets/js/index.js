define(["require", "exports", "./WebSocketClient", "./config"], function (require, exports, WebSocketClient_1, config_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    (function main() {
        const url = "ws://" + window.location.host + "/ws";
        const ws = new WebSocket(url); // WebSocket接続開始
        const client = new WebSocketClient_1.WebSocketClient(ws);
        client.addOnOpen();
        client.addOnMessage();
        // WebSocket接続が確立されるまで試行する
        const timerId = setInterval(() => {
            if (client.isOpened()) {
                const initMsg = {
                    type: 'init',
                    userName: 'testUser',
                    config: JSON.stringify(config_1.config)
                };
                client.sendMessage(initMsg);
                clearInterval(timerId);
            }
        }, 25);
    })();
});
