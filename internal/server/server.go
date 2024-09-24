package server

import (
	v1 "applicationDesignTest/internal/delivery/http/v1"
	"applicationDesignTest/internal/pkg/log"
	"errors"
	"net/http"
	"os"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", v1.CreateOrder)

	log.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.LogInfo("Server closed")
	} else if err != nil {
		log.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
