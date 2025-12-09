'use client'

import { useEffect, useState } from 'react'
import { useAuthStore } from '@/stores/authStore'

export function Providers({ children }: { children: React.ReactNode }) {
  const [isHydrated, setIsHydrated] = useState(false)

  useEffect(() => {
    // ストアのハイドレーション
    useAuthStore.persist.rehydrate()
    setIsHydrated(true)
  }, [])

  // ハイドレーション完了まで何も表示しない
  if (!isHydrated) {
    return null
  }

  return <>{children}</>
}
