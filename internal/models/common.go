package models

type TypeMessage struct {
	Type string `json:"type"`
}

func NewTypeResponse(typ string) *TypeMessage {
	return &TypeMessage{
		Type: typ,
	}
}
