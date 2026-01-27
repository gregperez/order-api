package repository

import "example.com/order-api/internal/models"

type OrderRepositoryPort interface {
	GetByID(id string) (models.Order, error)
	GetAll() map[string]models.Order
	Create(customer string, amount float64) models.Order
	Update(id string, customer string, amount float64) (models.Order, error)
	Delete(id string) error
}
