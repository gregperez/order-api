package repository

import (
	"errors"
	"time"

	"example.com/order-api/internal/models"
	"github.com/google/uuid"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

// OrderRepository maneja el almacenamiento en memoria de las órdenes (MOCK)
type OrderRepository struct {
	orders map[string]models.Order
}

// NewOrderRepository crea una nueva instancia del repositorio con datos de prueba
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: map[string]models.Order{
			"1111": {ID: "1111", Customer: "Ana", Amount: 1000, CreatedAt: time.Now()},
			"2222": {ID: "2222", Customer: "Juan", Amount: 2000, CreatedAt: time.Now()},
			"3333": {ID: "3333", Customer: "Maria", Amount: 5500, CreatedAt: time.Now()},
			"4444": {ID: "4444", Customer: "Pedro", Amount: 9990, CreatedAt: time.Now()},
			"5555": {ID: "5555", Customer: "Pablo", Amount: 8000, CreatedAt: time.Now()},
		},
	}
}

// GetByID obtiene una orden por su ID
func (r *OrderRepository) GetByID(id string) (models.Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return models.Order{}, ErrOrderNotFound
	}
	return order, nil
}

// GetAll obtiene todas las órdenes
func (r *OrderRepository) GetAll() map[string]models.Order {
	return r.orders
}

// Create crea una nueva orden
func (r *OrderRepository) Create(customer string, amount float64) models.Order {
	order := models.Order{
		ID:        uuid.New().String(),
		Customer:  customer,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	r.orders[order.ID] = order
	return order
}

// Update actualiza una orden existente
func (r *OrderRepository) Update(id string, customer string, amount float64) (models.Order, error) {
	existing, exists := r.orders[id]
	if !exists {
		return models.Order{}, ErrOrderNotFound
	}

	updated := models.Order{
		ID:        id,
		Customer:  customer,
		Amount:    amount,
		CreatedAt: existing.CreatedAt,
	}
	r.orders[id] = updated
	return updated, nil
}

// Delete elimina una orden
func (r *OrderRepository) Delete(id string) error {
	if _, exists := r.orders[id]; !exists {
		return ErrOrderNotFound
	}
	delete(r.orders, id)
	return nil
}
