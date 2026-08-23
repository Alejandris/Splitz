package handlers

import (
	"encoding/json"
	"net/http"
)

// CalculateBudget valida la entrada del salario sin implementar el cálculo aún.
func CalculateBudget(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Salary float64 `json:"salary"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if request.Salary < 0 {
		http.Error(w, "salary must be greater than or equal to 0", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
