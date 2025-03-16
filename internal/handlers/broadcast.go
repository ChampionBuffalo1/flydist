package handlers

import (
	"github.com/ChampionBuffalo1/flydist/internal/handlers/broadcast"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func AddBroadcastHandlers(node *maelstrom.Node) {
	state := broadcast.NewBroadcastState()
	node.Handle("read", broadcast.GetReadHandler(node, state))
	node.Handle("topology", broadcast.GetTopologyHandler(node, state))
	node.Handle("broadcast", broadcast.GetBroadcastHandler(node, state))
}
