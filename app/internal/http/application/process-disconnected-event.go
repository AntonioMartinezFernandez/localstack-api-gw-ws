package application

import (
	http_domain "app/internal/http/domain"
	processor_domain "app/internal/processor/domain"
	"fmt"
)

func ProcessDisconnectedEvent(disconnectedRequest interface{}, de processor_domain.DisconnectedUsecase) http_domain.DataResponse {
	fmt.Println("Processing disconnected event...")

	err := de.Execute(disconnectedRequest)

	res := http_domain.DataResponse{Data: err, Message: "Disconnected Event Processed"}
	return res
}
