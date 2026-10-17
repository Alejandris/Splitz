# Sistema de diseño de Splitz

> Fuente única de verdad para la identidad visual y la experiencia de usuario de Splitz.

## 1. Dirección visual

Splitz es una aplicación de planificación financiera. Su interfaz debe transmitir una sensación **tranquila, confiable y clara**, ayudando a que las decisiones sobre el dinero se sientan ordenadas y alcanzables.

### Principios

1. **Claridad antes que decoración.** Cada pantalla debe tener una acción principal y una jerarquía fácil de escanear.
2. **Confianza sin rigidez.** La interfaz debe sentirse profesional, pero cercana y humana.
3. **Calma visual.** Usar espacios generosos, contrastes controlados y pocos elementos compitiendo entre sí.
4. **Progreso visible.** Los resultados, estados y próximos pasos deben ser fáciles de identificar.
5. **Consistencia.** Los mismos tokens y patrones deben utilizarse en dashboard, presupuesto y futuras funcionalidades.

## 2. Paleta de colores

Los colores se utilizan mediante nombres semánticos. Los componentes no deben depender directamente de nombres como `blue`, `green` u `orange`, porque esos nombres describen el tono y no su intención.

| Token | HEX | Uso principal |
| --- | --- | --- |
| `--color-ink` | `#061A24` | Fondo profundo del shell, navegación y zonas de alta concentración. |
| `--color-ocean` | `#0B3A53` | Superficies secundarias, bordes con presencia y estados activos sobrios. |
| `--color-primary` | `#0FA3B1` | Acción primaria, enlaces, selección activa y elementos interactivos principales. |
| `--color-sky` | `#B5E2FA` | Acentos suaves, fondos de apoyo y visualizaciones que necesiten ligereza. |
| `--color-surface` | `#F3FAFF` | Superficie clara, texto oscuro sobre fondos profundos y estados destacados. |
| `--color-success` | `#9EE4A8` | Texto pequeño de confirmacion, indicadores discretos y simbolos de estado positivo. |

### Reglas de uso

- `--color-ink` y `--color-ocean` deben formar la base estructural, no competir con el contenido.
- `--color-primary` se reserva para acciones que el usuario puede ejecutar. No usarlo para decorar cada bloque.
- `--color-success` comunica un estado positivo secundario; debe limitarse a letras pequeñas, indicadores discretos y simbolos. No usarlo en botones principales, iconos de marca, importes destacados ni como color general de marca.
- `--color-surface` debe mantener contraste suficiente cuando se use como texto o superficie.
- `--color-sky` funciona como apoyo visual y no sustituye el foco de teclado ni los mensajes de estado.
- No introducir nuevos colores categóricos sin documentar primero su significado y su contraste.
- Para categorías financieras, priorizar etiquetas, iconos y patrones además del color. La información no debe depender únicamente del color.

## 3. Tipografía

### Familias oficiales

- **Títulos:** `Cormorant Garamond`. Serif elegante, distintiva y de alta gama para encabezados, frases de portada y nombres de secciones.
- **Cuerpo e interfaz:** `Montserrat`. Fuente sans serif para párrafos, navegación, botones, formularios, métricas, etiquetas y mensajes de estado.

`Open Sans` queda como alternativa descartada para esta versión del sistema. No debe mezclarse con Montserrat en la misma interfaz.

### Jerarquía tipográfica

| Elemento | Familia | Peso recomendado | Tamaño orientativo | Altura de línea |
| --- | --- | --- | --- | --- |
| H1 de página | Cormorant Garamond | 600-700 | 3.25rem a 4.5rem | 0.95-1.05 |
| H2 de sección | Cormorant Garamond | 600 | 1.75rem a 2.25rem | 1.05-1.15 |
| H3 de tarjeta | Cormorant Garamond | 600 | 1.35rem a 1.6rem | 1.15 |
| Texto principal | Montserrat | 400-500 | 1rem | 1.55-1.7 |
| Texto secundario | Montserrat | 400 | 0.8125rem a 0.9375rem | 1.45-1.6 |
| Botón y navegación | Montserrat | 600-700 | 0.875rem a 1rem | 1.2 |
| Etiqueta técnica | Montserrat | 600 | 0.6875rem a 0.75rem | 1.3 |
| Métrica financiera | Montserrat | 600-700 | 1.5rem a 2.25rem | 1.1 |

### Reglas tipográficas

- Usar Cormorant Garamond para crear jerarquía y personalidad, no para párrafos largos ni controles pequeños.
- Usar Montserrat para cualquier texto que requiera lectura rápida o interacción.
- Las etiquetas técnicas pueden usar mayúsculas y espaciado de letras moderado, pero nunca deben dificultar la lectura.
- Evitar títulos completamente en mayúsculas.
- No usar `letter-spacing` negativo de forma global. Ajustarlo únicamente cuando una prueba visual lo justifique.
- Los títulos deben poder ocupar dos líneas en móvil sin solaparse con controles o contenido vecino.

## 4. Espaciado y composición

Usar una escala base de 4 px para mantener ritmo y previsibilidad:

```css
--space-1: 0.25rem;
--space-2: 0.5rem;
--space-3: 0.75rem;
--space-4: 1rem;
--space-5: 1.5rem;
--space-6: 2rem;
--space-7: 3rem;
--space-8: 4rem;
--space-9: 6rem;
```

- Separar etiqueta y control con `--space-2` o `--space-3`.
- Separar grupos relacionados con `--space-4` o `--space-5`.
- Reservar `--space-7` a `--space-9` para cambios de sección y composición de portada.
- Mantener una anchura de lectura aproximada de 60 a 75 caracteres en párrafos.
- El contenido principal debe tener una anchura máxima estable y márgenes fluidos.
- El espacio debe reducirse progresivamente en móvil, sin eliminar la separación entre grupos.

## 5. Forma, bordes y profundidad

```css
--radius-sm: 0.375rem;
--radius-md: 0.5rem;
--radius-lg: 0.75rem;
--border-subtle: 1px solid color-mix(in srgb, var(--color-surface) 16%, transparent);
--shadow-soft: 0 1rem 3rem rgb(6 26 36 / 18%);
```

- Preferir radios pequeños y consistentes: de `6px` a `12px`.
- Usar bordes sutiles para separar superficies; no rodear cada elemento de la interfaz.
- Las tarjetas deben representar unidades de información reales, no convertirse en contenedores decorativos anidados.
- Las sombras deben ser suaves y reservadas para elementos elevados, resultados o diálogos.
- Evitar gradientes intensos, efectos de brillo excesivos y superficies visualmente ruidosas.

## 6. Componentes

### Navegación y shell

- La navegación lateral debe distinguir claramente la ruta activa mediante superficie, borde o color primario.
- La marca debe utilizar Cormorant Garamond cuando funcione como expresión editorial y Montserrat cuando funcione como etiqueta operativa.
- El estado de conexión de la API debe incluir texto y señal visual; el color por sí solo no es suficiente.
- En pantallas pequeñas, la navegación puede compactarse, pero debe conservar nombres accesibles y estados perceptibles.

### Botones

- Cada vista debe tener un único botón primario por flujo principal.
- El botón primario usa `--color-primary` y texto con contraste suficiente. Nunca usar `--color-success` como fondo de un botón principal.
- Los botones secundarios deben tener menor peso visual mediante borde o superficie, no mediante texto ilegible.
- Los botones deben mantener una altura mínima aproximada de `44px` para interacción táctil.
- Estados requeridos: reposo, hover, focus-visible, activo, deshabilitado y carga.
- Durante la carga, conservar el ancho del botón para evitar saltos de layout.

### Formularios

- Cada campo debe tener una etiqueta visible y asociada mediante `for`/`id`.
- El valor monetario debe destacar sin ocultar moneda, formato ni errores.
- Los campos seleccionables deben comunicar selección con más de una señal: color, borde, icono o texto.
- Los mensajes de error deben aparecer junto al campo o acción que los provoca y explicar cómo corregirlo.
- El foco debe ser visible sobre cualquier superficie.

### Tarjetas y resultados

- Usar tarjetas para grupos de información relacionados, como metodologías y distribución calculada.
- El resultado financiero debe priorizar el importe disponible, el método elegido y después el desglose.
- Las métricas deben alinearse de forma consistente y conservar sus etiquetas.
- Los estados vacíos deben explicar qué falta y ofrecer una acción clara para comenzar.

### Iconografía

- Los iconos deben reforzar el significado del texto, nunca reemplazarlo en acciones críticas.
- Mantener un estilo de trazo y una escala coherentes.
- Los iconos interactivos necesitan nombre accesible o `aria-label`.
- No usar símbolos ambiguos si una etiqueta breve puede evitar confusión.

## 7. Estados y accesibilidad

- El contraste mínimo objetivo es WCAG AA: `4.5:1` para texto normal y `3:1` para texto grande o elementos gráficos relevantes.
- Todo control debe ser usable con teclado y mostrar `:focus-visible`.
- No comunicar éxito, error o selección únicamente mediante color.
- Los errores deben conservarse hasta que el usuario pueda corregir la causa o repetir la acción.
- El estado de carga debe indicar que la acción continúa y bloquear dobles envíos cuando corresponda.
- Las transiciones deben ser breves y discretas. Respetar `prefers-reduced-motion`.
- El contenido debe permanecer legible con zoom del navegador y tamaños de fuente aumentados.

## 8. Responsive

### Escritorio: desde 1024 px

- Mantener navegación lateral completa y un área de contenido centrada.
- Usar grids de dos o tres columnas únicamente cuando cada bloque conserve una anchura cómoda.
- La portada puede combinar texto y visualización, dejando una acción principal claramente dominante.

### Tableta: entre 721 px y 1023 px

- Reducir padding y separación antes de eliminar información.
- Permitir que grids de tres columnas pasen a dos.
- Mantener visible el texto de navegación o proporcionar un mecanismo equivalente accesible.

### Móvil: hasta 720 px

- Priorizar una sola columna para formularios, tarjetas y resultados.
- Reducir títulos de manera fluida, sin usar tamaños que impidan la lectura.
- La navegación puede compactarse a iconos, pero debe conservar tooltips, nombres accesibles o una alternativa equivalente.
- Los botones principales deben ocupar el ancho disponible cuando el flujo lo necesite.
- Evitar que importes, etiquetas, errores o controles se desborden horizontalmente.
- Revisar cada pantalla en orientación vertical y con contenido real, no solo con estados vacíos.

## 9. Tokens CSS de referencia

La implementación debe centralizar los tokens en `frontend/src/app/styles.css` o en un módulo de estilos compartido:

```css
:root {
  --color-ink: ...;
  --color-ocean: ...;
  --color-primary: ...;
  --color-sky: ...;
  --color-surface: ...;
  --color-success: ...;

  --font-display: 'Cormorant Garamond', serif;
  --font-body: 'Montserrat', sans-serif;

  --space-1: 0.25rem;
  --space-2: 0.5rem;
  --space-3: 0.75rem;
  --space-4: 1rem;
  --space-5: 1.5rem;
  --space-6: 2rem;
  --space-7: 3rem;
  --space-8: 4rem;
  --space-9: 6rem;
}
```

Los valores reales de color deben declararse una sola vez en la implementación y consumirse mediante variables semánticas.

## 10. Ruta de migración del frontend

La migración visual se realizará sin alterar la lógica de cálculo ni los contratos de la API.

1. Actualizar la importación de fuentes y sustituir `Manrope`/`DM Mono` por Cormorant Garamond/Montserrat en `frontend/src/app/styles.css`.
2. Reemplazar los tokens actuales (`--bg`, `--panel`, `--muted`, `--accent` y colores de categorías) por los tokens semánticos de esta guía.
3. Revisar `h1`, `h2`, `h3`, etiquetas, navegación y métricas para aplicar la jerarquía tipográfica definida.
4. Adaptar `.app-shell`, `.sidebar`, `.topbar`, `.primary-button`, `.method-card` y `.result-card` a la nueva composición y profundidad visual.
5. Revisar los colores de categorías en `frontend/src/pages/budget/BudgetPage.tsx`. La paleta base no debe ampliarse automáticamente; cada categoría debe conservar también una etiqueta textual y un patrón accesible.
6. Alinear el `theme-color` de `frontend/index.html` con el token de fondo elegido para la aplicación.
7. Verificar `AppShell`, `DashboardPage` y `BudgetPage` en estados de reposo, selección, carga, error y resultado.
8. Ejecutar el build de Vite y realizar una revisión visual en escritorio y móvil antes de considerar terminada la migración.

### Fuera de alcance

- Cambios en la lógica de negocio de Go.
- Cambios en los endpoints o modelos de la API.
- Cambios en la navegación funcional de React.
- Incorporación de una librería de componentes sin una decisión posterior del proyecto.

## 11. Criterio de aceptación

Una pantalla cumple este sistema cuando:

- Usa únicamente tokens documentados para color, tipografía y espaciado.
- Tiene una jerarquía clara entre título, contexto, acción y resultado.
- Comunica estados sin depender únicamente del color.
- Conserva legibilidad, foco y acciones utilizables en móvil.
- Mantiene la sensación de calma y confianza sin sacrificar información financiera.
- Puede revisarse contra esta guía sin depender de decisiones implícitas en una pantalla anterior.
