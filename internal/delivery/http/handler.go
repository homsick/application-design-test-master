package http

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	services *service.OrderService
}

func NewHandler(services *service.OrderService) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder domain.Order
	json.NewDecoder(r.Body).Decode(&newOrder)

	createdOrder, err := h.services.CreateOrder(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)

	log.LogInfo("Order successfully created: %v", createdOrder)
}
