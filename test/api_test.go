package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"splitz/internal/models"
	"splitz/internal/server"
	"splitz/internal/service"
	"testing"
)

// setupServer configura el router de prueba equivalente a tu WebApplicationFactory en .NET
func setupServer() http.Handler {
	return server.NewRouter()
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

	payload := []byte(`{"salary": -500, "method": 1}`)
	req, _ := http.NewRequest("POST", "/api/v1/budget/calculate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Se esperaba código %d para sueldo inválido, pero se obtuvo %d", http.StatusBadRequest, rr.Code)
	}
}

// TestCalculateBudgetByMethod_Enfoque valida la lógica del método 1.
func TestCalculateBudgetByMethod_Enfoque(t *testing.T) {
	result, err := service.CalculateBudgetByMethod(3000, 1)
	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if result.Needs != 1500 {
		t.Fatalf("Se esperaba needs=1500, pero se obtuvo %v", result.Needs)
	}
	if result.Debt != 900 {
		t.Fatalf("Se esperaba debt=900, pero se obtuvo %v", result.Debt)
	}
	if result.Savings != 300 {
		t.Fatalf("Se esperaba savings=300, pero se obtuvo %v", result.Savings)
	}
	if result.Desires != 300 {
		t.Fatalf("Se esperaba desires=300, pero se obtuvo %v", result.Desires)
	}
}

// TestCalculateBudgetByMethod_AllProfiles valida las tres metodologías fijas.
func TestCalculateBudgetByMethod_AllProfiles(t *testing.T) {
	tests := []struct {
		name      string
		salary    float64
		method    int
		needs     float64
		debt      float64
		savings   float64
		desires   float64
		lifestyle float64
	}{
		{name: "Modo enfoque", salary: 3000, method: 1, needs: 1500, debt: 900, savings: 300, desires: 300},
		{name: "Modo inversionista", salary: 3000, method: 2, needs: 1200, savings: 1200, desires: 600},
		{name: "Modo disfrute", salary: 4000, method: 3, needs: 1600, lifestyle: 1600, savings: 800},
	}

	for _, tc := range tests {
		result, err := service.CalculateBudgetByMethod(tc.salary, tc.method)
		if err != nil {
			t.Fatalf("%s: no se esperaba error: %v", tc.name, err)
		}

		if result.Needs != tc.needs {
			t.Fatalf("%s: esperaba needs=%v, pero se obtuvo %v", tc.name, tc.needs, result.Needs)
		}
		if result.Debt != tc.debt {
			t.Fatalf("%s: esperaba debt=%v, pero se obtuvo %v", tc.name, tc.debt, result.Debt)
		}
		if result.Savings != tc.savings {
			t.Fatalf("%s: esperaba savings=%v, pero se obtuvo %v", tc.name, tc.savings, result.Savings)
		}
		if result.Desires != tc.desires {
			t.Fatalf("%s: esperaba desires=%v, pero se obtuvo %v", tc.name, tc.desires, result.Desires)
		}
		if result.Lifestyle != tc.lifestyle {
			t.Fatalf("%s: esperaba lifestyle=%v, pero se obtuvo %v", tc.name, tc.lifestyle, result.Lifestyle)
		}
	}
}

// TestCalculateBudgetEndpoint_Success valida el cálculo según método seleccionado.
func TestCalculateBudgetEndpoint_Success(t *testing.T) {
	server := setupServer()

	payload := []byte(`{"salary": 3000, "method": 1}`)
	req, _ := http.NewRequest("POST", "/api/v1/budget/calculate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba código %d, pero se obtuvo %d", http.StatusOK, rr.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error al parsear respuesta JSON: %v", err)
	}

	if response["method"] != float64(1) {
		t.Errorf("Se esperaba method=1, pero se obtuvo %v", response["method"])
	}
	if response["needs"] != float64(1500) {
		t.Errorf("Se esperaba needs=1500, pero se obtuvo %v", response["needs"])
	}
}

// TestCalculateBudgetEndpoint_InvalidMethod valida opción no permitida.
func TestCalculateBudgetEndpoint_InvalidMethod(t *testing.T) {
	server := setupServer()

	payload := []byte(`{"salary": 3000, "method": 9}`)
	req, _ := http.NewRequest("POST", "/api/v1/budget/calculate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Se esperaba código %d para método inválido, pero se obtuvo %d", http.StatusBadRequest, rr.Code)
	}
}

// TestCalculateCustomBudgetByService_ValidPercentages valida la lógica personalizada.
func TestCalculateCustomBudgetByService_ValidPercentages(t *testing.T) {
	request := models.CustomBudgetRequest{Salary: 4000, Needs: 40, Debt: 20, Savings: 20, Desires: 20}
	result, err := service.CalculateCustomBudget(request)
	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if result.TotalPercent != 100 {
		t.Fatalf("Se esperaba total 100, pero se obtuvo %v", result.TotalPercent)
	}
	if result.Needs != 1600 {
		t.Fatalf("Se esperaba needs=1600, pero se obtuvo %v", result.Needs)
	}
	if result.Debt != 800 {
		t.Fatalf("Se esperaba debt=800, pero se obtuvo %v", result.Debt)
	}
	if result.Savings != 800 {
		t.Fatalf("Se esperaba savings=800, pero se obtuvo %v", result.Savings)
	}
	if result.Desires != 800 {
		t.Fatalf("Se esperaba desires=800, pero se obtuvo %v", result.Desires)
	}
}

// TestCalculateCustomBudgetByService_InvalidPercentages valida suma no permitida.
func TestCalculateCustomBudgetByService_InvalidPercentages(t *testing.T) {
	request := models.CustomBudgetRequest{Salary: 4000, Needs: 40, Debt: 20, Savings: 20, Desires: 10}
	_, err := service.CalculateCustomBudget(request)
	if err == nil {
		t.Fatal("Se esperaba error cuando los porcentajes no suman 100")
	}
}

// TestCalculateCustomBudgetEndpoint_Success valida el endpoint personalizado.
func TestCalculateCustomBudgetEndpoint_Success(t *testing.T) {
	server := setupServer()

	payload := []byte(`{"salary":4000,"needs":40,"debt":20,"savings":20,"desires":20}`)
	req, _ := http.NewRequest("POST", "/api/v1/budget/custom", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba código %d, pero se obtuvo %d", http.StatusOK, rr.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error al parsear respuesta JSON: %v", err)
	}

	if response["salary"] != float64(4000) {
		t.Errorf("Se esperaba salary=4000, pero se obtuvo %v", response["salary"])
	}
	if response["total_percent"] != float64(100) {
		t.Errorf("Se esperaba total_percent=100, pero se obtuvo %v", response["total_percent"])
	}
	if response["needs"] != float64(1600) {
		t.Errorf("Se esperaba needs=1600, pero se obtuvo %v", response["needs"])
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
	if response["salary"] != float64(2500) {
		t.Errorf("Se esperaba salary 2500, pero se obtuvo %v", response["salary"])
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
