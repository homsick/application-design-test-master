package repository

import (
	"applicationDesignTest/internal/domain"
	"time"
)

type InMemoryRepository struct {
	Orders       []domain.Order
	Availability []domain.RoomAvailability
}

var Repo = &InMemoryRepository{
	Orders:       []domain.Order{},
	Availability: []domain.RoomAvailability{},
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

func (r *InMemoryRepository) AddRoomAvailability(roomAvailability []domain.RoomAvailability) {
	r.Availability = append(r.Availability, roomAvailability...)
}
