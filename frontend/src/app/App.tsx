import { useState } from 'react'
import { AuthProvider, useAuth } from '../features/auth/AuthProvider'
import { AuthPage } from '../pages/auth/AuthPage'
import { BudgetPage } from '../pages/budget/BudgetPage'
import { DashboardPage } from '../pages/dashboard/DashboardPage'
import { AppShell } from './layout/AppShell'

export type AppRoute = 'dashboard' | 'budget'

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
      {route === 'dashboard' ? <DashboardPage onStart={() => setRoute('budget')} /> : <BudgetPage />}
    </AppShell>
  )
}
