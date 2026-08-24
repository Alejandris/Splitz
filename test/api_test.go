package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"splitz/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// setupServer configura el router de prueba equivalente a tu WebApplicationFactory en .NET
func setupServer() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)
	r.Post("/api/v1/salary", handlers.AddSalary)
	r.Post("/api/v1/budget/calculate", handlers.CalculateBudget)
	return r
}

// TestHealthCheckEndpoint prueba el estado de salud de la API (US-03)
func TestHealthCheckEndpoint(t *testing.T) {
	server := setupServer()

	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba status %d, pero se obtuvo %d", http.StatusOK, rr.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error al parsear respuesta JSON: %v", err)
	}

	if response["status"] != "UP" {
		t.Errorf("Se esperaba status 'UP', pero se obtuvo '%s'", response["status"])
	}
}

// TestCalculateBudgetEndpoint_InvalidSalary prueba el manejo de errores (BadRequest)
func TestCalculateBudgetEndpoint_InvalidSalary(t *testing.T) {
	server := setupServer()

	payload := []byte(`{"salary": -500}`)
	req, _ := http.NewRequest("POST", "/api/v1/budget/calculate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Se esperaba código %d para sueldo inválido, pero se obtuvo %d", http.StatusBadRequest, rr.Code)
	}
}

// TestAddSalaryEndpoint_Success prueba el ingreso de salario neto.
func TestAddSalaryEndpoint_Success(t *testing.T) {
	server := setupServer()

	payload := []byte(`{"net_salary": 2500}`)
	req, _ := http.NewRequest("POST", "/api/v1/salary", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("Se esperaba código %d, pero se obtuvo %d", http.StatusCreated, rr.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error al parsear respuesta JSON: %v", err)
	}

	if response["net_salary"] != float64(2500) {
		t.Errorf("Se esperaba net_salary 2500, pero se obtuvo %v", response["net_salary"])
	}

	if response["ready_for_processing"] != true {
		t.Errorf("Se esperaba ready_for_processing=true, pero se obtuvo %v", response["ready_for_processing"])
	}
}

// TestAddSalaryEndpoint_InvalidNetSalary prueba manejo de errores del ingreso de salario neto.
func TestAddSalaryEndpoint_InvalidNetSalary(t *testing.T) {
	server := setupServer()

	payload := []byte(`{"net_salary": -50}`)
	req, _ := http.NewRequest("POST", "/api/v1/salary", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Se esperaba código %d para salario neto inválido, pero se obtuvo %d", http.StatusBadRequest, rr.Code)
	}
}
