import type { ReactNode } from 'react'
import type { AppRoute } from '../App'

type AppShellProps = {
  activeRoute: AppRoute
  onNavigate: (route: AppRoute) => void
  children: ReactNode
}

export function AppShell({ activeRoute, onNavigate, children }: AppShellProps) {
  return (
    <div className="app-shell">
      <aside className="sidebar">
        <div className="brand"><span className="brand-mark">S</span><span>splitz</span></div>
        <p className="sidebar-label">Workspace</p>
        <nav aria-label="Navegación principal">
          <button className={`nav-item ${activeRoute === 'dashboard' ? 'active' : ''}`} onClick={() => onNavigate('dashboard')}>
            <span>⌂</span> Resumen
          </button>
          <button className={`nav-item ${activeRoute === 'budget' ? 'active' : ''}`} onClick={() => onNavigate('budget')}>
            <span>◈</span> Presupuesto
          </button>
        </nav>
        <div className="sidebar-footer"><span className="status-dot" /> API conectada</div>
      </aside>
      <main className="main-content">
        <header className="topbar">
          <span className="eyebrow">CONTROL FINANCIERO</span>
          <div className="avatar">A</div>
        </header>
        {children}
      </main>
    </div>
  )
}
