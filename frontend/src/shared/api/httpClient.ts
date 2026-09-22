import { firebaseAuth } from '../../features/auth/firebase'

const API_URL = import.meta.env.VITE_API_URL ?? ''

async function request<TResponse>(path: string, init: RequestInit): Promise<TResponse> {
  const headers = new Headers(init.headers)
  headers.set('Content-Type', 'application/json')
  if (firebaseAuth.currentUser) headers.set('Authorization', `Bearer ${await firebaseAuth.currentUser.getIdToken()}`)

  const response = await fetch(`${API_URL}${path}`, { ...init, headers })

  if (!response.ok) {
    if (response.status === 401 && firebaseAuth.currentUser) await firebaseAuth.signOut()
    const message = await response.text()
    throw new Error(message || 'No se pudo completar la solicitud')
  }

  return response.json() as Promise<TResponse>
}

export function get<TResponse>(path: string): Promise<TResponse> {
  return request<TResponse>(path, { method: 'GET' })
}

export function post<TRequest, TResponse>(path: string, body: TRequest): Promise<TResponse> {
  return request<TResponse>(path, {
    method: 'POST',
    body: JSON.stringify(body),
  })
}
