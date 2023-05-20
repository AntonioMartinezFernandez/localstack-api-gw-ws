package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"time"

	processor_domain "app/internal/processor/domain"
)

func StartHttpServer(port string, writeTimeout int, readtimeout int, ce processor_domain.ConnectedUsecase, de processor_domain.DisconnectedUsecase, def processor_domain.DefaultUsecase) {
	fmt.Println("starting http server at port", port)

	muxRouter := createRouter(ce, de, def)

	srv := &http.Server{
		Handler: muxRouter,
		Addr:    fmt.Sprint("0.0.0.0:", port),
		// Good practice: enforce timeouts for servers you create
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readtimeout) * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
