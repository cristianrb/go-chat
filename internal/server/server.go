package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan ChatMessage)
var upgrader = websocket.Upgrader{}

type ChatMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func StartServer(port string, path string) {
	http.HandleFunc(path, handleConnections)
	go handleMessage()

	serverUrl := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(serverUrl, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		fmt.Println(msg.Message)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcaster <- msg
	}
}

func handleMessage() {
	for {
		msg := <-broadcaster
		sendMsgToEveryClient(msg)
	}
}

func sendMsgToEveryClient(msg ChatMessage) {
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
