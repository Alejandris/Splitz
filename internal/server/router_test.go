package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouterDoesNotExposeFinancialRoutesWithoutDependencies(t *testing.T) {
	router := NewRouter()

	health := httptest.NewRecorder()
	router.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/health", nil))
	if health.Code != http.StatusOK {
		t.Fatalf("health debe permanecer público: se obtuvo %d", health.Code)
	}

	financial := httptest.NewRecorder()
	router.ServeHTTP(financial, httptest.NewRequest(http.MethodPost, "/api/v1/budget/calculate", nil))
	if financial.Code != http.StatusNotFound {
		t.Fatalf("la ruta financiera no configurada debe estar ausente: se obtuvo %d", financial.Code)
	}
}
