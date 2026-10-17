import { post } from '../../shared/api/httpClient'
import type { BudgetRequest, BudgetResponse } from './budget.types'

export function calculateBudget(request: BudgetRequest) {
  return post<BudgetRequest, BudgetResponse>('/api/v1/budget/calculate', request)
}
