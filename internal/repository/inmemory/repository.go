package repository

import (
	"applicationDesignTest/internal/domain"
	"time"
)

type Orders interface {
	Add(newOrder domain.Order)
	GetAll() []domain.Order
	GetByID(orderID string) (domain.Order, error)
	Update(order domain.Order) error
	Delete(orderID string) error
}

type RoomsAvailability interface {
	Add(roomAvailability []domain.RoomAvailability)
	Update(availability domain.RoomAvailability) error
	Delete(hotelID, roomID string, date time.Time) error
	GetByDate(hotelID, roomID string, date time.Time) (domain.RoomAvailability, error)
	CheckAvailability(daysToBook []time.Time) map[time.Time]struct{}
}

type Repositories struct {
	Orders            Orders
	RoomsAvailability RoomsAvailability
}

func NewRepositories() *Repositories {
	return &Repositories{
		Orders:            NewInMemoryOrdersRepository(),
		RoomsAvailability: NewInMemoryRoomsAvailabilityRepository(),
	}
}
