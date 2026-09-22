import { createContext, useContext, useEffect, useState, type ReactNode } from 'react'
import { createUserWithEmailAndPassword, onAuthStateChanged, signInWithEmailAndPassword, signInWithPopup, signOut, updateProfile, type User } from 'firebase/auth'
import { get } from '../../shared/api/httpClient'
import { firebaseAuth, googleProvider } from './firebase'
import type { AuthContextValue, Profile } from './auth.types'

const AuthContext = createContext<AuthContextValue | null>(null)

function readableError(error: unknown): string {
  const code = error instanceof Error ? (error as { code?: string }).code : ''
  const messages: Record<string, string> = {
    'auth/invalid-credential': 'El correo o la contraseña no son correctos.',
    'auth/email-already-in-use': 'Este correo ya está registrado.',
    'auth/weak-password': 'La contraseña debe tener al menos 6 caracteres.',
    'auth/popup-closed-by-user': 'El acceso con Google fue cancelado.',
  }
  return messages[code ?? ''] ?? 'No se pudo completar la autenticación.'
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [profile, setProfile] = useState<Profile | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    let active = true
    const unsubscribe = onAuthStateChanged(firebaseAuth, async (nextUser) => {
      if (!active) return
      setUser(nextUser)
      if (!nextUser) {
        setProfile(null)
        setError('')
        setLoading(false)
        return
      }

      setLoading(true)
      try {
        const nextProfile = await get<Profile>('/api/v1/auth/me')
        if (!active) return
        setProfile(nextProfile)
        setError('')
      } catch {
        if (!active) return
        setProfile(null)
        setError('No se pudo sincronizar tu perfil con Splitz.')
        await signOut(firebaseAuth)
      } finally {
        if (active) setLoading(false)
      }
    })

    return () => {
      active = false
      unsubscribe()
    }
  }, [])

  async function run(action: () => Promise<void>) {
    setError('')
    try {
      await action()
    } catch (actionError) {
      setError(readableError(actionError))
      throw actionError
    }
  }

  const value: AuthContextValue = {
    user,
    profile,
    loading,
    error,
    signIn: (email, password) => run(async () => { await signInWithEmailAndPassword(firebaseAuth, email, password) }),
    signUp: (name, email, password) => run(async () => {
      const result = await createUserWithEmailAndPassword(firebaseAuth, email, password)
      await updateProfile(result.user, { displayName: name })
      await result.user.getIdToken(true)
      setProfile(await get<Profile>('/api/v1/auth/me'))
    }),
    signInWithGoogle: () => run(async () => { await signInWithPopup(firebaseAuth, googleProvider) }),
    signOutUser: () => run(async () => { await signOut(firebaseAuth) }),
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (!context) throw new Error('useAuth debe utilizarse dentro de AuthProvider')
  return context
}