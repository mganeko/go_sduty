// simple echo server with websocket

package main

import (
    "io"
		"net/http"
		"fmt"

    "golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
    http.Handle("/", websocket.Handler(EchoServer)) // REFUSED
    /*
    http.HandleFunc("/echo",
        func(w http.ResponseWriter, req *http.Request) {
            s := websocket.Server{Handler: websocket.Handler(echoHandler)}
            s.ServeHTTP(w, req)
        }
    )
		*/
		fmt.Printf("start ws server localhost:3001\r\n")
    err := http.ListenAndServe(":3001", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}