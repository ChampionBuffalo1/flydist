package broadcast

import (
	"encoding/json"

	"github.com/ChampionBuffalo1/flydist/internal/constants"
	"github.com/ChampionBuffalo1/flydist/internal/models"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type topologyMessage struct {
	Topology map[string][]string `json:"topology,omitempty"`
	models.TypeMessage
}

func GetTopologyHandler(node *maelstrom.Node, state *BroadcastState) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		var body topologyMessage
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		state.AddTopology(body.Topology)
		return node.Reply(msg, models.NewTypeResponse(constants.TopologyOK))
	}
}
