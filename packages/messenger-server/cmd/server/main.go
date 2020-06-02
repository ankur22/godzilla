package main

import (
	"context"
	"log"

	"github.com/ankur22/godzilla/packages/messenger-server/internal"
)

func main() {
	if err := internal.Run(context.Background()); err != nil {
		log.Fatalf("Exiting service: %v", err)
	}
}
