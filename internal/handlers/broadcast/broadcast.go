package broadcast

import (
	"encoding/json"

	"github.com/ChampionBuffalo1/flydist/internal/constants"
	"github.com/ChampionBuffalo1/flydist/internal/models"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type broadcastMessage struct {
	Message uint64 `json:"message"`
	models.TypeMessage
}

func GetBroadcastHandler(node *maelstrom.Node, state *BroadcastState) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		var body broadcastMessage
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		state.AddMessageRecv(body.Message)
		return node.Reply(msg, models.NewTypeResponse(constants.BroadcastOK))
	}
}
