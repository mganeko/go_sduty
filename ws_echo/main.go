// simple echo server with websocket

package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)

	/*
		// --- use Receive / Send --
		var message string
		for {
			err := websocket.Message.Receive(ws, &message)
			if err != nil {
				fmt.Printf("error: %v", err)
				break;
			}

			fmt.Printf("receive message: %s\r\n", message)
			err = websocket.Message.Send(ws, message)
			if err != nil {
				fmt.Printf("error: %v", err)
				break;
			}
		}
	*/
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
