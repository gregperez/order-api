package main

import (
	"log"
	"net/http"

	"example.com/order-api/examples"
	"example.com/order-api/internal/handlers"
	"example.com/order-api/internal/middleware"
	"example.com/order-api/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	// Ejecutar demos de ejercicios (opcional)
	examples.RunExercisesDemo()

	// Inicializar repositorio y handlers
	orderRepo := repository.NewOrderRepository()
	orderHandler := handlers.NewOrderHandler(orderRepo)

	// Configurar router
	r := mux.NewRouter()
	orderHandler.RegisterRoutes(r)

	// Aplicar middleware
	r.Use(middleware.Logging)

	// Iniciar servidor
	log.Println("🚀 Server listening on :8080")
	log.Println("📚 API Endpoints:")
	log.Println("   GET    /order      - List all orders")
	log.Println("   GET    /order/{id} - Get order by ID")
	log.Println("   POST   /order      - Create new order")
	log.Println("   PUT    /order/{id} - Update order")
	log.Println("   DELETE /order/{id} - Delete order")

	log.Fatal(http.ListenAndServe(":8080", r))
}
