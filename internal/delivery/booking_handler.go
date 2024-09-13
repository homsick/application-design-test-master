package delivery

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/usecase"
	"encoding/json"
	"net/http"
)

type BookingHandler struct {
	service *usecase.BookingService
}

func NewBookingHandler(service *usecase.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder domain.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateOrder(newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}
