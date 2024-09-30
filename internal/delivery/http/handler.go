package http

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder domain.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	createdOrder, err := h.services.Order.CreateOrder(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)

	log.LogInfo("Order successfully created: %v", createdOrder)
}
