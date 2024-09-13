package main

import (
	"log"

	"github.com/SpaceSlow/gophkeeper/internal/server"
)

func main() {
	if srv, err := server.NewServer(); err != nil || srv.Run() != nil {
		log.Fatalf("Error occured: %s.\r\nExiting...", err)
	}
}
