package service

import (
	"fmt"
	"math"

	"splitz/internal/models"
	"splitz/internal/money"
)

// CalculateBudgetByMethod calcula la distribución del salario según el método fijo elegido.
func CalculateBudgetByMethod(salary float64, method int) (models.BudgetResponse, error) {
	if salary <= 0 {
		return models.BudgetResponse{}, fmt.Errorf("salary must be greater than 0")
	}

	var result models.BudgetResponse
	result.Salary = money.Round(salary)
	result.Method = method

	switch method {
	case 1:
		result.MethodName = "Modo Enfoque"
		result.Needs = money.Round(salary * 0.50)
		result.Debt = money.Round(salary * 0.30)
		result.Savings = money.Round(salary * 0.10)
		result.Desires = money.Round(salary * 0.10)
		result.ReadyToSpend = result.Needs + result.Desires
	case 2:
		result.MethodName = "Modo Inversionista o Meta Grande"
		result.Needs = money.Round(salary * 0.40)
		result.Savings = money.Round(salary * 0.40)
		result.Investment = result.Savings
		result.Desires = money.Round(salary * 0.20)
		result.ReadyToSpend = result.Needs + result.Desires
	case 3:
		result.MethodName = "Modo Disfrute"
		result.Needs = money.Round(salary * 0.40)
		result.Lifestyle = money.Round(salary * 0.40)
		result.Savings = money.Round(salary * 0.20)
		result.ReadyToSpend = result.Needs + result.Lifestyle
	default:
		return models.BudgetResponse{}, fmt.Errorf("method must be 1, 2, or 3")
	}

	result.ReadyToSpend = money.Round(result.ReadyToSpend)
	return result, nil
}

// CalculateCustomBudget calcula la distribución del salario usando porcentajes personalizados.
func CalculateCustomBudget(request models.CustomBudgetRequest) (models.CustomBudgetResponse, error) {
	if request.Salary <= 0 {
		return models.CustomBudgetResponse{}, fmt.Errorf("salary must be greater than 0")
	}

	percentages := []float64{request.Needs, request.Debt, request.Savings, request.Desires, request.Lifestyle, request.Investment}
	var total float64
	for _, value := range percentages {
		if value < 0 {
			return models.CustomBudgetResponse{}, fmt.Errorf("percentages must be greater than or equal to 0")
		}
		total += value
	}

	if math.Abs(total-100) > 0.0001 {
		return models.CustomBudgetResponse{}, fmt.Errorf("distribution percentages must add up to 100")
	}

	result := models.CustomBudgetResponse{
		Salary:       money.Round(request.Salary),
		Needs:        money.Round(request.Salary * (request.Needs / 100)),
		Debt:         money.Round(request.Salary * (request.Debt / 100)),
		Savings:      money.Round(request.Salary * (request.Savings / 100)),
		Desires:      money.Round(request.Salary * (request.Desires / 100)),
		Lifestyle:    money.Round(request.Salary * (request.Lifestyle / 100)),
		Investment:   money.Round(request.Salary * (request.Investment / 100)),
		TotalPercent: math.Round(total*10000) / 10000,
	}

	result.ReadyToSpend = money.Round(result.Needs + result.Desires + result.Lifestyle)
	return result, nil
}
