package service

import (
	"applicationDesignTest/internal/domain"
	"time"
)

type RoomAvailabilitiesRepository interface {
	Add(roomAvailability []domain.RoomAvailability)
	Update(availability domain.RoomAvailability) error
	Delete(hotelID, roomID string, date time.Time) error
	GetByDate(hotelID, roomID string, date time.Time) (domain.RoomAvailability, error)
	CheckAvailability(daysToBook []time.Time) map[time.Time]struct{}
}

type RoomAvailabilityService struct {
	repo RoomAvailabilitiesRepository
}

func NewRoomAvailabilityService(repo RoomAvailabilitiesRepository) *RoomAvailabilityService {
	return &RoomAvailabilityService{
		repo: repo,
	}
}

func (s *RoomAvailabilityService) CheckAvailability(daysToBook []time.Time) map[time.Time]struct{} {
	return s.repo.CheckAvailability(daysToBook)
}
