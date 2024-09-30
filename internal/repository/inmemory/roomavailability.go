package repository

import (
	"applicationDesignTest/internal/domain"
	"time"
)

type InMemoryRoomsAvailabilityRepository struct {
	RoomsAvailability []domain.RoomAvailability
}

func NewInMemoryRoomsAvailabilityRepository() *InMemoryRoomsAvailabilityRepository {
	return &InMemoryRoomsAvailabilityRepository{
		RoomsAvailability: []domain.RoomAvailability{},
	}
}

func (r *InMemoryRoomsAvailabilityRepository) Add(roomAvailability []domain.RoomAvailability) {
	r.RoomsAvailability = append(r.RoomsAvailability, roomAvailability...)
}

func (r *InMemoryRoomsAvailabilityRepository) CheckAvailability(daysToBook []time.Time) map[time.Time]struct{} {
	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	r.updateAvailability(daysToBook, unavailableDays)

	return unavailableDays
}

func (r *InMemoryRoomsAvailabilityRepository) updateAvailability(daysToBook []time.Time, unavailableDays map[time.Time]struct{}) {
	for _, dayToBook := range daysToBook {
		for i, availability := range r.RoomsAvailability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			r.RoomsAvailability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}
}

func (r *InMemoryRoomsAvailabilityRepository) Update(availability domain.RoomAvailability) error {
	return nil // TODO: Implement room availability update logic
}
func (r *InMemoryRoomsAvailabilityRepository) Delete(hotelID, roomID string, date time.Time) error {
	return nil // TODO: Implement room availability deletion logic
}

func (r *InMemoryRoomsAvailabilityRepository) GetByDate(hotelID, roomID string, date time.Time) (domain.RoomAvailability, error) {
	return domain.RoomAvailability{}, nil // TODO: Implement room availability retrieval logic
}
