'use client'

import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { Layout } from '@/components/layout'
import { Button, Loading, Error } from '@/components/ui'
import { CartItem } from '@/components/features/CartItem'
import { useCartStore } from '@/stores/cartStore'
import { useAuthStore } from '@/stores/authStore'

export default function CartPage() {
  const router = useRouter()
  const { isAuthenticated } = useAuthStore()
  const {
    items = [],  // デフォルト値を設定
    totalAmount = 0,
    itemCount = 0,
    isLoading,
    fetchCart,
    updateItemQuantity,
    removeItem,
    clearCart,
  } = useCartStore()

  const [isUpdating, setIsUpdating] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login')
      return
    }

    fetchCart().catch((err) => {
      setError('カート情報の取得に失敗しました')
      console.error(err)
    })
  }, [isAuthenticated, fetchCart, router])

  const handleUpdateQuantity = async (id: number, quantity: number) => {
    try {
      setIsUpdating(true)
      await updateItemQuantity(id, quantity)
    } catch (err) {
      alert('数量の更新に失敗しました')
      console.error(err)
    } finally {
      setIsUpdating(false)
    }
  }

  const handleRemove = async (id: number) => {
    if (!confirm('この商品をカートから削除しますか？')) return

    try {
      setIsUpdating(true)
      await removeItem(id)
    } catch (err) {
      alert('削除に失敗しました')
      console.error(err)
    } finally {
      setIsUpdating(false)
    }
  }

  const handleClearCart = async () => {
    if (!confirm('カート内の全ての商品を削除しますか？')) return

    try {
      setIsUpdating(true)
      await clearCart()
    } catch (err) {
      alert('カートのクリアに失敗しました')
      console.error(err)
    } finally {
      setIsUpdating(false)
    }
  }

  const handleCheckout = () => {
    router.push('/checkout')
  }

  if (!isAuthenticated) {
    return null
  }

  if (isLoading) {
    return (
      <Layout>
        <Loading fullScreen text="カート情報を読み込み中..." />
      </Layout>
    )
  }

  if (error) {
    return (
      <Layout>
        <Error message={error} onRetry={() => window.location.reload()} />
      </Layout>
    )
  }

  return (
    <Layout>
      <div className="max-w-4xl mx-auto">
        <div className="mb-8 flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold text-gray-800 mb-2">
              ショッピングカート
            </h1>
            <p className="text-gray-600">
              {itemCount}個の商品
            </p>
          </div>

          {items.length > 0 && (
            <Button
              variant="outline"
              size="sm"
              onClick={handleClearCart}
              disabled={isUpdating}
            >
              カートを空にする
            </Button>
          )}
        </div>

        {items.length === 0 ? (
          <div className="text-center py-16">
            <svg
              className="mx-auto h-24 w-24 text-gray-400 mb-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={1}
                d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
              />
            </svg>
            <p className="text-gray-600 text-lg mb-6">カートは空です</p>
            <Button onClick={() => router.push('/products')}>
              商品一覧へ
            </Button>
          </div>
        ) : (
          <>
            {/* カート商品リスト */}
            <div className="space-y-4 mb-8">
              {items.map((item) => (
                <CartItem
                  key={item.id}
                  item={item}
                  onUpdateQuantity={handleUpdateQuantity}
                  onRemove={handleRemove}
                  isUpdating={isUpdating}
                />
              ))}
            </div>

            {/* 合計金額とチェックアウト */}
            <div className="border-t pt-6">
              <div className="bg-gray-50 rounded-lg p-6">
                <div className="flex items-center justify-between mb-6">
                  <span className="text-xl font-semibold text-gray-800">
                    合計金額
                  </span>
                  <span className="text-3xl font-bold text-blue-600">
                    ¥{totalAmount.toLocaleString()}
                  </span>
                </div>

                <Button
                  onClick={handleCheckout}
                  size="lg"
                  className="w-full"
                  disabled={isUpdating}
                >
                  チェックアウトに進む
                </Button>

                <div className="mt-4 text-center">
                  <button
                    onClick={() => router.push('/products')}
                    className="text-blue-600 hover:text-blue-700 text-sm"
                  >
                    ← 買い物を続ける
                  </button>
                </div>
              </div>
            </div>
          </>
        )}
      </div>
    </Layout>
  )
}
