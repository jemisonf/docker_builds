package main

import (
	"log"

	"github.com/jemisonf/docker_builds/server"
)

func main() {
	logger := log.Default()
	err := server.StartServer(":3333")

	if err != nil {
		logger.Fatalf("error starting server: %s", err)
	}

	logger.Println("started server on http://localhost:3333")
}
