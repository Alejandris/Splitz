import { useState } from 'react'
import { AuthProvider, useAuth } from '../features/auth/AuthProvider'
import { AuthPage } from '../pages/auth/AuthPage'
import { BudgetPage } from '../pages/budget/BudgetPage'
import { DashboardPage } from '../pages/dashboard/DashboardPage'
import { PeriodsPage } from '../pages/periods/PeriodsPage'
import { AppShell } from './layout/AppShell'

export type AppRoute = 'dashboard' | 'budget' | 'periods'

export function App() {
  return <AuthProvider><AuthenticatedApp /></AuthProvider>
}

function AuthenticatedApp() {
  const { user, loading } = useAuth()
  const [route, setRoute] = useState<AppRoute>('dashboard')

  if (!user) return <AuthPage />
  if (loading) return <main className="auth-splash" aria-label="Cargando Splitz"><span className="auth-mark">$</span></main>

  return (
    <AppShell activeRoute={route} onNavigate={setRoute}>
      {route === 'dashboard' ? <DashboardPage onStart={() => setRoute('periods')} /> : route === 'budget' ? <BudgetPage onOpenPeriod={() => setRoute('periods')} /> : <PeriodsPage />}
    </AppShell>
  )
}
