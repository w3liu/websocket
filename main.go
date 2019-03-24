package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"w3liu.com/websocket/handler"
)

func main() {
	http.Handle("/ws", websocket.Handler(handler.Handle))
	err := http.ListenAndServe(":88", nil)
	if err != nil {
		fmt.Println(err)
	}
}
