package infrastructure

import (
	"fmt"
	"io"

	"net/http"

	http_application "app/internal/http/application"
	processor_domain "app/internal/processor/domain"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request, de processor_domain.DisconnectedUsecase) {
	// Read request body
	requestBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendResponseUnprocessableEntity(w)
	}

	fmt.Println("Handling /disconnect request: \n", string(requestBody))

	// Run usecase
	handlerResponse := http_application.ProcessDisconnectedEvent(string(requestBody), de)
	if handlerResponse.Data != nil {
		fmt.Println("process disconnected handler error: ", handlerResponse)
		SendResponseInternalServerError(w)
	} else {
		SendResponseWithData(w, handlerResponse)
	}

}
