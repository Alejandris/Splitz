import type { ReactNode } from 'react'
import type { AppRoute } from '../App'
import { useAuth } from '../../features/auth/AuthProvider'

type AppShellProps = {
  activeRoute: AppRoute
  onNavigate: (route: AppRoute) => void
  children: ReactNode
}

export function AppShell({ activeRoute, onNavigate, children }: AppShellProps) {
  const { profile, user, signOutUser } = useAuth()
  const profileName = profile?.name || user?.displayName || user?.email || 'Usuario'

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
          <div className="profile-menu"><div className="avatar">{profileName.charAt(0).toUpperCase()}</div><span>{profileName}</span><button className="logout-button" onClick={() => void signOutUser()}>Salir</button></div>
        </header>
        {children}
      </main>
    </div>
  )
}
