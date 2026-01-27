package repository

import (
	"errors"
	"testing"
)

func TestOrderRepository_GetByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		expectErr error
	}{
		{"Existing ID", "1111", nil},
		{"Non-existing ID", "9999", ErrOrderNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewOrderRepository()
			_, err := repo.GetByID(tt.id)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
		})
	}
}

func TestOrderRepository_GetAll(t *testing.T) {
	repo := NewOrderRepository()
	orders := repo.GetAll()
	if len(orders) != 5 {
		t.Errorf("expected 5 orders, got %d", len(orders))
	}
}

func TestOrderRepository_Create(t *testing.T) {
	tests := []struct {
		name     string
		customer string
		amount   float64
	}{
		{"Create Order 1", "Luis", 1500},
		{"Create Order 2", "Sofia", 2500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewOrderRepository()
			order := repo.Create(tt.customer, tt.amount)
			if order.Customer != tt.customer || order.Amount != tt.amount {
				t.Errorf("expected customer %s and amount %f, got customer %s and amount %f",
					tt.customer, tt.amount, order.Customer, order.Amount)
			}
		})
	}
}

func TestOrderRepository_Update(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		customer  string
		amount    float64
		expectErr error
	}{
		{"Update Existing Order", "1111", "Ana Updated", 1200, nil},
		{"Update Non-existing Order", "9999", "Ghost", 0, ErrOrderNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewOrderRepository()
			order, err := repo.Update(tt.id, tt.customer, tt.amount)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
			if err == nil {
				if order.Customer != tt.customer || order.Amount != tt.amount {
					t.Errorf("expected customer %s and amount %f, got customer %s and amount %f",
						tt.customer, tt.amount, order.Customer, order.Amount)
				}
			}
		})
	}
}

func TestOrderRepository_Delete(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		expectErr error
	}{
		{"Delete Existing Order", "1111", nil},
		{"Delete Non-existing Order", "9999", ErrOrderNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewOrderRepository()
			err := repo.Delete(tt.id)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
		})
	}
}
