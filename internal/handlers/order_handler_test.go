package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"example.com/order-api/internal/models"
	"example.com/order-api/internal/repository"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) GetByID(id string) (models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(models.Order), args.Error(1)
}

func (m *MockOrderRepository) GetAll() map[string]models.Order {
	args := m.Called()
	return args.Get(0).(map[string]models.Order)
}

func (m *MockOrderRepository) Create(customer string, amount float64) models.Order {
	args := m.Called(customer, amount)
	return args.Get(0).(models.Order)
}

func (m *MockOrderRepository) Update(id string, customer string, amount float64) (models.Order, error) {
	args := m.Called(id, customer, amount)
	return args.Get(0).(models.Order), args.Error(1)
}

func (m *MockOrderRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestOrderHandler_GetOrder(t *testing.T) {
	location, _ := time.LoadLocation("UTC")
	createdAt := time.Date(2026, 01, 01, 00, 00, 00, 00, location)

	tests := []struct {
		name           string
		orderID        string
		mockSetup      func(m *MockOrderRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Order exists",
			orderID: "1234",
			mockSetup: func(m *MockOrderRepository) {
				m.On("GetByID", "1234").
					Return(
						models.Order{
							ID:        "1234",
							Customer:  "Alice",
							Amount:    1000.0,
							CreatedAt: createdAt},
						nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"id":"1234","customer":"Alice","amount":1000,"created_at":"2026-01-01T00:00:00Z"}` + "\n",
		},
		{
			name:    "Order does not exist",
			orderID: "2",
			mockSetup: func(m *MockOrderRepository) {
				m.On("GetByID", "2").Return(models.Order{}, repository.ErrOrderNotFound)
			},
			expectedStatus: 404,
			expectedBody:   "Order not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el mock
			mockRepo := new(MockOrderRepository)
			tt.mockSetup(mockRepo)

			// Crear el handler con el mock
			handler := NewOrderHandler(mockRepo)

			// Crear la petición HTTP
			req := httptest.NewRequest("GET", "/order/"+tt.orderID, nil)
			// Simular las variables de ruta que mux.Vars extraería
			req = mux.SetURLVars(req, map[string]string{"id": tt.orderID})
			rr := httptest.NewRecorder()

			// Ejecutar la petición
			handler.GetOrder(rr, req)

			// Verificar la respuesta
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			if rr.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestOrderHandler_GetOrders(t *testing.T) {
	tests := []struct {
		name                 string
		mockSetup            func(m *MockOrderRepository)
		expectedStatus       int
		expectedQuantityBody int
	}{
		{
			name: "Get all orders",
			mockSetup: func(m *MockOrderRepository) {
				m.On("GetAll").Return(map[string]models.Order{
					"1111": {ID: "1111", Customer: "Ana", Amount: 1000},
					"2222": {ID: "2222", Customer: "Juan", Amount: 2000},
				})
			},
			expectedStatus:       200,
			expectedQuantityBody: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el mock
			mockRepo := new(MockOrderRepository)
			tt.mockSetup(mockRepo)

			// Crear el handler con el mock
			handler := NewOrderHandler(mockRepo)

			// Crear la petición HTTP
			req := httptest.NewRequest("GET", "/order", nil)
			rr := httptest.NewRecorder()

			// Ejecutar la petición
			handler.GetOrders(rr, req)

			// Verificar la respuesta
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			var orders map[string]models.Order
			err := json.Unmarshal(rr.Body.Bytes(), &orders)
			if err != nil {
				t.Errorf("error unmarshaling response: %v", err)
			}
			if len(orders) != tt.expectedQuantityBody {
				t.Errorf("expected %d orders, got %d", tt.expectedQuantityBody, len(orders))
			}
		})
	}
}

func TestOrderHandler_CreateOrder(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(m *MockOrderRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Create order successfully",
			requestBody: `{"customer":"Bob","amount":1500}`,
			mockSetup: func(m *MockOrderRepository) {
				m.On("Create", "Bob", 1500.0).
					Return(models.Order{
						ID:        "new-id",
						Customer:  "Bob",
						Amount:    1500.0,
						CreatedAt: time.Now(),
					})
			},
			expectedStatus: 201,
			expectedBody:   `"customer":"Bob","amount":1500`, // Partial match
		},
		{
			name:           "Create order with invalid data",
			requestBody:    `{"customer":"","amount":0}`,
			mockSetup:      func(m *MockOrderRepository) {},
			expectedStatus: 400,
			expectedBody:   "Missing customer or amount\n",
		},
		{
			name:           "Create order with invalid JSON",
			requestBody:    `invalid-json`,
			mockSetup:      func(m *MockOrderRepository) {},
			expectedStatus: 400,
			expectedBody:   "Invalid order data\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el mock
			mockRepo := new(MockOrderRepository)
			tt.mockSetup(mockRepo)

			// Crear el handler con el mock
			handler := NewOrderHandler(mockRepo)

			// Crear la petición HTTP
			req := httptest.NewRequest("POST", "/order", nil)
			if bodyStr, ok := tt.requestBody.(string); ok {
				req = httptest.NewRequest("POST", "/order", strings.NewReader(bodyStr))
			}
			rr := httptest.NewRecorder()

			// Ejecutar la petición
			handler.CreateOrder(rr, req)

			// Verificar la respuesta
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			if tt.expectedBody != "" && !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestOrderHandler_UpdateOrder(t *testing.T) {
	tests := []struct {
		name           string
		orderID        string
		requestBody    interface{}
		mockSetup      func(m *MockOrderRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Update order successfully",
			orderID: "1234",
			requestBody: `{
				"customer":"Charlie",
				"amount":2000
			}`,
			mockSetup: func(m *MockOrderRepository) {
				m.On("Update", "1234", "Charlie", 2000.0).
					Return(models.Order{
						ID:        "1234",
						Customer:  "Charlie",
						Amount:    2000.0,
						CreatedAt: time.Now(),
					}, nil)
			},
			expectedStatus: 200,
			expectedBody:   `"customer":"Charlie","amount":2000`, // Partial match
		},
		{
			name:           "Update order with invalid JSON",
			orderID:        "1234",
			requestBody:    `invalid-json`,
			mockSetup:      func(m *MockOrderRepository) {},
			expectedStatus: 400,
			expectedBody:   "Invalid order data\n",
		},
		{
			name:    "Update non-existing order",
			orderID: "9999",
			requestBody: `{
				"customer":"NonExistent",
				"amount":500
			}`,
			mockSetup: func(m *MockOrderRepository) {
				m.On("Update", "9999", "NonExistent", 500.0).
					Return(models.Order{}, repository.ErrOrderNotFound)
			},
			expectedStatus: 404,
			expectedBody:   "Order not found\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el mock
			mockRepo := new(MockOrderRepository)
			tt.mockSetup(mockRepo)

			// Crear el handler con el mock
			handler := NewOrderHandler(mockRepo)

			// Crear la petición HTTP
			req := httptest.NewRequest("PUT", "/order/"+tt.orderID, nil)
			if bodyStr, ok := tt.requestBody.(string); ok {
				req = httptest.NewRequest("PUT", "/order/"+tt.orderID, strings.NewReader(bodyStr))
			}
			// Simular las variables de ruta que mux.Vars extraería
			req = mux.SetURLVars(req, map[string]string{"id": tt.orderID})
			rr := httptest.NewRecorder()

			// Ejecutar la petición
			handler.UpdateOrder(rr, req)

			// Verificar la respuesta
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			if tt.expectedBody != "" && !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestOrderHandler_DeleteOrder(t *testing.T) {
	tests := []struct {
		name           string
		orderID        string
		mockSetup      func(m *MockOrderRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Delete order successfully",
			orderID: "1234",
			mockSetup: func(m *MockOrderRepository) {
				m.On("Delete", "1234").Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"Order deleted successfully"}` + "\n",
		},
		{
			name:    "Delete non-existing order",
			orderID: "9999",
			mockSetup: func(m *MockOrderRepository) {
				m.On("Delete", "9999").Return(repository.ErrOrderNotFound)
			},
			expectedStatus: 404,
			expectedBody:   "Order not found\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el mock
			mockRepo := new(MockOrderRepository)
			tt.mockSetup(mockRepo)

			// Crear el handler con el mock
			handler := NewOrderHandler(mockRepo)

			// Crear la petición HTTP
			req := httptest.NewRequest("DELETE", "/order/"+tt.orderID, nil)
			// Simular las variables de ruta que mux.Vars extraería
			req = mux.SetURLVars(req, map[string]string{"id": tt.orderID})
			rr := httptest.NewRecorder()

			// Ejecutar la petición
			handler.DeleteOrder(rr, req)

			// Verificar la respuesta
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			if rr.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestOrderHandler_RegisterRoutes(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	handler := NewOrderHandler(mockRepo)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	expectedRoutes := []string{
		"/order/{id}",
		"/order",
	}
	for _, route := range expectedRoutes {
		req := httptest.NewRequest("GET", route, nil)
		match := &mux.RouteMatch{}
		if !router.Match(req, match) {
			t.Errorf("expected route %s to be registered", route)
		}
	}
}
