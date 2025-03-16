package handlers

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type broadcastMessage struct {
}

func GetBroadcastHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		return nil
	}
}
