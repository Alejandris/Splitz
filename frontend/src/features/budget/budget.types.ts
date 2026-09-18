export type BudgetMethod = 1 | 2 | 3

export type BudgetRequest = { salary: number; method: BudgetMethod }

export type BudgetResponse = {
  salary: number
  method: BudgetMethod
  method_name: string
  needs: number
  debt: number
  savings: number
  desires: number
  lifestyle: number
  investment: number
  ready_to_spend: number
}
