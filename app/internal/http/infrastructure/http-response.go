package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	http_domain "app/internal/http/domain"
)

type HttpResponse struct {
	Status         int                 `json:"status"`
	Data           interface{}         `json:"data"`
	Message        string              `json:"message"`
	ContentType    string              `json:"contentType"`
	ResponseWriter http.ResponseWriter `json:"responseWriter"`
}

func CreateDefaultHttpResponse(w http.ResponseWriter) HttpResponse {
	return HttpResponse{
		Status:         http.StatusOK,
		ResponseWriter: w,
		ContentType:    "application/json",
	}
}

func (httpResponse *HttpResponse) SendWithData() {
	httpResponse.ResponseWriter.Header().Set("Content-Type", httpResponse.ContentType)
	httpResponse.ResponseWriter.WriteHeader(httpResponse.Status)

	output, _ := json.Marshal(&httpResponse.Data)

	res := string(output)

	if res == "null" {
		fmt.Fprint(httpResponse.ResponseWriter, httpResponse.Message)
	} else {
		fmt.Fprint(httpResponse.ResponseWriter, res)
	}
}

func SendResponseWithData(w http.ResponseWriter, data http_domain.DataResponse) {
	response := CreateDefaultHttpResponse(w)
	response.Data = data
	response.SendWithData()
}

func SendResponseHealthcheck(w http.ResponseWriter, data http_domain.HealthcheckResponse) {
	response := CreateDefaultHttpResponse(w)
	response.Data = data
	response.SendWithData()
}

func (httpResponse *HttpResponse) SendNotFound() {
	httpResponse.Status = http.StatusNotFound
	httpResponse.Message = "Not Found"
	http.Error(httpResponse.ResponseWriter, httpResponse.Message, httpResponse.Status)
}

func SendResponseNotFound(w http.ResponseWriter) {
	response := CreateDefaultHttpResponse(w)
	response.SendNotFound()
}

func (httpResponse *HttpResponse) SendUnprocessableEntity() {
	httpResponse.Status = http.StatusUnprocessableEntity
	httpResponse.Message = "Unprocessable Entity"
	http.Error(httpResponse.ResponseWriter, httpResponse.Message, httpResponse.Status)
}

func SendResponseUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultHttpResponse(w)
	response.SendUnprocessableEntity()
}

func (httpResponse *HttpResponse) SendInternalServerError() {
	httpResponse.Status = http.StatusInternalServerError
	httpResponse.Message = "Internal Server Error"
	http.Error(httpResponse.ResponseWriter, httpResponse.Message, httpResponse.Status)
}

func SendResponseInternalServerError(w http.ResponseWriter) {
	response := CreateDefaultHttpResponse(w)
	response.Status = http.StatusInternalServerError
	response.SendInternalServerError()
}
