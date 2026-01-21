package main

import (
	"time"

	"github.com/ArthurBitt/Client-Server-API-Challenge/client"
	"github.com/ArthurBitt/Client-Server-API-Challenge/server"
)

func main() {
	server.InitDB()

	go server.StartServer()

	// pequeno delay só para garantir que o servidor subiu
	time.Sleep(200 * time.Millisecond)

	client.Run()
}
