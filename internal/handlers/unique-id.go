package handlers

import (
	"log"

	"github.com/ChampionBuffalo1/flydist/internal/constants"
	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type uniqueIdResponse struct {
	ID   uint32 `json:"id"`
	Type string `json:"type"`
}

// generate unique ID challenge: https://fly.io/dist-sys/2
func GetUniqueIDHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		var generateResponse uniqueIdResponse
		generateResponse.Type = constants.GenerateOK
		if uid, err := uuid.NewRandom(); err == nil {
			generateResponse.ID = uid.ID()
		} else {
			log.Fatalf("Failed to generate unique ID: %v", err)
		}
		return node.Reply(msg, generateResponse)
	}
}
