package service

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/pkg/utils"
	"fmt"
)

type OrdersRepository interface {
	Add(newOrder domain.Order)
	GetAll() []domain.Order
	GetByID(orderID string) (domain.Order, error)
	Update(order domain.Order) error
	Delete(orderID string) error
}

type OrderService struct {
	repo OrdersRepository

	roomAvailabilityService RoomAvailability
}

func NewOrderService(repo OrdersRepository, roomAvailabilityService RoomAvailability) *OrderService {
	return &OrderService{
		repo: repo,

		roomAvailabilityService: roomAvailabilityService,
	}
}

func (s *OrderService) CreateOrder(newOrder domain.Order) (domain.Order, error) {
	daysToBook := utils.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := s.roomAvailabilityService.CheckAvailability(daysToBook)

	if len(unavailableDays) != 0 {
		log.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return domain.Order{}, fmt.Errorf("hotel room is not available for selected dates: \n%v\n%v", newOrder, unavailableDays)
	}

	s.repo.Add(newOrder)
	return newOrder, nil
}
