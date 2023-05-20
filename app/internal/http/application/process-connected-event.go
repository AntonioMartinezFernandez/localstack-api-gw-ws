package application

import (
	http_domain "app/internal/http/domain"
	processor_domain "app/internal/processor/domain"
	"fmt"
)

func ProcessConnectedEvent(connectedRequest interface{}, chargerParam string, ce processor_domain.ConnectedUsecase) http_domain.DataResponse {
	fmt.Println("Processing connected event...")

	err := ce.Execute(connectedRequest, chargerParam)

	res := http_domain.DataResponse{Data: err, Message: "Connected Event Processed"}
	return res
}
