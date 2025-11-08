'use client'

import { useEffect } from 'react'
import { useAuthStore } from '@/stores/authStore'

export function Providers({ children }: { children: React.ReactNode }) {
  useEffect(() => {
    // ストアのハイドレーション
    useAuthStore.persist.rehydrate()
  }, [])

  return <>{children}</>
}
