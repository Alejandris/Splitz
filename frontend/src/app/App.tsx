import { useState } from 'react'
import { BudgetPage } from '../pages/budget/BudgetPage'
import { DashboardPage } from '../pages/dashboard/DashboardPage'
import { AppShell } from './layout/AppShell'

export type AppRoute = 'dashboard' | 'budget'

export function App() {
  const [route, setRoute] = useState<AppRoute>('dashboard')

  return (
    <AppShell activeRoute={route} onNavigate={setRoute}>
      {route === 'dashboard' ? <DashboardPage onStart={() => setRoute('budget')} /> : <BudgetPage />}
    </AppShell>
  )
}
