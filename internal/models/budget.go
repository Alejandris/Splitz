package models

// BudgetRequest representa la solicitud para calcular un presupuesto.
type BudgetRequest struct {
	Salary float64 `json:"salary"`
}

// BudgetResponse representa la respuesta del cálculo de presupuesto.
type BudgetResponse struct {
	MonthlyBudget float64 `json:"monthly_budget"`
	AnnualBudget  float64 `json:"annual_budget"`
}
