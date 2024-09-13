package usecase

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/repository"
	"time"
)

type BookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) CreateOrder(order domain.Order) error {
	daysToBook := daysBetween(order.From, order.To)

	for _, day := range daysToBook {
		availability, err := s.repo.GetAvailability(day, order.RoomID, order.HotelID)
		if err != nil {
			return err
		}
		if availability.Quota < 1 {
			return err
		}
		err = s.repo.UpdateAvailability(day, order.RoomID, order.HotelID)
		if err != nil {
			return err
		}
	}

	return s.repo.CreateOrder(order)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
