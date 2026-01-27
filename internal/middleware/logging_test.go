package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		requestURI    string
		expectedLog   string
		statusCode    int
		handlerCalled bool
	}{
		{
			name:          "GET request logs correctly",
			method:        "GET",
			requestURI:    "/api/orders",
			expectedLog:   "GET /api/orders",
			statusCode:    http.StatusOK,
			handlerCalled: true,
		},
		{
			name:          "POST request logs correctly",
			method:        "POST",
			requestURI:    "/api/orders",
			expectedLog:   "POST /api/orders",
			statusCode:    http.StatusCreated,
			handlerCalled: true,
		},
		{
			name:          "DELETE request logs correctly",
			method:        "DELETE",
			requestURI:    "/api/orders/123",
			expectedLog:   "DELETE /api/orders/123",
			statusCode:    http.StatusNoContent,
			handlerCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capturar el output del log
			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer log.SetOutput(nil)

			// Flag para verificar si el handler fue llamado
			handlerCalled := false

			// Handler mock que será envuelto por el middleware
			mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlerCalled = true
				w.WriteHeader(tt.statusCode)
			})

			// Aplicar el middleware
			handler := Logging(mockHandler)

			// Crear una petición de prueba
			req := httptest.NewRequest(tt.method, tt.requestURI, nil)
			rr := httptest.NewRecorder()

			// Ejecutar el handler
			handler.ServeHTTP(rr, req)

			// Verificar que el handler interno fue llamado
			if handlerCalled != tt.handlerCalled {
				t.Errorf("expected handler called = %v, got %v", tt.handlerCalled, handlerCalled)
			}

			// Verificar que el log contiene el mensaje esperado
			logOutput := buf.String()
			if !strings.Contains(logOutput, tt.expectedLog) {
				t.Errorf("expected log to contain '%s', got '%s'", tt.expectedLog, logOutput)
			}

			// Verificar el status code
			if status := rr.Code; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.statusCode)
			}
		})
	}
}
