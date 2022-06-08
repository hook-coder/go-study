package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		//err为空的话收到数据 如果err不为空的话 websoket无法接收数据
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		//客户端发来的数据
		fmt.Println("Received back from client: " + reply)
		//发回给客户端的数据
		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)
		//err为空时发送数据过去，如果为空的话不能发送
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
