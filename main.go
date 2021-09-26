package main

import (
	"log"
	"os"

	"github.com/jemisonf/docker_builds/server"
)

func main() {
	logger := log.Default()

	port := ":3333"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	err := server.StartServer(port)

	if err != nil {
		logger.Fatalf("error starting server: %s", err)
	}

	logger.Println("started server on http://localhost:3333")
}
