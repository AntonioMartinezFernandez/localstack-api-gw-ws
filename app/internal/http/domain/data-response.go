package domain

type DataResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
