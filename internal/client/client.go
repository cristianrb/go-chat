package client

import (
	"bufio"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"test3/internal/server"

	"github.com/gorilla/websocket"
)

type WebsocketClient interface {
	Start()
	Stop()
}

type WebsocketClientImpl struct {
	Conn *websocket.Conn
}

func NewClient(host string, path string) WebsocketClient {
	url := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   path,
	}

	c, _, _ := websocket.DefaultDialer.Dial(url.String(), nil)
	client := WebsocketClientImpl{
		Conn: c,
	}
	return client
}

func (client WebsocketClientImpl) Start() {
	go readMessages(client.Conn)
	listenMessages(client.Conn)
}

func (client WebsocketClientImpl) Stop() {
	client.Conn.Close()
}

func readMessages(c *websocket.Conn) {
	for {
		_, message, _ := c.ReadMessage()
		log.Printf("Message received: %s", message)
		log.Printf("Send new message: ")
	}
}

func listenMessages(c *websocket.Conn) {
	log.Printf("Send new message: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		chatMessage := server.ChatMessage{}
		json.Unmarshal([]byte(message), &chatMessage)
		c.WriteJSON(chatMessage)
	}
}
