package main

import (
	"log"

	"github.com/SpaceSlow/gophkeeper/internal"
)

func main() {
	if err := internal.RunClient(); err != nil {
		log.Fatalf("Error running program: %s", err)
	}
}
