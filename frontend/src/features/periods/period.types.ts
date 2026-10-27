export type Period = {
  id: string
  year: number
  month: number
  start_date: string
  end_date: string
  currency_code: string
  is_closed: boolean
  created_at: string
}

export type FixedExpense = {
  id?: string
  category_id: string
  name: string
  expected_amount: number
  is_active: boolean
}

export type DraftFixedExpense = Omit<FixedExpense, 'id' | 'is_active'> & { id?: string }

export type HistoricalFixedExpense = {
  period_id: string
  year: number
  month: number
  fixed_expense_id: string
  category_id: string
  name: string
  expected_amount: number
}

export type Category = {
  id: string
  name: string
  type: 'INCOME' | 'EXPENSE' | 'SAVING' | 'DEBT'
  created_by_user_id?: string
  is_active: boolean
}

export type FinancialSummary = {
  period_id: string
  total_income: number
  total_fixed_expenses: number
  fixed_expenses_percent: number
  exceeds_fifty_percent: boolean
  has_savings_goal: boolean
  has_debts: boolean
  available_amount: number
  has_deficit: boolean
  method_name: string
  warnings: string[]
  allocations: DistributionAllocation[]
}

export type DistributionAllocation = {
  category_id: string
  category_name: string
  category_type: string
  configured_percent: number
  effective_percent: number
  amount: number
  category_used: boolean
}

export type Method = {
  id: string
  name: string
  is_custom: boolean
  created_by_user_id?: string
}
