package service

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/log"
	"applicationDesignTest/internal/pkg/utils"
	"applicationDesignTest/internal/repository"
	"fmt"
	"time"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(newOrder domain.Order) (domain.Order, error) {
	daysToBook := utils.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range repository.Repo.Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			repository.Repo.Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		log.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return domain.Order{}, fmt.Errorf("hotel room is not available for selected dates: \n%v\n%v", newOrder, unavailableDays)
	}

	repository.Repo.Orders = append(repository.Repo.Orders, newOrder)
	return newOrder, nil
}
