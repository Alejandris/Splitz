const API_URL = import.meta.env.VITE_API_URL ?? ''

export async function post<TRequest, TResponse>(path: string, body: TRequest): Promise<TResponse> {
  const response = await fetch(`${API_URL}${path}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })

  if (!response.ok) {
    const message = await response.text()
    throw new Error(message || 'No se pudo completar la solicitud')
  }

  return response.json() as Promise<TResponse>
}
