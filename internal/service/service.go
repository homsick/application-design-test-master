package service

import (
	"applicationDesignTest/internal/domain"
	repository "applicationDesignTest/internal/repository/inmemory"
	"time"
)

type Order interface {
	CreateOrder(newOrder domain.Order) (domain.Order, error)
}

type RoomAvailability interface {
	CheckAvailability(daysToBook []time.Time) map[time.Time]struct{}
}

type Services struct {
	Order            Order
	RoomAvailability RoomAvailability
}

func NewServices(repos *repository.Repositories) *Services {
	roomAvailabilityService := NewRoomAvailabilityService(repos.RoomsAvailability)
	orderService := NewOrderService(repos.Orders, roomAvailabilityService)

	return &Services{
		Order:            orderService,
		RoomAvailability: roomAvailabilityService,
	}
}
