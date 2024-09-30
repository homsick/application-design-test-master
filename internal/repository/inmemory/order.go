package repository

import "applicationDesignTest/internal/domain"

type InMemoryOrdersRepository struct {
	Orders []domain.Order
}

func NewInMemoryOrdersRepository() *InMemoryOrdersRepository {
	return &InMemoryOrdersRepository{
		Orders: []domain.Order{},
	}
}

func (r *InMemoryOrdersRepository) Add(newOrder domain.Order) {
	r.Orders = append(r.Orders, newOrder)
}

func (r *InMemoryOrdersRepository) GetAll() []domain.Order {
	return r.Orders
}

func (r *InMemoryOrdersRepository) GetByID(orderID string) (domain.Order, error) {
	return domain.Order{}, nil // TODO: Implement order retrieval logic
}

func (r *InMemoryOrdersRepository) Delete(orderID string) error {
	return nil // TODO: Implement order deletion logic
}

func (r *InMemoryOrdersRepository) Update(order domain.Order) error {
	return nil // TODO: Implement order update logic
}
