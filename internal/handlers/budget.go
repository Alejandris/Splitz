package handlers

import (
	"encoding/json"
	"net/http"

	"splitz/internal/models"
	"splitz/internal/service"
)

// CalculateBudget aplica una de las tres metodologías fijas de reparto del saldo.
func CalculateBudget(w http.ResponseWriter, r *http.Request) {
	var request models.BudgetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if request.Salary <= 0 {
		http.Error(w, "salary must be greater than 0", http.StatusBadRequest)
		return
	}

	if request.Method < 1 || request.Method > 3 {
		http.Error(w, "method must be 1, 2, or 3", http.StatusBadRequest)
		return
	}

	result, err := service.CalculateBudgetByMethod(request.Salary, request.Method)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

// CalculateCustomBudget calcula una distribución personalizada según porcentajes definidos por el usuario.
func CalculateCustomBudget(w http.ResponseWriter, r *http.Request) {
	var request models.CustomBudgetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := service.CalculateCustomBudget(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
