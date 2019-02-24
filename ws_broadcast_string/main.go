package main

import (
	"log"
	"net/http"

	//"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

var clients = make(map[*websocket.Conn]bool) // 接続されるクライアント

func main() {
	// websockerへのルーティングを紐づけ
	//http.HandleFunc("/", handleConnections)
	http.Handle("/", websocket.Handler(handleConnections))

	// WebSocketサーバーをlocalhostのポート3001で立ち上げる
	log.Println("ws server started on :3001")
	err := http.ListenAndServe(":3001", nil)
	// エラーがあった場合ロギングする
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(ws *websocket.Conn) {
	// 関数が終わった際に必ずwebsocketnのコネクションを閉じる
	defer ws.Close()

	// クライアントを新しく登録
	clients[ws] = true

	// ---  Receive / Send string --
	var message string

	for {
		// 文字列メッセージを受け取る
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			log.Printf("error: %v", err)
			return
		}

		// メッセージを同期的に送る
		broadcastStringMessage(ws, message)
	}
}

// 文字列メッセージを送る
func broadcastStringMessage(sender *websocket.Conn, message string) {
	for client := range clients {
		if sender == client {
			log.Printf("skip sender\r\n")
			continue
		}

		err := websocket.Message.Send(client, message)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
