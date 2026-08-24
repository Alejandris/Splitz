package models

// BudgetRequest representa la solicitud para calcular el presupuesto según un método fijo.
type BudgetRequest struct {
	Salary float64 `json:"salary"`
	Method int     `json:"method"`
}

// BudgetResponse representa el resultado del cálculo del presupuesto.
type BudgetResponse struct {
	Salary       float64 `json:"salary"`
	Method       int     `json:"method"`
	MethodName   string  `json:"method_name"`
	Needs        float64 `json:"needs"`
	Debt         float64 `json:"debt"`
	Savings      float64 `json:"savings"`
	Desires      float64 `json:"desires"`
	Lifestyle    float64 `json:"lifestyle"`
	Investment   float64 `json:"investment"`
	ReadyToSpend float64 `json:"ready_to_spend"`
}

// CustomBudgetRequest representa una distribución porcentual personalizada del salario.
type CustomBudgetRequest struct {
	Salary     float64 `json:"salary"`
	Needs      float64 `json:"needs"`
	Debt       float64 `json:"debt"`
	Savings    float64 `json:"savings"`
	Desires    float64 `json:"desires"`
	Lifestyle  float64 `json:"lifestyle"`
	Investment float64 `json:"investment"`
}

// CustomBudgetResponse representa el resultado de un cálculo personalizado.
type CustomBudgetResponse struct {
	Salary       float64 `json:"salary"`
	Needs        float64 `json:"needs"`
	Debt         float64 `json:"debt"`
	Savings      float64 `json:"savings"`
	Desires      float64 `json:"desires"`
	Lifestyle    float64 `json:"lifestyle"`
	Investment   float64 `json:"investment"`
	TotalPercent float64 `json:"total_percent"`
	ReadyToSpend float64 `json:"ready_to_spend"`
}
