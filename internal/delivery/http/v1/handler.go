package v1

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/repository"
	"applicationDesignTest/internal/service"
	"encoding/json"
	"net/http"
)

var inmemoryrepository = &repository.InMemoryRepository{}

var OrderService = service.NewOrderService(inmemoryrepository)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder domain.Order
	json.NewDecoder(r.Body).Decode(&newOrder)

	createdOrder, err := OrderService.CreateOrder(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)

	log.LogInfo("Order successfully created: %v", createdOrder)
}
