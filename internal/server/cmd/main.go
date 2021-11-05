package main

import (
	"test3/internal/server"
)

const PATH = "/websocket"
const PORT = "8080"

func main() {
	server.StartServer(PORT, PATH)
}
