package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"splitz/internal/models"
	"splitz/internal/money"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrPeriodNotFound = errors.New("period not found")
var ErrPeriodClosed = errors.New("period is closed")

type PeriodService struct {
	db *pgxpool.Pool
}

func NewPeriodService(db *pgxpool.Pool) *PeriodService {
	return &PeriodService{db: db}
}

func validCurrency(code string) bool {
	return code == "MXN" || code == "COP" || code == "USD" || code == "EUR"
}

func periodDates(request models.CreatePeriodRequest) (string, string, error) {
	if request.StartDate == "" && request.EndDate == "" {
		start := time.Date(request.Year, time.Month(request.Month), 1, 0, 0, 0, 0, time.UTC)
		return start.Format("2006-01-02"), start.AddDate(0, 1, -1).Format("2006-01-02"), nil
	}
	start, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		return "", "", errors.New("start_date must use YYYY-MM-DD")
	}
	end, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		return "", "", errors.New("end_date must use YYYY-MM-DD")
	}
	if end.Before(start) {
		return "", "", errors.New("end_date must be on or after start_date")
	}
	return start.Format("2006-01-02"), end.Format("2006-01-02"), nil
}

func (service *PeriodService) CreateOrGetPeriod(ctx context.Context, userID uuid.UUID, request models.CreatePeriodRequest) (models.Period, error) {
	if request.Year < 2000 || request.Year > 9999 {
		return models.Period{}, errors.New("year must be between 2000 and 9999")
	}
	if request.Month < 1 || request.Month > 12 {
		return models.Period{}, errors.New("month must be between 1 and 12")
	}
	startDate, endDate, err := periodDates(request)
	if err != nil {
		return models.Period{}, err
	}
	if request.CurrencyCode == "" {
		request.CurrencyCode = "MXN"
	}
	if !validCurrency(request.CurrencyCode) {
		return models.Period{}, errors.New("currency_code must be MXN, COP, USD, or EUR")
	}

	tx, err := service.db.Begin(ctx)
	if err != nil {
		return models.Period{}, fmt.Errorf("begin period: %w", err)
	}
	defer tx.Rollback(ctx)

	var period models.Period
	err = tx.QueryRow(ctx, `
		INSERT INTO periods (user_id, year, month, start_date, end_date, currency_code)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (user_id, year, month) DO UPDATE SET user_id = EXCLUDED.user_id
		RETURNING id, year, month, start_date, end_date, currency_code, is_closed, created_at`,
		userID, request.Year, request.Month, startDate, endDate, request.CurrencyCode,
	).Scan(&period.ID, &period.Year, &period.Month, &period.StartDate, &period.EndDate, &period.CurrencyCode, &period.IsClosed, &period.CreatedAt)
	if err != nil {
		return models.Period{}, fmt.Errorf("create or get period: %w", err)
	}
	if !period.IsClosed {
		_, err = tx.Exec(ctx, `
			INSERT INTO period_fixed_expenses
				(period_id, fixed_expense_id, expected_amount, paid_exact_amount, actual_paid_amount, is_paid)
			SELECT $1, id, expected_amount, FALSE, expected_amount, FALSE
			FROM fixed_expenses
			WHERE user_id = $2 AND is_active = TRUE
			ON CONFLICT (period_id, fixed_expense_id) DO NOTHING`, period.ID, userID)
		if err != nil {
			return models.Period{}, fmt.Errorf("snapshot fixed expenses: %w", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return models.Period{}, fmt.Errorf("commit period: %w", err)
	}
	return period, nil
}

func (service *PeriodService) GetPeriod(ctx context.Context, userID, periodID uuid.UUID) (models.Period, error) {
	var period models.Period
	err := service.db.QueryRow(ctx, `
		SELECT id, year, month, start_date, end_date, currency_code, is_closed, created_at
		FROM periods
		WHERE id = $1 AND user_id = $2`,
		periodID, userID,
	).Scan(&period.ID, &period.Year, &period.Month, &period.StartDate, &period.EndDate, &period.CurrencyCode, &period.IsClosed, &period.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return models.Period{}, ErrPeriodNotFound
	}
	if err != nil {
		return models.Period{}, fmt.Errorf("get period: %w", err)
	}
	return period, nil
}

func (service *PeriodService) CreateFixedExpense(ctx context.Context, userID uuid.UUID, request models.CreateFixedExpenseRequest) (models.FixedExpense, error) {
	name := strings.TrimSpace(request.Name)
	if name == "" {
		return models.FixedExpense{}, errors.New("name is required")
	}
	if request.ExpectedAmount < 0 {
		return models.FixedExpense{}, errors.New("expected_amount must be greater than or equal to 0")
	}
	if request.CategoryID == uuid.Nil {
		return models.FixedExpense{}, errors.New("category_id is required")
	}

	var expense models.FixedExpense
	err := service.db.QueryRow(ctx, `
		INSERT INTO fixed_expenses (user_id, category_id, name, expected_amount)
		SELECT $1, id, $3, $4
		FROM categories
		WHERE id = $2 AND is_active = TRUE
			AND (created_by_user_id IS NULL OR created_by_user_id = $1)
		RETURNING id, category_id, name, expected_amount, is_active`,
		userID, request.CategoryID, name, money.Round(request.ExpectedAmount),
	).Scan(&expense.ID, &expense.CategoryID, &expense.Name, &expense.ExpectedAmount, &expense.IsActive)
	if err != nil {
		return models.FixedExpense{}, fmt.Errorf("create fixed expense: %w", err)
	}
	expense.ExpectedAmount = money.Round(expense.ExpectedAmount)
	return expense, nil
}

func (service *PeriodService) ListFixedExpenses(ctx context.Context, userID uuid.UUID) ([]models.FixedExpense, error) {
	rows, err := service.db.Query(ctx, `
		SELECT id, category_id, name, expected_amount, is_active
		FROM fixed_expenses
		WHERE user_id = $1
		ORDER BY name, id`, userID)
	if err != nil {
		return nil, fmt.Errorf("list fixed expenses: %w", err)
	}
	defer rows.Close()

	expenses := make([]models.FixedExpense, 0)
	for rows.Next() {
		var expense models.FixedExpense
		if err := rows.Scan(&expense.ID, &expense.CategoryID, &expense.Name, &expense.ExpectedAmount, &expense.IsActive); err != nil {
			return nil, fmt.Errorf("scan fixed expense: %w", err)
		}

		expense.ExpectedAmount = money.Round(expense.ExpectedAmount)
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate fixed expenses: %w", err)
	}
	return expenses, nil
}

func (service *PeriodService) ListHistoricalFixedExpenses(ctx context.Context, userID uuid.UUID) ([]models.HistoricalFixedExpense, error) {
	rows, err := service.db.Query(ctx, `
		SELECT p.id, p.year, p.month, f.id, f.category_id, f.name, pfe.expected_amount
		FROM periods p
		JOIN period_fixed_expenses pfe ON pfe.period_id = p.id
		JOIN fixed_expenses f ON f.id = pfe.fixed_expense_id
		WHERE p.user_id = $1 AND f.user_id = $1
		ORDER BY p.year DESC, p.month DESC, f.name`, userID)
	if err != nil {
		return nil, fmt.Errorf("list historical fixed expenses: %w", err)
	}
	defer rows.Close()

	result := make([]models.HistoricalFixedExpense, 0)
	for rows.Next() {
		var item models.HistoricalFixedExpense
		if err := rows.Scan(&item.PeriodID, &item.Year, &item.Month, &item.FixedExpenseID, &item.CategoryID, &item.Name, &item.ExpectedAmount); err != nil {
			return nil, fmt.Errorf("scan historical fixed expense: %w", err)
		}
		item.ExpectedAmount = money.Round(item.ExpectedAmount)
		result = append(result, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate historical fixed expenses: %w", err)
	}
	return result, nil
}

func (service *PeriodService) UpdateFixedExpense(ctx context.Context, userID, expenseID uuid.UUID, request models.UpdateFixedExpenseRequest) (models.FixedExpense, error) {
	name := strings.TrimSpace(request.Name)
	if name == "" || request.CategoryID == uuid.Nil {
		return models.FixedExpense{}, errors.New("name and category_id are required")
	}
	if request.ExpectedAmount < 0 {
		return models.FixedExpense{}, errors.New("expected_amount must be greater than or equal to 0")
	}

	var expense models.FixedExpense
	err := service.db.QueryRow(ctx, `
		UPDATE fixed_expenses expense
		SET category_id = category.id, name = $3, expected_amount = $4,
			is_active = COALESCE($5, expense.is_active)
		FROM categories category
		WHERE expense.id = $1 AND expense.user_id = $2
			AND category.id = $6 AND category.is_active = TRUE
			AND (category.created_by_user_id IS NULL OR category.created_by_user_id = $2)
		RETURNING expense.id, expense.category_id, expense.name, expense.expected_amount, expense.is_active`,
		expenseID, userID, name, money.Round(request.ExpectedAmount), request.IsActive, request.CategoryID,
	).Scan(&expense.ID, &expense.CategoryID, &expense.Name, &expense.ExpectedAmount, &expense.IsActive)
	if errors.Is(err, pgx.ErrNoRows) {
		return models.FixedExpense{}, ErrPeriodNotFound
	}
	if err != nil {
		return models.FixedExpense{}, fmt.Errorf("update fixed expense: %w", err)
	}
	expense.ExpectedAmount = money.Round(expense.ExpectedAmount)
	return expense, nil
}

func (service *PeriodService) UpsertFinancialSummary(ctx context.Context, userID, periodID uuid.UUID, request models.CreateFinancialSummaryRequest) (models.FinancialSummary, error) {
	if request.TotalIncome <= 0 {
		return models.FinancialSummary{}, errors.New("total_income must be greater than 0")
	}
	if request.MethodID == uuid.Nil {
		return models.FinancialSummary{}, errors.New("method_id is required")
	}

	tx, err := service.db.Begin(ctx)
	if err != nil {
		return models.FinancialSummary{}, fmt.Errorf("begin financial summary: %w", err)
	}
	defer tx.Rollback(ctx)

	var closed bool
	if err := tx.QueryRow(ctx, `SELECT is_closed FROM periods WHERE id = $1 AND user_id = $2 FOR UPDATE`, periodID, userID).Scan(&closed); errors.Is(err, pgx.ErrNoRows) {
		return models.FinancialSummary{}, ErrPeriodNotFound
	} else if err != nil {
		return models.FinancialSummary{}, fmt.Errorf("lock period: %w", err)
	}
	if closed {
		return models.FinancialSummary{}, ErrPeriodClosed
	}
	var methodExists bool
	if err := tx.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM methods
			WHERE id = $1 AND (created_by_user_id IS NULL OR created_by_user_id = $2)
		)`, request.MethodID, userID).Scan(&methodExists); err != nil {
		return models.FinancialSummary{}, fmt.Errorf("validate method: %w", err)
	}
	if !methodExists {
		return models.FinancialSummary{}, errors.New("method not found or not available")
	}

	var totalFixed float64
	if err := tx.QueryRow(ctx, `
		SELECT COALESCE(SUM(expected_amount), 0)
		FROM period_fixed_expenses
		WHERE period_id = $1`, periodID).Scan(&totalFixed); err != nil {
		return models.FinancialSummary{}, fmt.Errorf("calculate fixed expenses: %w", err)
	}

	totalIncome := money.Round(request.TotalIncome)
	totalFixed = money.Round(totalFixed)
	percentage := money.Round((totalFixed / totalIncome) * 100)
	summary := models.FinancialSummary{
		PeriodID:             periodID,
		TotalIncome:          totalIncome,
		TotalFixedExpenses:   totalFixed,
		FixedExpensesPercent: percentage,
		ExceedsFiftyPercent:  percentage > 50,
		HasSavingsGoal:       request.HasSavingsGoal,
		HasDebts:             request.HasDebts,
	}
	if err := service.populateDistribution(ctx, tx, userID, periodID, request, totalIncome, totalFixed, &summary); err != nil {
		return models.FinancialSummary{}, err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO financial_summaries (
			period_id, total_income, total_fixed_expenses,
			fixed_expenses_pct, has_savings_goal, has_debts, method_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (period_id) DO UPDATE SET
			total_income = EXCLUDED.total_income,
			total_fixed_expenses = EXCLUDED.total_fixed_expenses,
			fixed_expenses_pct = EXCLUDED.fixed_expenses_pct,
			has_savings_goal = EXCLUDED.has_savings_goal,
			has_debts = EXCLUDED.has_debts,
			method_id = EXCLUDED.method_id,
			updated_at = CURRENT_TIMESTAMP`,
		periodID, totalIncome, totalFixed, percentage,
		summary.HasSavingsGoal, summary.HasDebts, request.MethodID)
	if err != nil {
		return models.FinancialSummary{}, fmt.Errorf("save financial summary: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return models.FinancialSummary{}, fmt.Errorf("commit financial summary: %w", err)
	}
	return summary, nil
}

func (service *PeriodService) PreviewFinancialSummary(ctx context.Context, userID, periodID uuid.UUID, request models.CreateFinancialSummaryRequest) (models.FinancialSummary, error) {
	if request.TotalIncome <= 0 {
		return models.FinancialSummary{}, errors.New("total_income must be greater than 0")
	}
	if request.MethodID == uuid.Nil {
		return models.FinancialSummary{}, errors.New("method_id is required")
	}
	tx, err := service.db.Begin(ctx)
	if err != nil {
		return models.FinancialSummary{}, fmt.Errorf("begin summary preview: %w", err)
	}
	defer tx.Rollback(ctx)

	var closed bool
	if err := tx.QueryRow(ctx, `SELECT is_closed FROM periods WHERE id = $1 AND user_id = $2`, periodID, userID).Scan(&closed); errors.Is(err, pgx.ErrNoRows) {
		return models.FinancialSummary{}, ErrPeriodNotFound
	} else if err != nil {
		return models.FinancialSummary{}, fmt.Errorf("load period for preview: %w", err)
	}
	if closed {
		return models.FinancialSummary{}, ErrPeriodClosed
	}
	var totalFixed float64
	if err := tx.QueryRow(ctx, `SELECT COALESCE(SUM(expected_amount), 0) FROM period_fixed_expenses WHERE period_id = $1`, periodID).Scan(&totalFixed); err != nil {
		return models.FinancialSummary{}, fmt.Errorf("calculate preview fixed expenses: %w", err)
	}
	totalIncome := money.Round(request.TotalIncome)
	totalFixed = money.Round(totalFixed)
	percentage := money.Round((totalFixed / totalIncome) * 100)
	summary := models.FinancialSummary{
		PeriodID: periodID, TotalIncome: totalIncome, TotalFixedExpenses: totalFixed,
		FixedExpensesPercent: percentage, ExceedsFiftyPercent: percentage > 50,
		HasSavingsGoal: request.HasSavingsGoal, HasDebts: request.HasDebts,
	}
	if err := service.populateDistribution(ctx, tx, userID, periodID, request, totalIncome, totalFixed, &summary); err != nil {
		return models.FinancialSummary{}, err
	}
	return summary, nil
}

func (service *PeriodService) populateDistribution(ctx context.Context, tx pgx.Tx, userID, periodID uuid.UUID, request models.CreateFinancialSummaryRequest, income, fixed float64, summary *models.FinancialSummary) error {
	var methodName string
	if err := tx.QueryRow(ctx, `SELECT name FROM methods WHERE id = $1 AND (created_by_user_id IS NULL OR created_by_user_id = $2)`, request.MethodID, userID).Scan(&methodName); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("method not found or not available")
		}
		return fmt.Errorf("load method: %w", err)
	}
	summary.MethodName = methodName
	summary.AvailableAmount = money.Round(income - fixed)
	summary.HasDeficit = summary.AvailableAmount < 0
	if summary.HasDeficit {
		summary.Warnings = append(summary.Warnings, "Tus gastos fijos superan tu ingreso; no se asignarán importes positivos.")
	}
	if summary.ExceedsFiftyPercent {
		summary.Warnings = append(summary.Warnings, "Tus gastos fijos superan el 50% del ingreso.")
	}

	rows, err := tx.Query(ctx, `
		SELECT c.id, c.name, c.type, ma.percentage,
			EXISTS (SELECT 1 FROM fixed_expenses f WHERE f.user_id = $2 AND f.category_id = c.id AND f.is_active = TRUE)
			OR EXISTS (SELECT 1 FROM transactions t WHERE t.period_id = $3 AND t.category_id = c.id)
		FROM method_allocations ma
		JOIN categories c ON c.id = ma.category_id
		WHERE ma.method_id = $1
		ORDER BY ma.id`, request.MethodID, userID, periodID)
	if err != nil {
		return fmt.Errorf("load method allocations: %w", err)
	}
	defer rows.Close()

	type allocation struct{ item models.DistributionAllocation }
	allocations := make([]allocation, 0)
	usedTotal := 0.0
	unusedTotal := 0.0
	configuredTotal := 0.0
	for rows.Next() {
		var item models.DistributionAllocation
		if err := rows.Scan(&item.CategoryID, &item.CategoryName, &item.CategoryType, &item.ConfiguredPercent, &item.CategoryUsed); err != nil {
			return fmt.Errorf("scan method allocation: %w", err)
		}
		item.ConfiguredPercent = money.Round(item.ConfiguredPercent)
		configuredTotal += item.ConfiguredPercent
		if item.CategoryUsed {
			usedTotal += item.ConfiguredPercent
		} else {
			unusedTotal += item.ConfiguredPercent
			summary.Warnings = append(summary.Warnings, fmt.Sprintf("La categoría %s no tiene compromisos registrados; su porcentaje se redistribuye entre categorías usadas.", item.CategoryName))
		}
		allocations = append(allocations, allocation{item: item})
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate method allocations: %w", err)
	}
	if money.Round(configuredTotal) != 100 {
		summary.Warnings = append(summary.Warnings, "El método no suma exactamente 100%; revisa sus asignaciones antes de guardarlo.")
	}
	if usedTotal == 0 && unusedTotal > 0 {
		summary.Warnings = append(summary.Warnings, "El método no tiene categorías usadas en este período; se muestran sus porcentajes sin redistribución.")
	}

	var activeDebts int
	if err := tx.QueryRow(ctx, `SELECT COUNT(*) FROM debt_accounts WHERE user_id = $1 AND is_paid_off = FALSE`, userID).Scan(&activeDebts); err != nil {
		return fmt.Errorf("check active debts: %w", err)
	}
	hasDebtAllocation := false
	for _, current := range allocations {
		if current.item.CategoryType == "DEBT" {
			hasDebtAllocation = true
			break
		}
	}
	if request.HasDebts && activeDebts > 0 {
		if hasDebtAllocation {
			summary.Warnings = append(summary.Warnings, "La asignación de deuda se conserva como prioridad porque tienes deudas activas.")
		} else {
			summary.Warnings = append(summary.Warnings, "Tienes deudas activas, pero el método no incluye una categoría de deuda.")
		}
	}
	if request.HasSavingsGoal {
		summary.Warnings = append(summary.Warnings, "La asignación de ahorro se conserva porque activaste una meta de ahorro.")
	}
	for _, current := range allocations {
		item := current.item
		if item.CategoryUsed && usedTotal > 0 {
			item.EffectivePercent = money.Round(item.ConfiguredPercent + unusedTotal*(item.ConfiguredPercent/usedTotal))
		} else if !item.CategoryUsed && usedTotal > 0 {
			item.EffectivePercent = 0
		} else {
			item.EffectivePercent = item.ConfiguredPercent
		}
		if summary.HasDeficit {
			item.Amount = 0
		} else {
			item.Amount = money.Round(summary.AvailableAmount * item.EffectivePercent / 100)
		}
		summary.Allocations = append(summary.Allocations, item)
	}
	return nil
}
