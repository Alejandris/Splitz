package service

import (
	"fmt"
	"math"

	"splitz/internal/models"
)

// CalculateBudgetByMethod calcula la distribución del salario según el método fijo elegido.
func CalculateBudgetByMethod(salary float64, method int) (models.BudgetResponse, error) {
	if salary <= 0 {
		return models.BudgetResponse{}, fmt.Errorf("salary must be greater than 0")
	}

	var result models.BudgetResponse
	result.Salary = salary
	result.Method = method

	switch method {
	case 1:
		result.MethodName = "Modo Enfoque"
		result.Needs = salary * 0.50
		result.Debt = salary * 0.30
		result.Savings = salary * 0.10
		result.Desires = salary * 0.10
		result.ReadyToSpend = result.Needs + result.Desires
	case 2:
		result.MethodName = "Modo Inversionista o Meta Grande"
		result.Needs = salary * 0.40
		result.Savings = salary * 0.40
		result.Investment = result.Savings
		result.Desires = salary * 0.20
		result.ReadyToSpend = result.Needs + result.Desires
	case 3:
		result.MethodName = "Modo Disfrute"
		result.Needs = salary * 0.40
		result.Lifestyle = salary * 0.40
		result.Savings = salary * 0.20
		result.ReadyToSpend = result.Needs + result.Lifestyle
	default:
		return models.BudgetResponse{}, fmt.Errorf("method must be 1, 2, or 3")
	}

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
		Salary:       request.Salary,
		Needs:        request.Salary * (request.Needs / 100),
		Debt:         request.Salary * (request.Debt / 100),
		Savings:      request.Salary * (request.Savings / 100),
		Desires:      request.Salary * (request.Desires / 100),
		Lifestyle:    request.Salary * (request.Lifestyle / 100),
		Investment:   request.Salary * (request.Investment / 100),
		TotalPercent: total,
	}

	result.ReadyToSpend = result.Needs + result.Desires + result.Lifestyle
	return result, nil
}
