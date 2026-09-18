import { useState } from 'react'
import { calculateBudget } from '../../features/budget/budget.api'
import type { BudgetMethod, BudgetResponse } from '../../features/budget/budget.types'

const methods: Array<{ id: BudgetMethod; name: string; description: string }> = [
  { id: 1, name: 'Modo enfoque', description: 'Prioriza necesidades y deudas.' },
  { id: 2, name: 'Modo inversionista', description: 'Acelera ahorro e inversión.' },
  { id: 3, name: 'Modo disfrute', description: 'Equilibra vida y objetivos.' },
]

const categories = [
  ['needs', 'Necesidades', 'var(--teal)'], ['debt', 'Deudas', 'var(--orange)'], ['savings', 'Ahorro', 'var(--blue)'],
  ['desires', 'Deseos', 'var(--pink)'], ['lifestyle', 'Estilo de vida', 'var(--purple)'], ['investment', 'Inversión', 'var(--green)'],
] as const

export function BudgetPage() {
  const [salary, setSalary] = useState('3000')
  const [method, setMethod] = useState<BudgetMethod>(1)
  const [result, setResult] = useState<BudgetResponse | null>(null)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault()
    setError('')
    setLoading(true)
    try {
      setResult(await calculateBudget({ salary: Number(salary), method }))
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo calcular el presupuesto')
    } finally {
      setLoading(false)
    }
  }

  return <section className="page budget-page">
    <div className="page-heading"><div><p className="eyebrow accent">NUEVO PRESUPUESTO</p><h1>Haz que cada peso<br /><em>tenga un propósito.</em></h1></div><span className="step-count">01 <span>/ 01</span></span></div>
    <form className="budget-form" onSubmit={handleSubmit}>
      <label className="field-label" htmlFor="salary">¿Cuál es tu ingreso neto mensual?</label>
      <div className="salary-input"><span>$</span><input id="salary" type="number" min="1" step="0.01" value={salary} onChange={(event) => setSalary(event.target.value)} required /><span className="currency">MXN</span></div>
      <div className="field-label method-label">Elige una metodología</div>
      <div className="method-list">{methods.map((item) => <button type="button" key={item.id} className={`method-card ${method === item.id ? 'selected' : ''}`} onClick={() => setMethod(item.id)}><span className="method-number">0{item.id}</span><span><strong>{item.name}</strong><small>{item.description}</small></span><span className="radio">{method === item.id ? '✓' : ''}</span></button>)}</div>
      <button className="primary-button calculate-button" type="submit" disabled={loading}>{loading ? 'Calculando...' : 'Calcular distribución'} <span>→</span></button>
      {error && <p className="error-message">{error}</p>}
    </form>
    {result && <div className="result-card"><div className="result-heading"><div><p className="eyebrow accent">TU DISTRIBUCIÓN</p><h2>{result.method_name}</h2></div><div className="ready-amount"><small>LISTO PARA GASTAR</small><strong>${result.ready_to_spend.toLocaleString('es-MX', { minimumFractionDigits: 2 })}</strong></div></div><div className="category-grid">{categories.map(([key, label, color]) => <div className="category" key={key}><div className="category-label"><span className="category-dot" style={{ background: color }} />{label}</div><strong>${result[key].toLocaleString('es-MX', { minimumFractionDigits: 2 })}</strong></div>)}</div></div>}
  </section>
}
