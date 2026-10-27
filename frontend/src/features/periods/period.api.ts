import { get, patch, post } from '../../shared/api/httpClient'
import type { Category, DraftFixedExpense, FinancialSummary, FixedExpense, HistoricalFixedExpense, Method, Period, PeriodOverview } from './period.types'

export function listCategories() {
  return get<Category[]>('/api/v1/categories')
}

export function createCategory(name: string, type: Category['type']) {
  return post<{ name: string; type: Category['type'] }, Category>('/api/v1/categories', { name, type })
}

export function createPeriod(year: number, month: number, start_date: string, end_date: string, currency_code: string) {
  return post<{ year: number; month: number; start_date: string; end_date: string; currency_code: string }, Period>('/api/v1/periods', { year, month, start_date, end_date, currency_code })
}

export function listPeriods() {
  return get<PeriodOverview[]>('/api/v1/periods')
}

export function closePeriod(periodId: string) {
  return patch<Record<string, never>, Period>(`/api/v1/periods/${periodId}/close`, {})
}

export function createFixedExpense(category_id: string, name: string, expected_amount: number) {
  return post<{ category_id: string; name: string; expected_amount: number }, FixedExpense>('/api/v1/fixed-expenses', { category_id, name, expected_amount })
}

export function listFixedExpenses() {
  return get<FixedExpense[]>('/api/v1/fixed-expenses')
}

export function updateFixedExpense(expense: FixedExpense, is_active: boolean) {
  return patch<{
    category_id: string
    name: string
    expected_amount: number
    is_active: boolean
  }, FixedExpense>(`/api/v1/fixed-expenses/${expense.id}`, {
    category_id: expense.category_id,
    name: expense.name,
    expected_amount: expense.expected_amount,
    is_active,
  })
}

export function listHistoricalFixedExpenses() {
  return get<HistoricalFixedExpense[]>('/api/v1/fixed-expenses/history')
}

export function listMethods() {
  return get<Method[]>('/api/v1/methods')
}

export function saveFinancialSummary(periodId: string, total_income: number, has_savings_goal: boolean, has_debts: boolean, method_id: string, fixed_expenses: DraftFixedExpense[]) {
  return post<{ total_income: number; has_savings_goal: boolean; has_debts: boolean; method_id: string; fixed_expenses: DraftFixedExpense[] }, FinancialSummary>(
    `/api/v1/periods/${periodId}/financial-summary`,
    { total_income, has_savings_goal, has_debts, method_id, fixed_expenses },
  )
}

export function previewFinancialSummary(periodId: string, total_income: number, has_savings_goal: boolean, has_debts: boolean, method_id: string, fixed_expenses: DraftFixedExpense[]) {
  return post<{ total_income: number; has_savings_goal: boolean; has_debts: boolean; method_id: string; fixed_expenses: DraftFixedExpense[] }, FinancialSummary>(
    `/api/v1/periods/${periodId}/financial-summary/preview`,
    { total_income, has_savings_goal, has_debts, method_id, fixed_expenses },
  )
}
