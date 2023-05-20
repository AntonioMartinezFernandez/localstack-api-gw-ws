package infrastructure

import (
	"fmt"
	"io"

	"net/http"

	http_application "app/internal/http/application"
	processor_domain "app/internal/processor/domain"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request, def processor_domain.DefaultUsecase) {
	// Read request body
	requestBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendResponseUnprocessableEntity(w)
	}

	fmt.Println("Handling /default request: \n", string(requestBody))

	// Run usecase
	handlerResponse := http_application.ProcessDefaultEvent(string(requestBody), def)
	if handlerResponse.Data != nil {
		fmt.Println("process default handler error: ", handlerResponse)
		SendResponseInternalServerError(w)
	} else {
		SendResponseWithData(w, handlerResponse)
	}

}
