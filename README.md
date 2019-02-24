# Go study

- personal study of Go lang.
- Go言語の勉強用のレポジトリです

## samples / サンプル

### ws_echo

- Simple echo servert with WebSocket
  - http server is not included. Please copy www/index.html to your web server.
- WebSocketを使ったシンプルなエコーサーバー
  - http サーバー機能は省いています。www/index.htmlを自分のWebサーバーにコピーしてください

### ws_broadcast_data

- Simple broadcast servert with WebSocket
  - data will not send back to sender client.
  - http server is not included. Please copy www/chat.html to your web server.
- WebSocketを使ったシンプルなブロードキャストサーバー
  - データーは送信者には戻ってこないようにしています。送信者以外の接続クライアントに送信されます
  - httpサーバー機能は省いています。www/chat.htmlを自分のWebサーバーにコピーしてください

### ws_broadcast_json

- Simple broadcast servert with WebSocket
  - data will not send back to sender client.
  - use Json struct
  - http server is not included. Please copy www/chat.html to your web server.
- WebSocketを使ったシンプルなブロードキャストサーバー
  - データーは送信者には戻ってこないようにしています。
  - JSONに相当するstructを定義して使っています
  - httpサーバー機能は省いています。www/chat.htmlを自分のWebサーバーにコピーしてください

### ws_gorutine

- Simple broadcast servert with WebSocket
  - use Json struct
  - use gorutine for broadcasting
  - http server is not included. Please copy www/chat.html to your web server.
- WebSocketを使ったシンプルなブロードキャストサーバー
  - JSONに相当するstructを定義して使っています
  - goroutineを使って、送信を行っています
  - httpサーバー機能は省いています。www/chat.htmlを自分のWebサーバーにコピーしてください




## License / ライセンス

* This repository is provided under the MIT license
* このレポジトリはMITランセンスで提供されます

