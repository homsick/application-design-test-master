package repository

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/utils"
)

var Orders = []domain.Order{}

var Availability = []domain.RoomAvailability{
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 1), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 2), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 3), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 4), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 5), Quota: 0},
}
