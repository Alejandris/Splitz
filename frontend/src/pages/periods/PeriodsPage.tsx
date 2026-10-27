import { useEffect, useMemo, useState } from 'react'
import { createCategory, createFixedExpense, createPeriod, listCategories, listFixedExpenses, listHistoricalFixedExpenses, listMethods, previewFinancialSummary, saveFinancialSummary } from '../../features/periods/period.api'
import type { Category, FinancialSummary, FixedExpense, HistoricalFixedExpense, Method, Period } from '../../features/periods/period.types'

function CategoryIcon({ type }: { type: Category['type'] }) {
  const path = type === 'SAVING'
    ? 'M12 3v18M5 8h14M5 16h14M7 3h10v18H7z'
    : type === 'DEBT'
      ? 'M4 7h16v10H4zM7 12h4M16 12h1'
      : 'M5 4h14v16H5zM8 8h8M8 12h8M8 16h5'
  return <svg className="category-icon" viewBox="0 0 24 24" aria-hidden="true"><path d={path} /></svg>
}

export function PeriodsPage() {
  const today = new Date()
  const [year, setYear] = useState(String(today.getFullYear()))
  const [month, setMonth] = useState(String(today.getMonth() + 1))
  const [startDate, setStartDate] = useState(`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}-01`)
  const [endDate, setEndDate] = useState(`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}-${String(new Date(today.getFullYear(), today.getMonth() + 1, 0).getDate()).padStart(2, '0')}`)
  const [currency, setCurrency] = useState('MXN')
  const [income, setIncome] = useState('')
  const [period, setPeriod] = useState<Period | null>(null)
  const [expenses, setExpenses] = useState<FixedExpense[]>([])
  const [history, setHistory] = useState<HistoricalFixedExpense[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [methods, setMethods] = useState<Method[]>([])
  const [categoryId, setCategoryId] = useState('')
  const [categoryName, setCategoryName] = useState('')
  const [categoryType, setCategoryType] = useState<Category['type']>('EXPENSE')
  const [expenseName, setExpenseName] = useState('')
  const [expenseAmount, setExpenseAmount] = useState('')
  const [selectedMethod, setSelectedMethod] = useState('')
  const [hasSavingsGoal, setHasSavingsGoal] = useState(false)
  const [hasDebts, setHasDebts] = useState(false)
  const [showExpenseModal, setShowExpenseModal] = useState(false)
  const [showHistoryModal, setShowHistoryModal] = useState(false)
  const [showCategoryModal, setShowCategoryModal] = useState(false)
  const [summary, setSummary] = useState<{ percent: number; warning: boolean } | null>(null)
  const [preview, setPreview] = useState<FinancialSummary | null>(null)
  const [error, setError] = useState('')

  useEffect(() => {
    Promise.all([listFixedExpenses(), listCategories(), listMethods(), listHistoricalFixedExpenses()])
      .then(([loadedExpenses, loadedCategories, loadedMethods, loadedHistory]) => {
        setExpenses(loadedExpenses)
        setHistory(loadedHistory)
        setCategories(loadedCategories)
        setMethods(loadedMethods)
        setCategoryId(loadedCategories.find((category) => category.type === 'EXPENSE' || category.type === 'DEBT')?.id ?? '')
        setSelectedMethod(loadedMethods[0]?.id ?? '')
      })
      .catch(() => setError('No se pudieron cargar las opciones del período'))
  }, [])

  const expenseCategories = useMemo(
    () => categories.filter((category) => category.type === 'EXPENSE' || category.type === 'DEBT'),
    [categories],
  )

  function formatMoney(value: number) {
    return value.toLocaleString('es-MX', { style: 'currency', currency, minimumFractionDigits: 2 })
  }

  function formatInputAmount(value: string) {
    const normalized = value.replace(/[^\d.]/g, '')
    const [integer = '', decimals] = normalized.split('.')
    const formattedInteger = (integer || '0').replace(/^0+(?=\d)/, '').replace(/\B(?=(\d{3})+(?!\d))/g, ',')
    return decimals === undefined ? formattedInteger : `${formattedInteger}.${decimals.slice(0, 2)}`
  }

  function parseAmount(value: string) {
    return Number(value.replace(/,/g, ''))
  }

  function repeatExpense(expense: HistoricalFixedExpense) {
    setCategoryId(expense.category_id)
    setExpenseName(expense.name)
    setExpenseAmount(String(expense.expected_amount))
    setShowHistoryModal(false)
    setShowExpenseModal(true)
  }

  async function handlePeriod(event: React.FormEvent) {
    event.preventDefault()
    setError('')
    try {
      setPeriod(await createPeriod(Number(year), Number(month), startDate, endDate, currency))
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo abrir el período')
    }
  }

  async function handleExpense(event: React.FormEvent) {
    event.preventDefault()
    setError('')
    try {
      const created = await createFixedExpense(categoryId, expenseName, parseAmount(expenseAmount))
      setExpenses((current) => [...current, created])
      setExpenseName('')
      setExpenseAmount('')
      setShowExpenseModal(false)
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo guardar el gasto fijo')
    }
  }

  async function handleCategory(event: React.FormEvent) {
    event.preventDefault()
    setError('')
    try {
      const created = await createCategory(categoryName, categoryType)
      setCategories((current) => [...current, created])
      setCategoryId(created.id)
      setCategoryName('')
      setShowCategoryModal(false)
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo crear la categoría')
    }
  }

  async function handleSummary(event: React.FormEvent) {
    event.preventDefault()
    if (!period) return
    setError('')
    try {
      const result = await previewFinancialSummary(period.id, parseAmount(income), hasSavingsGoal, hasDebts, selectedMethod)
      setPreview(result)
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo calcular el resumen')
    }

  }

  async function confirmSummary() {
    if (!period || !preview) return
    setError('')
    try {
      const result = await saveFinancialSummary(period.id, parseAmount(income), hasSavingsGoal, hasDebts, selectedMethod)
      setSummary({ percent: result.fixed_expenses_percent, warning: result.exceeds_fifty_percent })
      setPreview(null)
    } catch (requestError) {
      setError(requestError instanceof Error ? requestError.message : 'No se pudo guardar el período')
    }
  }

  return <section className="page period-page">
    <div className="page-heading">
      <div><p className="eyebrow accent">PLAN MENSUAL</p><h1>Tu dinero, <em>en orden.</em></h1><p className="page-intro">Abre el período, elige tus compromisos y deja que tu presupuesto se ajuste a tu realidad.</p></div>
      <span className="step-count">01 <span>/ 03</span></span>
    </div>

    <form className="period-start" onSubmit={handlePeriod}>
      <div><label className="field-label" htmlFor="period-month">Período que quieres organizar</label><p className="field-hint">Podrás volver a editar tus plantillas sin cambiar meses cerrados.</p></div>
      <div className="period-selects"><select id="period-month" value={month} onChange={(event) => setMonth(event.target.value)} aria-label="Mes"><option value="1">Enero</option><option value="2">Febrero</option><option value="3">Marzo</option><option value="4">Abril</option><option value="5">Mayo</option><option value="6">Junio</option><option value="7">Julio</option><option value="8">Agosto</option><option value="9">Septiembre</option><option value="10">Octubre</option><option value="11">Noviembre</option><option value="12">Diciembre</option></select><input id="period-year" type="number" min="2000" max="9999" value={year} onChange={(event) => setYear(event.target.value)} aria-label="Año" required /><label className="date-field">Inicio<input type="date" value={startDate} onChange={(event) => setStartDate(event.target.value)} required /></label><label className="date-field">Fin<input type="date" value={endDate} onChange={(event) => setEndDate(event.target.value)} required /></label><select value={currency} onChange={(event) => setCurrency(event.target.value)} aria-label="Moneda"><option value="MXN">MXN · Peso mexicano</option><option value="COP">COP · Peso colombiano</option><option value="USD">USD · Dólar</option><option value="EUR">EUR · Euro</option></select><button className="primary-button compact-button" type="submit">{period ? 'Actualizar período' : 'Abrir período'} <span>→</span></button></div>
    </form>

    {period && <form className="period-workspace" onSubmit={handleSummary}>
      <div className="workspace-heading"><div><p className="eyebrow accent">PASO 02 · INGRESO Y COMPROMISOS</p><h2>{period.month}/{period.year}: construyamos tu base</h2></div><span className="workspace-status">Período abierto</span></div>
      <div className="income-row"><label className="field-label" htmlFor="period-income">Ingreso neto del mes</label><div className="money-input"><span>$</span><input id="period-income" type="text" inputMode="decimal" value={income} onChange={(event) => setIncome(formatInputAmount(event.target.value))} placeholder="0.00" required /><span className="currency">{currency}</span></div></div>
      <div className="goal-options"><label><input type="checkbox" checked={hasSavingsGoal} onChange={(event) => setHasSavingsGoal(event.target.checked)} /> Tengo una meta de ahorro</label><label><input type="checkbox" checked={hasDebts} onChange={(event) => setHasDebts(event.target.checked)} /> Tengo deudas activas</label></div>
      <div className="section-heading"><div><p className="eyebrow accent">GASTOS FIJOS</p><h3>Lo que ya sabes que vas a pagar</h3></div><button className="secondary-button" type="button" onClick={() => setShowExpenseModal(true)}>Agregar gasto <span>+</span></button></div>
      <div className="category-board">{expenseCategories.map((category) => { const items = expenses.filter((expense) => expense.category_id === category.id && expense.is_active); return <div className="category-tile" key={category.id}><div className="tile-top"><span className="tile-icon"><CategoryIcon type={category.type} /></span><span className="tile-count">{items.length} {items.length === 1 ? 'gasto' : 'gastos'}</span></div><h3>{category.name}</h3>{items.length ? <ul>{items.map((expense) => <li key={expense.id}><span>{expense.name}</span><strong>{formatMoney(expense.expected_amount)}</strong></li>)}</ul> : <p className="tile-empty">Sin gastos registrados</p>}</div> })}<button className="category-tile add-tile" type="button" onClick={() => setShowExpenseModal(true)}><span className="add-icon">+</span><strong>Agregar otra categoría</strong><small>Usa una categoría existente o crea una propia.</small></button></div>
      <button className="repeat-button" type="button" onClick={() => setShowHistoryModal(true)}>Repetir gastos de otro mes <span>↗</span></button>
      <div className="section-heading method-heading"><div><p className="eyebrow accent">PASO 03 · MÉTODO</p><h3>Elige cómo repartir tu disponible</h3></div><button className="text-button" type="button" onClick={() => setShowCategoryModal(true)}>Crear categoría</button></div>
      <div className="method-picker">{methods.map((method) => <button type="button" key={method.id} className={`method-option ${selectedMethod === method.id ? 'selected' : ''}`} onClick={() => setSelectedMethod(method.id)}><span><strong>{method.name}</strong><small>{method.is_custom ? 'Tu método personalizado' : 'Método global Splitz'}</small></span><span className="radio" aria-hidden="true">{selectedMethod === method.id ? '✓' : ''}</span></button>)}</div>
      <button className="primary-button calculate-button" type="submit">Calcular mi distribución <span>→</span></button>
      {summary && <p className={summary.warning ? 'error-message summary-message' : 'muted summary-message'}>Período guardado. Tus gastos fijos representan {summary.percent.toFixed(2)}% del ingreso.{summary.warning ? ' Superan el 50% recomendado.' : ' Tu base está dentro del rango recomendado.'}</p>}
      {preview && <div className="summary-review"><p className="eyebrow accent">REVISIÓN ANTES DE GUARDAR</p><h3>{preview.method_name}</h3><p>Disponible después de gastos fijos: <strong>{formatMoney(preview.available_amount)}</strong></p><p>Gastos fijos: {preview.fixed_expenses_percent.toFixed(2)}% · Moneda: {currency}</p><ul>{preview.allocations.map((allocation) => <li key={allocation.category_id}><span>{allocation.category_name} · {allocation.effective_percent.toFixed(2)}%</span><strong>{formatMoney(allocation.amount)}</strong></li>)}</ul>{preview.warnings.map((warning) => <p className="summary-warning" key={warning}>{warning}</p>)}<div className="review-actions"><button className="primary-button" type="button" onClick={confirmSummary}>Guardar período definitivo</button><button className="text-button" type="button" onClick={() => setPreview(null)}>Editar</button></div></div>}
    </form>}

    {!period && <div className="empty-period"><span className="empty-period-mark">01</span><div><h2>Empieza por abrir tu período</h2><p>Al hacerlo verás tus categorías, gastos fijos y métodos en una sola vista.</p></div></div>}
    {error && <p className="error-message">{error}</p>}

    {showExpenseModal && <div className="modal-backdrop" role="presentation" onMouseDown={(event) => { if (event.currentTarget === event.target) setShowExpenseModal(false) }}><form className="modal" onSubmit={handleExpense}><button className="modal-close" type="button" onClick={() => setShowExpenseModal(false)} aria-label="Cerrar">×</button><p className="eyebrow accent">NUEVO COMPROMISO</p><h2>Agrega un gasto fijo</h2><p className="modal-copy">La plantilla se podrá reutilizar en tus próximos períodos.</p><label className="field-label" htmlFor="expense-category">Categoría</label><select id="expense-category" value={categoryId} onChange={(event) => setCategoryId(event.target.value)} required>{expenseCategories.map((category) => <option value={category.id} key={category.id}>{category.name}</option>)}</select><label className="field-label" htmlFor="expense-name">Nombre del gasto</label><input id="expense-name" value={expenseName} onChange={(event) => setExpenseName(event.target.value)} placeholder="Ej. Arriendo, internet..." required /><label className="field-label" htmlFor="expense-amount">Valor mensual</label><div className="modal-money"><span>$</span><input id="expense-amount" type="text" inputMode="decimal" value={expenseAmount} onChange={(event) => setExpenseAmount(formatInputAmount(event.target.value))} placeholder="0.00" required /><span>{currency}</span></div><button className="primary-button" type="submit">Guardar gasto <span>→</span></button><button className="text-button modal-category" type="button" onClick={() => { setShowExpenseModal(false); setShowCategoryModal(true) }}>Crear una categoría nueva</button></form></div>}
    {showCategoryModal && <div className="modal-backdrop" role="presentation" onMouseDown={(event) => { if (event.currentTarget === event.target) setShowCategoryModal(false) }}><form className="modal" onSubmit={handleCategory}><button className="modal-close" type="button" onClick={() => setShowCategoryModal(false)} aria-label="Cerrar">×</button><p className="eyebrow accent">PERSONALIZA TU TABLERO</p><h2>Nueva categoría</h2><label className="field-label" htmlFor="category-name">Nombre</label><input id="category-name" value={categoryName} onChange={(event) => setCategoryName(event.target.value)} placeholder="Ej. Mascotas" required /><label className="field-label" htmlFor="category-type">Tipo</label><select id="category-type" value={categoryType} onChange={(event) => setCategoryType(event.target.value as Category['type'])}><option value="EXPENSE">Gasto</option><option value="SAVING">Ahorro</option><option value="DEBT">Deuda</option><option value="INCOME">Ingreso</option></select><button className="primary-button" type="submit">Crear categoría <span>→</span></button></form></div>}
    {showHistoryModal && <div className="modal-backdrop" role="presentation" onMouseDown={(event) => { if (event.currentTarget === event.target) setShowHistoryModal(false) }}><div className="modal history-modal"><button className="modal-close" type="button" onClick={() => setShowHistoryModal(false)} aria-label="Cerrar">×</button><p className="eyebrow accent">HISTORIAL</p><h2>Repite un gasto</h2><p className="modal-copy">Elige un gasto de un período anterior para convertirlo en una plantilla activa.</p>{history.length ? <ul className="history-list">{history.map((expense) => <li key={`${expense.period_id}-${expense.fixed_expense_id}`}><div><strong>{expense.name}</strong><small>{expense.month}/{expense.year} · {formatMoney(expense.expected_amount)}</small></div><button className="text-button" type="button" onClick={() => repeatExpense(expense)}>Usar</button></li>)}</ul> : <p className="tile-empty">Aún no hay snapshots históricos disponibles.</p>}</div></div>}
  </section>
}
