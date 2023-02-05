define(["require", "exports", "./config"], function (require, exports, config_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    (function main() {
        (0, config_1.testFunc)();
        var url = "ws://" + window.location.host + "/ws";
        // WebSocket接続を作成(/wsにリクエストが飛ぶ)
        var ws = new WebSocket(url);
        // 接続通知
        ws.onopen = function () {
            console.log("WebSocket connected");
            // サーバーにメッセージ送信
            ws.send("test");
        };
        // メッセージ受信
        ws.onmessage = function (event) {
            console.log("WebSocket receive message");
            console.log(event);
        };
    })();
});
