package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:      PortID,
		MessageList: []Message{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in message
	messageIdMap := make(map[uint64]bool)
	messageCount := gs.GetMessageCount()
	for _, elem := range gs.MessageList {
		if _, ok := messageIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for message")
		}
		if elem.Id >= messageCount {
			return fmt.Errorf("message id should be lower or equal than the last id")
		}
		messageIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
