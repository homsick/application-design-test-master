package app

import (
	delivery "applicationDesignTest/internal/delivery/http"
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/pkg/utils"
	repository "applicationDesignTest/internal/repository/inmemory"
	"applicationDesignTest/internal/server"
	"applicationDesignTest/internal/service"
)

func Run() {

	repos := repository.NewRepositories()
	initialAvailability := []domain.RoomAvailability{
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 1), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 2), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 3), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 4), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 5), Quota: 0},
	}
	repos.RoomsAvailability.Add(initialAvailability)
	services := service.NewServices(repos)
	handlers := delivery.NewHandler(services)
	server.StartServer(handlers)
}
