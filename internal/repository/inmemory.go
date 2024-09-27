package repository

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/utils"
	"time"
)

type InMemoryRepository struct {
	Orders       []domain.Order
	Availability []domain.RoomAvailability
}

var Repo = &InMemoryRepository{
	Orders: []domain.Order{},
	Availability: []domain.RoomAvailability{
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 1), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 2), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 3), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 4), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 5), Quota: 0},
	},
}

func (r *InMemoryRepository) AddOrder(newOrder domain.Order) {
	r.Orders = append(r.Orders, newOrder)
}

func (r *InMemoryRepository) updateAvailability(daysToBook []time.Time, unavailableDays map[time.Time]struct{}) {

	for _, dayToBook := range daysToBook {
		for i, availability := range Repo.Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			Repo.Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}
}

func (r *InMemoryRepository) CheckAvailability(daysToBook []time.Time) map[time.Time]struct{} {
	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	r.updateAvailability(daysToBook, unavailableDays)

	return unavailableDays
}