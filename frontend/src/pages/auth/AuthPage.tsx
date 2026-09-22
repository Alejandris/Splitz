import { useEffect, useState } from 'react'
import { useAuth } from '../../features/auth/AuthProvider'

const AUTH_SPLASH_DURATION = 8_000

export function AuthPage() {
  const { error, signIn, signUp, signInWithGoogle } = useAuth()
  const [showForm, setShowForm] = useState(false)
  const [registering, setRegistering] = useState(false)
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [submitting, setSubmitting] = useState(false)

  useEffect(() => {
    const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
    const timeout = window.setTimeout(() => setShowForm(true), reducedMotion ? 0 : AUTH_SPLASH_DURATION)
    return () => window.clearTimeout(timeout)
  }, [])

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault()
    setSubmitting(true)
    try {
      if (registering) await signUp(name, email, password)
      else await signIn(email, password)
    } catch {
      // AuthProvider exposes the translated error beside the form.
    } finally {
      setSubmitting(false)
    }
  }

  if (!showForm) return <main className="auth-splash" aria-label="Cargando Splitz">
    <span className="auth-mark">$</span>
    <button className="auth-skip" type="button" onClick={() => setShowForm(true)}>Continuar</button>
  </main>

  return <main className="auth-page">
    <section className="auth-panel">
      <div className="brand auth-brand"><span className="brand-mark">S</span><span>splitz</span></div>
      <p className="eyebrow accent">CONTROL FINANCIERO</p>
      <h1>{registering ? <>Empieza a darle<br /><em>intención a tu dinero.</em></> : <>Tu dinero,<br /><em>dividido con intención.</em></>}</h1>
      <p className="auth-copy">Una vista clara para tomar mejores decisiones con cada peso.</p>
      <form className="auth-form" onSubmit={handleSubmit}>
        {registering && <label className="auth-field">Nombre visible<input value={name} onChange={(event) => setName(event.target.value)} autoComplete="name" required /></label>}
        <label className="auth-field">Correo electrónico<input type="email" value={email} onChange={(event) => setEmail(event.target.value)} autoComplete="email" required /></label>
        <label className="auth-field">Contraseña<input type="password" value={password} onChange={(event) => setPassword(event.target.value)} autoComplete={registering ? 'new-password' : 'current-password'} minLength={6} required /></label>
        <button className="primary-button auth-submit" disabled={submitting}>{submitting ? 'Conectando...' : registering ? 'Crear cuenta' : 'Ingresar'} <span>→</span></button>
        {error && <p className="error-message" role="alert">{error}</p>}
      </form>
      <div className="auth-divider"><span>o continúa con</span></div>
      <button type="button" className="secondary-button google-button" onClick={() => void signInWithGoogle()} disabled={submitting}>G&nbsp;&nbsp; Google</button>
      <button type="button" className="auth-switch" onClick={() => setRegistering(!registering)}>{registering ? 'Ya tengo una cuenta' : 'Crear una cuenta nueva'}</button>
    </section>
    <div className="auth-orbit" aria-hidden="true"><div className="orbit-ring ring-one" /><div className="orbit-ring ring-two" /><div className="orbit-core">$</div></div>
  </main>
}