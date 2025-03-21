package main

import (
	"log"

	"github.com/ChampionBuffalo1/flydist/internal/handlers"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	handlers.AddBroadcastHandlers(n)
	n.Handle("echo", handlers.GetEchoHandler(n))
	n.Handle("generate", handlers.GetUniqueIDHandler(n))

	if err := n.Run(); err != nil {
		log.Fatalf("Failed to run the handler: %v", err)
	}
}
