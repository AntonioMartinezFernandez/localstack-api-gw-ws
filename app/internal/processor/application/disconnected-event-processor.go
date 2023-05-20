package application

import (
	"fmt"
)

type DisconnectedEventProcessor struct {
}

func NewDisconnectedEventProcessor() *DisconnectedEventProcessor {
	return &DisconnectedEventProcessor{}
}

func (dep *DisconnectedEventProcessor) Execute(disconnectedEvent interface{}) error {

	fmt.Println("### DISCONNECTED EVENT PROCESSED ###\n", disconnectedEvent)

	return nil
}
