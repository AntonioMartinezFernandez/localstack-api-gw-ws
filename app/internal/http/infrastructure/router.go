package infrastructure

import (
	processor_domain "app/internal/processor/domain"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter(ce processor_domain.ConnectedUsecase, de processor_domain.DisconnectedUsecase, def processor_domain.DefaultUsecase) *mux.Router {
	gorillaMuxRouter := mux.NewRouter()

	// Healthcheck
	gorillaMuxRouter.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		HealthCheckHandler(w, r)
	}).Methods("GET")

	// Connect
	gorillaMuxRouter.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		ConnectHandler(w, r, ce)
	}).Methods("POST")

	// Disconnect
	gorillaMuxRouter.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		DisconnectHandler(w, r, de)
	}).Methods("POST")

	// Default
	gorillaMuxRouter.HandleFunc("/default", func(w http.ResponseWriter, r *http.Request) {
		DefaultHandler(w, r, def)
	}).Methods("POST")

	return gorillaMuxRouter
}
