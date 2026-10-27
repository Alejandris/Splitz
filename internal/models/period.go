package models

import (
	"time"

	"github.com/google/uuid"
)

type Period struct {
	ID           uuid.UUID `json:"id"`
	Year         int       `json:"year"`
	Month        int       `json:"month"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
	CurrencyCode string    `json:"currency_code"`
	IsClosed     bool      `json:"is_closed"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreatePeriodRequest struct {
	Year         int    `json:"year"`
	Month        int    `json:"month"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	CurrencyCode string `json:"currency_code"`
}

type FixedExpense struct {
	ID             uuid.UUID `json:"id"`
	CategoryID     uuid.UUID `json:"category_id"`
	Name           string    `json:"name"`
	ExpectedAmount float64   `json:"expected_amount"`
	IsActive       bool      `json:"is_active"`
}

type HistoricalFixedExpense struct {
	PeriodID       uuid.UUID `json:"period_id"`
	Year           int       `json:"year"`
	Month          int       `json:"month"`
	FixedExpenseID uuid.UUID `json:"fixed_expense_id"`
	CategoryID     uuid.UUID `json:"category_id"`
	Name           string    `json:"name"`
	ExpectedAmount float64   `json:"expected_amount"`
}

type CreateFixedExpenseRequest struct {
	CategoryID     uuid.UUID `json:"category_id"`
	Name           string    `json:"name"`
	ExpectedAmount float64   `json:"expected_amount"`
}

type UpdateFixedExpenseRequest struct {
	CategoryID     uuid.UUID `json:"category_id"`
	Name           string    `json:"name"`
	ExpectedAmount float64   `json:"expected_amount"`
	IsActive       *bool     `json:"is_active"`
}

type FinancialSummary struct {
	PeriodID             uuid.UUID                `json:"period_id"`
	TotalIncome          float64                  `json:"total_income"`
	TotalFixedExpenses   float64                  `json:"total_fixed_expenses"`
	FixedExpensesPercent float64                  `json:"fixed_expenses_percent"`
	ExceedsFiftyPercent  bool                     `json:"exceeds_fifty_percent"`
	HasSavingsGoal       bool                     `json:"has_savings_goal"`
	HasDebts             bool                     `json:"has_debts"`
	AvailableAmount      float64                  `json:"available_amount"`
	HasDeficit           bool                     `json:"has_deficit"`
	MethodName           string                   `json:"method_name"`
	Warnings             []string                 `json:"warnings"`
	Allocations          []DistributionAllocation `json:"allocations"`
}

type DistributionAllocation struct {
	CategoryID        uuid.UUID `json:"category_id"`
	CategoryName      string    `json:"category_name"`
	CategoryType      string    `json:"category_type"`
	ConfiguredPercent float64   `json:"configured_percent"`
	EffectivePercent  float64   `json:"effective_percent"`
	Amount            float64   `json:"amount"`
	CategoryUsed      bool      `json:"category_used"`
}

type CreateFinancialSummaryRequest struct {
	TotalIncome    float64   `json:"total_income"`
	HasSavingsGoal bool      `json:"has_savings_goal"`
	HasDebts       bool      `json:"has_debts"`
	MethodID       uuid.UUID `json:"method_id"`
}

type SummaryPreviewRequest = CreateFinancialSummaryRequest
