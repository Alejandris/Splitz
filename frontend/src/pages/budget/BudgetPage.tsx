import { useEffect, useState } from 'react'
import { closePeriod, listPeriods } from '../../features/periods/period.api'
import type { PeriodOverview } from '../../features/periods/period.types'

type BudgetPageProps = { onOpenPeriod: () => void }

export function BudgetPage({ onOpenPeriod }: BudgetPageProps) {
  const [periods, setPeriods] = useState<PeriodOverview[]>([])
  const [error, setError] = useState('')

  useEffect(() => {
    listPeriods().then(setPeriods).catch(() => setError('No se pudieron cargar los períodos registrados'))
  }, [])

  function formatMoney(value: number, currency: string) {
    return value.toLocaleString('es-MX', { style: 'currency', currency, minimumFractionDigits: 2 })
  }

  async function handleClose(periodId: string) {
    if (!window.confirm('¿Quieres cerrar este período? Después no podrás modificarlo.')) return
    try {
      const closed = await closePeriod(periodId)
      setPeriods((current) => current.map((period) => period.id === closed.id ? { ...period, is_closed: true } : period))
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo cerrar el período')
    }
  }

  return <section className="page budget-page">
    <div className="page-heading">
      <div><p className="eyebrow accent">HISTORIAL FINANCIERO</p><h1>Tus períodos,<br /><em>en perspectiva.</em></h1><p className="page-intro">Consulta los ciclos registrados y revisa el detalle de tu período actual.</p></div>
      <button className="primary-button compact-button" type="button" onClick={onOpenPeriod}>Abrir período <span>→</span></button>
    </div>
    {error && <p className="error-message">{error}</p>}
    {!periods.length && !error && <div className="empty-period"><span className="empty-period-mark">01</span><div><h2>Aún no hay períodos</h2><p>Abre tu primer período para comenzar a organizar tus ingresos y gastos.</p><button className="secondary-button" type="button" onClick={onOpenPeriod}>Abrir período</button></div></div>}
    {periods.length > 0 && <div className="period-history-grid">{periods.map((period) => <article className="period-history-card" key={period.id}><div className="period-history-card__top"><div><p className="eyebrow accent">{period.month}/{period.year}</p><h2>{period.start_date} — {period.end_date}</h2></div><div className="period-history-card__actions"><span className={period.is_closed ? 'workspace-status closed' : 'workspace-status'}>{period.is_closed ? 'Cerrado' : 'Abierto'}</span>{!period.is_closed && <button className="secondary-button close-period-button" type="button" onClick={() => handleClose(period.id)}>Cerrar período</button>}</div></div><div className="period-history-card__details"><div><small>Ingreso</small><strong>{period.has_summary ? formatMoney(period.total_income, period.currency_code) : 'Sin configurar'}</strong></div><div><small>Gastos fijos</small><strong>{period.has_summary ? formatMoney(period.total_fixed_expenses, period.currency_code) : '—'}</strong></div><div><small>Proporción fija</small><strong>{period.has_summary ? `${period.fixed_expenses_percent.toFixed(2)}%` : '—'}</strong></div></div></article>)}</div>}
  </section>
}
