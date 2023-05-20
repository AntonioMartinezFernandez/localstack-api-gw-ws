package application

import (
	http_domain "app/internal/http/domain"
	processor_domain "app/internal/processor/domain"
	"fmt"
)

func ProcessDefaultEvent(defaultRequest interface{}, def processor_domain.DefaultUsecase) http_domain.DataResponse {
	fmt.Println("Processing default event...")

	err := def.Execute(defaultRequest)

	res := http_domain.DataResponse{Data: err, Message: "Default Event Processed"}
	return res
}
