package service

import (
	"math"
	"testing"

	"splitz/internal/models"

	"github.com/google/uuid"
)

func TestCalculateDistributionAllocation_Traditional503020UsesAvailableAmount(t *testing.T) {
	income := 1750000.0
	fixed := 790000.0
	available := income - fixed
	if got := math.Round(fixed/income*100*100) / 100; got != 45.14 {
		t.Fatalf("fixed percentage: got %.2f, want 45.14", got)
	}

	allocations := []models.DistributionAllocation{
		{CategoryName: "Necesidades", ConfiguredPercent: 50, CategoryUsed: true},
		{CategoryName: "Ahorro", ConfiguredPercent: 20, CategoryUsed: false},
		{CategoryName: "Deseos", ConfiguredPercent: 30, CategoryUsed: false},
	}
	want := []float64{480000, 192000, 288000}
	var total float64
	for index, allocation := range allocations {
		got := calculateDistributionAllocation(available, allocation, false)
		if got.EffectivePercent != allocation.ConfiguredPercent {
			t.Errorf("%s effective percentage: got %.2f, want %.2f", allocation.CategoryName, got.EffectivePercent, allocation.ConfiguredPercent)
		}
		if got.Amount != want[index] {
			t.Errorf("%s amount: got %.2f, want %.2f", allocation.CategoryName, got.Amount, want[index])
		}
		total += got.Amount
	}
	if total+fixed != income {
		t.Fatalf("general total: got %.2f, want %.2f", total+fixed, income)
	}
}

func TestCalculateDistributionAllocation_DeficitProducesNoPositiveAmount(t *testing.T) {
	item := calculateDistributionAllocation(-1, models.DistributionAllocation{ConfiguredPercent: 50}, true)
	if item.Amount != 0 {
		t.Fatalf("deficit amount: got %.2f, want 0", item.Amount)
	}
}

func TestCalculateDistributionAllocation_DebtToggleDoesNotChangeConfiguredDistribution(t *testing.T) {
	available := 960000.0
	allocations := []models.DistributionAllocation{
		{CategoryName: "Necesidades", CategoryType: "EXPENSE", ConfiguredPercent: 50},
		{CategoryName: "Ahorro", CategoryType: "SAVING", ConfiguredPercent: 20},
		{CategoryName: "Deseos", CategoryType: "EXPENSE", ConfiguredPercent: 30},
	}
	want := []float64{480000, 192000, 288000}

	var total float64
	for index, allocation := range allocations {
		got := calculateDistributionAllocation(available, allocation, false)
		if got.Amount != want[index] {
			t.Errorf("%s amount: got %.2f, want %.2f", allocation.CategoryName, got.Amount, want[index])
		}
		total += got.Amount
	}
	if total != available {
		t.Fatalf("allocation total: got %.2f, want %.2f", total, available)
	}
}

func TestDebtPriorityWarningRequiresRealDebtAndDebtCategory(t *testing.T) {
	tests := []struct {
		name        string
		selected    bool
		activeDebts int
		hasCategory bool
		wantWarning bool
	}{
		{name: "button alone", selected: true, activeDebts: 0, hasCategory: true, wantWarning: false},
		{name: "real debt without debt category", selected: true, activeDebts: 1, hasCategory: false, wantWarning: false},
		{name: "real debt with debt category", selected: true, activeDebts: 1, hasCategory: true, wantWarning: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := debtPriorityWarning(test.selected, test.activeDebts, test.hasCategory)
			if (got != "") != test.wantWarning {
				t.Fatalf("warning presence: got %q, want warning=%t", got, test.wantWarning)
			}
		})
	}
}

func TestCalculateDraftFixedExpensesSumsOnlyCurrentDraft(t *testing.T) {
	first := uuid.New()
	second := uuid.New()
	total, err := calculateDraftFixedExpenses([]models.DraftFixedExpense{
		{ID: &first, CategoryID: uuid.New(), Name: "Deuda", ExpectedAmount: 700000},
		{ID: &second, CategoryID: uuid.New(), Name: "Netflix", ExpectedAmount: 30000},
		{CategoryID: uuid.New(), Name: "Gasolina", ExpectedAmount: 60000},
	})
	if err != nil {
		t.Fatal(err)
	}
	if total != 790000 {
		t.Fatalf("draft fixed expenses total: got %.2f, want 790000", total)
	}
}

func TestCalculateDraftFixedExpensesAllowsEmptyDraft(t *testing.T) {
	total, err := calculateDraftFixedExpenses(nil)
	if err != nil {
		t.Fatal(err)
	}
	if total != 0 {
		t.Fatalf("empty draft total: got %.2f, want 0", total)
	}
}

func TestCalculateDraftFixedExpensesIgnoresIncompleteDraft(t *testing.T) {
	total, err := calculateDraftFixedExpenses([]models.DraftFixedExpense{
		{Name: "", CategoryID: uuid.Nil, ExpectedAmount: 700000},
		{Name: "Válido", CategoryID: uuid.New(), ExpectedAmount: 30000},
	})
	if err != nil {
		t.Fatal(err)
	}
	if total != 30000 {
		t.Fatalf("incomplete draft total: got %.2f, want 30000", total)
	}
}
