package handlers

import (
	"encoding/json"

	"github.com/ChampionBuffalo1/flydist/internal/constants"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type echoMessage struct {
	ID   int    `json:"msg_id"`
	Type string `json:"type"`
	Echo string `json:"echo"`
}

// echo dist sys challenge: https://fly.io/dist-sys/1
func GetEchoHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		var body echoMessage
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body.Type = constants.EchoOK
		return node.Reply(msg, body)
	}
}
