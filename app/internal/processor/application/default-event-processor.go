package application

import (
	"fmt"
)

type DefaultEventProcessor struct {
}

func NewDefaultEventProcessor() *DefaultEventProcessor {
	return &DefaultEventProcessor{}
}

func (def *DefaultEventProcessor) Execute(defaultEvent interface{}) error {

	fmt.Println("### DEFAULT EVENT PROCESSED ###\n", defaultEvent)

	return nil
}
