package server

import (
	delivery "applicationDesignTest/internal/delivery/http"
	"applicationDesignTest/internal/pkg/log"
	"errors"
	"net/http"
	"os"
)

func StartServer(handlers *delivery.Handler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handlers.CreateOrder)

	log.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.LogInfo("Server closed")
	} else if err != nil {
		log.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
