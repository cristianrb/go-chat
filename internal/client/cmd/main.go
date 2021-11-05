package main

import (
	"test3/internal/client"
)

const HOST = "localhost:8080"
const PATH = "/websocket"

func main() {
	client := client.NewClient(HOST, PATH)
	client.Start()
	defer client.Stop()
}