package repository

import (
	"applicationDesignTest/internal/domain"
	"errors"
	"time"
)

// Availability represents the interface for room availability management.
type BookingRepository interface {
	CreateOrder(order domain.Order) error
	UpdateAvailability(date time.Time, roomID string, hotelID string) error
	GetAvailability(date time.Time, roomID string, hotelID string) (domain.RoomAvailability, error)
}

type InMemoryBookingRepository struct {
	orders       []domain.Order
	availability []domain.RoomAvailability
}

func NewInMemoryBookingRepository() *InMemoryBookingRepository {
	return &InMemoryBookingRepository{
		orders: []domain.Order{},
		availability: []domain.RoomAvailability{
			{"reddison", "lux", date(2024, 1, 1), 1},
			{"reddison", "lux", date(2024, 1, 2), 1},
			{"reddison", "lux", date(2024, 1, 3), 1},
			{"reddison", "lux", date(2024, 1, 4), 1},
			{"reddison", "lux", date(2024, 1, 5), 0},
		},
	}
}

func (r *InMemoryBookingRepository) CreateOrder(order domain.Order) error {
	r.orders = append(r.orders, order)
	return nil
}

func (r *InMemoryBookingRepository) UpdateAvailability(date time.Time, roomID string, hotelID string) error {
	for i, availability := range r.availability {
		if availability.Date.Equal(date) && availability.HotelID == hotelID && availability.RoomID == roomID {
			if availability.Quota < 1 {
				return errors.New("no rooms available")
			}
			r.availability[i].Quota -= 1
			return nil
		}
	}
	return errors.New("room not available on this date")
}

func (r *InMemoryBookingRepository) GetAvailability(date time.Time, roomID string, hotelID string) (domain.RoomAvailability, error) {
	for _, availability := range r.availability {
		if availability.Date.Equal(date) && availability.HotelID == hotelID && availability.RoomID == roomID {
			return availability, nil
		}
	}
	return domain.RoomAvailability{}, errors.New("no availability found")
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
