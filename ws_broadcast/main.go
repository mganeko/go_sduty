package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // 接続されるクライアント

// アップグレーダ
var upgrader = websocket.Upgrader{}

// メッセージ用構造体
type Message struct {
    //Email    string `json:"email"`
    //Username string `json:"username"`
    Message  string `json:"message"`
}

type BroadcastMsg struct {
	sender *websocket.Conn
	msg *Message
}

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
        var msg Message
        // 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピングする
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, ws)
            break
        }
				
				// メッセージを同期的に送る
				broadcastMessage(ws, msg)
    }
}

// --- チャネルを使わずに同期的に送る---
func broadcastMessage(sender *websocket.Conn, msg Message) {
	for client := range clients {
		if (sender == client) {
			log.Printf("skip sender\r\n")
			continue
		}

		err := client.WriteJSON(msg)
		if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
		}
	}
}