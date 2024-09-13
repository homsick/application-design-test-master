package domain

import "time"

type RoomAvailability struct {
	HotelID string
	RoomID  string
	Date    time.Time
	Quota   int
}
