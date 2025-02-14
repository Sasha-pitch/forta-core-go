package registry

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"

	"github.com/forta-network/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-network/forta-core-go/utils"
)

var Link = "Link"
var Unlink = "Unlink"

type DispatchMessage struct {
	Message
	ScannerID string `json:"scannerId"`
	AgentID   string `json:"agentId"`
}

func ParseDispatchMessage(msg string) (*DispatchMessage, error) {
	var m DispatchMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	if m.Action != Link && m.Action != Unlink {
		return nil, fmt.Errorf("invalid action for DispatchMessage: %s", m.Action)
	}
	return &m, nil
}

func NewDispatchMessage(evt *contract_dispatch.DispatchLink) *DispatchMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	agentID := utils.Hex(evt.AgentId)
	evtName := Unlink
	if evt.Enable {
		evtName = Link
	}
	return &DispatchMessage{
		Message: Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
		},
		ScannerID: scannerID,
		AgentID:   agentID,
	}
}

func NewAlreadyLinkedDispatchMessage(evt *contract_dispatch.DispatchAlreadyLinked) *DispatchMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	agentID := utils.Hex(evt.AgentId)
	evtName := Unlink
	if evt.Enable {
		evtName = Link
	}
	return &DispatchMessage{
		Message: Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
		},
		ScannerID: scannerID,
		AgentID:   agentID,
	}
}
