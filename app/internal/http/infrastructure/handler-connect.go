package infrastructure

import (
	"fmt"
	"io"

	"net/http"

	http_application "app/internal/http/application"
	processor_domain "app/internal/processor/domain"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request, ce processor_domain.ConnectedUsecase) {
	// Read request body
	requestBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendResponseUnprocessableEntity(w)
	}

	chargerParam := r.FormValue("charger")

	fmt.Println("Handling /connect request: \n", string(requestBody))

	// Run usecase
	handlerResponse := http_application.ProcessConnectedEvent(string(requestBody), chargerParam, ce)
	if handlerResponse.Data != nil {
		fmt.Println("process connected handler error: ", handlerResponse)
		SendResponseInternalServerError(w)
	} else {
		SendResponseWithData(w, handlerResponse)
	}

}
