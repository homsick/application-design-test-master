// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/pkg/utils"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

var Orders = []domain.Order{}

var Availability = []domain.RoomAvailability{
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 1), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 2), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 3), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 4), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 5), Quota: 0},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", createOrder)

	log.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.LogInfo("Server closed")
	} else if err != nil {
		log.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder domain.Order
	json.NewDecoder(r.Body).Decode(&newOrder)

	daysToBook := utils.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
		log.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	Orders = append(Orders, newOrder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	log.LogInfo("Order successfully created: %v", newOrder)
}
