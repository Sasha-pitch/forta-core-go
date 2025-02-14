package registry

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"

	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/utils"
)

var SaveAgent = "SaveAgent"
var DisableAgent = "DisableAgent"
var EnableAgent = "EnableAgent"

type AgentMessage struct {
	Message
	AgentID string `json:"agentId"`
	TxHash  string `json:"txHash"`
}

type AgentSaveMessage struct {
	AgentMessage
	Enabled  bool    `json:"enabled"`
	Name     string  `json:"name"`
	ChainIDs []int64 `json:"chainIds"`
	Metadata string  `json:"metadata"`
	Owner    string  `json:"owner"`
}

func ParseAgentSave(msg string) (*AgentSaveMessage, error) {
	var save AgentSaveMessage
	err := json.Unmarshal([]byte(msg), &save)
	if err != nil {
		return nil, err
	}
	if save.Action != SaveAgent {
		return nil, fmt.Errorf("invalid action for AgentSave: %s", save.Action)
	}
	return &save, nil
}

func ParseAgentMessage(msg string) (*AgentMessage, error) {
	var m AgentMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func NewAgentMessage(evt *contract_agent_registry.AgentRegistryAgentEnabled) *AgentMessage {
	agentID := utils.Hex(evt.AgentId)
	evtName := DisableAgent
	if evt.Enabled {
		evtName = EnableAgent
	}
	return &AgentMessage{
		Message: Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
		},
		AgentID: agentID,
		TxHash:  evt.Raw.TxHash.Hex(),
	}
}

func NewAgentSaveMessage(evt *contract_agent_registry.AgentRegistryAgentUpdated) *AgentSaveMessage {
	agentID := utils.Hex(evt.AgentId)
	return &AgentSaveMessage{
		AgentMessage: AgentMessage{
			AgentID: agentID,
			Message: Message{
				Action:    SaveAgent,
				Timestamp: time.Now().UTC(),
			},
			TxHash: evt.Raw.TxHash.Hex(),
		},
		Enabled:  true,
		Name:     evt.Metadata,
		ChainIDs: utils.IntArray(evt.ChainIds),
		Metadata: evt.Metadata,
		Owner:    evt.By.Hex(),
	}
}
