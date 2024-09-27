package service

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/pkg/utils"
	"fmt"
	"time"
)

type Repository interface {
	AddOrder(newOrder domain.Order)
	CheckAvailability(daysToBook []time.Time) map[time.Time]struct{}
}

type OrderService struct {
	repo Repository
}

func NewOrderService(repo Repository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(newOrder domain.Order) (domain.Order, error) {
	daysToBook := utils.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := s.repo.CheckAvailability(daysToBook)

	if len(unavailableDays) != 0 {
		log.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return domain.Order{}, fmt.Errorf("hotel room is not available for selected dates: \n%v\n%v", newOrder, unavailableDays)
	}

	s.repo.AddOrder(newOrder)
	return newOrder, nil
}
