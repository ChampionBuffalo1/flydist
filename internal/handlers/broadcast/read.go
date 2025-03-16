package broadcast

import (
	"github.com/ChampionBuffalo1/flydist/internal/constants"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type readResponse struct {
	Type     string   `json:"type"`
	Messages []uint64 `json:"messages"`
}

func GetReadHandler(node *maelstrom.Node, state *BroadcastState) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		resp := readResponse{
			constants.ReadOK,
			state.GetMessagesRecv(),
		}
		return node.Reply(msg, resp)
	}
}
