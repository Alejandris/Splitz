import type { User } from 'firebase/auth'

export type Profile = {
  id: string
  firebase_uid: string
  email: string
  name: string
}

export type AuthContextValue = {
  user: User | null
  profile: Profile | null
  loading: boolean
  error: string
  signIn: (email: string, password: string) => Promise<void>
  signUp: (name: string, email: string, password: string) => Promise<void>
  signInWithGoogle: () => Promise<void>
  signOutUser: () => Promise<void>
}