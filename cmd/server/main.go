package main

import (
	"net/http"
	"os"

	"applicationDesignTest/internal/delivery"
	"applicationDesignTest/internal/repository"
	"applicationDesignTest/internal/usecase"
)

func main() {
	mux := http.NewServeMux()
	repo := repository.NewBookingRepository()
	service := usecase.NewBookingService(repo)
	handler := delivery.NewBookingHandler(service)

	mux.HandleFunc("/orders", handler.CreateOrder)

	LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
