package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/order-api/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Order struct {
	ID        string    `json:"id"`
	Customer  string    `json:"customer"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

var ordersMap = map[string]Order{
	"1111": {ID: "1111", Customer: "Ana", Amount: 1000, CreatedAt: time.Now()},
	"2222": {ID: "2222", Customer: "Juan", Amount: 2000, CreatedAt: time.Now()},
	"3333": {ID: "3333", Customer: "Maria", Amount: 5500, CreatedAt: time.Now()},
	"4444": {ID: "4444", Customer: "Pedro", Amount: 9990, CreatedAt: time.Now()},
	"5555": {ID: "5555", Customer: "Pablo", Amount: 8000, CreatedAt: time.Now()},
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderID := mux.Vars(r)["id"]

	order, exists := ordersMap[orderID]
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ordersMap)
}

func postOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}
	if newOrder.Customer == "" || newOrder.Amount <= 0 {
		http.Error(w, "Missing customer or amount", http.StatusBadRequest)
		return
	}

	newOrder.ID = uuid.New().String()
	newOrder.CreatedAt = time.Now()
	ordersMap[newOrder.ID] = newOrder

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func putOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}
	orderID := mux.Vars(r)["id"]

	existing, exists := ordersMap[orderID]
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	updatedOrder.ID = orderID
	updatedOrder.CreatedAt = existing.CreatedAt
	ordersMap[orderID] = updatedOrder

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedOrder)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]
	if _, exists := ordersMap[orderID]; !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	delete(ordersMap, orderID)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	utils.Demo()
	log.Println("Demo completed")

	testCases := []string{
		"aja",
		"aba",
		"setset",
		"abcba",
		"abcta",
		"arepera",
		"c",
		"", // también probamos string vacío
	}

	for _, word := range testCases {
		fmt.Printf("¿Es palíndromo '%s'? → %t\n", word, utils.IsPalindrome(word))
	}

	r := mux.NewRouter()
	r.HandleFunc("/order/{id}", getOrder).Methods("GET")
	r.HandleFunc("/order", getOrders).Methods("GET")
	r.HandleFunc("/order", postOrder).Methods("POST")
	r.HandleFunc("/order/{id}", putOrder).Methods("PUT")
	r.HandleFunc("/order/{id}", deleteOrder).Methods("DELETE")
	r.Use(loggingMiddleware)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
