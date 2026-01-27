package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/order-api/internal/repository"
	"github.com/gorilla/mux"
)

// OrderHandler maneja las peticiones HTTP relacionadas con órdenes
type OrderHandler struct {
	repo repository.OrderRepositoryPort
}

// NewOrderHandler crea una nueva instancia del handler
func NewOrderHandler(repo repository.OrderRepositoryPort) *OrderHandler {
	return &OrderHandler{repo: repo}
}

// GetOrder maneja GET /order/{id}
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderID := mux.Vars(r)["id"]

	order, err := h.repo.GetByID(orderID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// GetOrders maneja GET /order
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := h.repo.GetAll()
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder maneja POST /order
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Customer string  `json:"customer"`
		Amount   float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	if req.Customer == "" || req.Amount <= 0 {
		http.Error(w, "Missing customer or amount", http.StatusBadRequest)
		return
	}

	order := h.repo.Create(req.Customer, req.Amount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// UpdateOrder maneja PUT /order/{id}
func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]

	var req struct {
		Customer string  `json:"customer"`
		Amount   float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	order, err := h.repo.Update(orderID, req.Customer, req.Amount)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// DeleteOrder maneja DELETE /order/{id}
func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]

	if err := h.repo.Delete(orderID); err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}

// RegisterRoutes registra todas las rutas de órdenes
func (h *OrderHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/order/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/order", h.GetOrders).Methods("GET")
	r.HandleFunc("/order", h.CreateOrder).Methods("POST")
	r.HandleFunc("/order/{id}", h.UpdateOrder).Methods("PUT")
	r.HandleFunc("/order/{id}", h.DeleteOrder).Methods("DELETE")
}
