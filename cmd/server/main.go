package main

import (
	"log"

	"github.com/SpaceSlow/gophkeeper/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Fatalf("Error occured while setup server: %s.\r\nExiting...", err)
	}
}
