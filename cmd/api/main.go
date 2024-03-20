package main

import (
	"image-server/internal/webservice"
)

func main() {
	// Initialize API server
	server := webservice.NewServer()
	server.Run()
}
