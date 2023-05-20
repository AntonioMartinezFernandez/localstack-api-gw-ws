package main

import (
	"fmt"

	http_infrastructure "app/internal/http/infrastructure"

	processor_application "app/internal/processor/application"
)

func main() {
	fmt.Println("starting connectivity processor")

	ceProcessor := processor_application.NewConnectedEventProcessor()
	deProcessor := processor_application.NewDisconnectedEventProcessor()
	defProcessor := processor_application.NewDefaultEventProcessor()

	http_infrastructure.StartHttpServer("8000", 15, 15, ceProcessor, deProcessor, defProcessor)
}
