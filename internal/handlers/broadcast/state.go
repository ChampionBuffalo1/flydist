package broadcast

type BroadcastState struct {
	messagesRecv []uint64
	topology     map[string][]string
}

func NewBroadcastState() *BroadcastState {
	return &BroadcastState{}
}

func (s *BroadcastState) GetMessagesRecv() []uint64 {
	return s.messagesRecv
}

func (s *BroadcastState) AddMessageRecv(message uint64) {
	s.messagesRecv = append(s.messagesRecv, message)
}

func (s *BroadcastState) AddTopology(topology map[string][]string) {
	s.topology = topology
}
