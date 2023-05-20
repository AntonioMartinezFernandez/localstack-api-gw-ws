package application

import (
	"fmt"
)

type ConnectedEventProcessor struct {
}

func NewConnectedEventProcessor() *ConnectedEventProcessor {
	return &ConnectedEventProcessor{}
}

func (cep *ConnectedEventProcessor) Execute(connectedEvent interface{}, chargerParam string) error {

	fmt.Println("### CONNECTED EVENT PROCESSED ###\n", connectedEvent)
	fmt.Println("'charger' parameter: ", chargerParam)

	return nil
}
