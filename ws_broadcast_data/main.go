package main

import (
    "log"
    "net/http"

		"github.com/gorilla/websocket"
		//"golang.org/x/net/websocket"
)

var clients = make(map[*websocket.Conn]bool) // 接続されるクライアント

// アップグレーダ
var upgrader = websocket.Upgrader{}

func main() {
		// websockerへのルーティングを紐づけ
    //http.HandleFunc("/ws", handleConnections)
    http.HandleFunc("/", handleConnections)

    // WebSocketサーバーをlocalhostのポート3001で立ち上げる
    log.Println("ws server started on :3001")
    err := http.ListenAndServe(":3001", nil)
    // エラーがあった場合ロギングする
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
		// 送られてきたGETリクエストをwebsocketにアップグレード
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true		// CROS ORIGIN を許す
		}
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    // 関数が終わった際に必ずwebsocketnのコネクションを閉じる
    defer ws.Close()

    // クライアントを新しく登録
    clients[ws] = true

    for {
				// メッセージを受け取る
				messageType, p, err := ws.ReadMessage()
				if err != nil {
						log.Println(err)
						return
				}

				// メッセージを同期的に送る
				broadcastMessage(ws, messageType, p)
    }
}

func broadcastMessage(sender *websocket.Conn, messageType int, data []byte) {
	for client := range clients {
		if (sender == client) {
			log.Printf("skip sender\r\n")
			continue
		}

		err := client.WriteMessage(messageType, data)
		if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
		}
	}
}