package main

import (
	"log"

	"github.com/SpaceSlow/gophkeeper/internal/server"
)

func main() {
	if srv, err := server.NewServer(); err != nil {
		log.Fatalf("Error occured while setup server: %s.\r\nExiting...", err)
	} else if err = srv.Run(); err != nil {
		log.Fatalf("Error occured while server running: %s.\r\nExiting...", err)
	}
}
