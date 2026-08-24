package handlers

import (
	"encoding/json"
	"net/http"

	"splitz/internal/models"
)

// AddSalary registra salario neto y devuelve salida lista para el siguiente endpoint.
func AddSalary(w http.ResponseWriter, r *http.Request) {
	var request models.AddSalaryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if request.NetSalary <= 0 {
		http.Error(w, "net_salary must be greater than 0", http.StatusBadRequest)
		return
	}

	response := models.AddSalaryResponse{
		NetSalary:          request.NetSalary,
		ReadyForProcessing: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
