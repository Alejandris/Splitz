type DashboardPageProps = { onStart: () => void }

export function DashboardPage({ onStart }: DashboardPageProps) {
  return (
    <section className="page dashboard-page">
      <div className="hero">
        <div>
          <p className="eyebrow accent">BUENAS, ALEJANDRO</p>
          <h1>Tu dinero,<br /><em>dividido con intención.</em></h1>
          <p className="hero-copy">Construye claridad sobre tus ingresos y convierte tus metas en hábitos.</p>
          <button className="primary-button" onClick={onStart}>Crear mi presupuesto <span>→</span></button>
        </div>
        <div className="hero-orbit" aria-hidden="true"><div className="orbit-ring ring-one" /><div className="orbit-ring ring-two" /><div className="orbit-core">$</div></div>
      </div>
      <div className="section-heading"><div><p className="eyebrow">PRÓXIMAMENTE</p><h2>Tu panorama financiero</h2></div><span className="muted">Aún no hay datos</span></div>
      <div className="empty-grid"><article><span className="card-icon">↗</span><h3>Ingresos</h3><p>Registra tu salario neto para comenzar.</p></article><article><span className="card-icon">◌</span><h3>Metas</h3><p>Define qué quieres lograr con tu dinero.</p></article><article><span className="card-icon">⌁</span><h3>Progreso</h3><p>Visualiza tus avances mes a mes.</p></article></div>
    </section>
  )
}
