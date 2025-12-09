'use client'

import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { Elements } from '@stripe/react-stripe-js'
import { Layout } from '@/components/layout'
import { Loading, Error } from '@/components/ui'
import { CheckoutForm } from '@/components/stripe/CheckoutForm'
import { getStripe } from '@/lib/stripe'
import { useCartStore } from '@/stores/cartStore'
import { useCheckoutStore } from '@/stores/checkoutStore'
import { useAuthStore } from '@/stores/authStore'

export default function CheckoutPage() {
  const router = useRouter()
  const { isAuthenticated } = useAuthStore()
  const { items } = useCartStore()
  const { createOrder, createPaymentIntent, clientSecret, currentOrder } = useCheckoutStore()

  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const stripePromise = getStripe()

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login')
      return
    }

    const initializeCheckout = async () => {
      try {
        setIsLoading(true)

        // 既に注文が作成されている場合はスキップ
        if (currentOrder) {
          // Payment Intent作成（既存の場合は既存のclient_secretを返す）
          await createPaymentIntent(currentOrder.id)
          setError(null)
          setIsLoading(false)
          return
        }

        // カートが空の場合はカートページへリダイレクト
        if (items.length === 0) {
          router.push('/cart')
          return
        }

        // 注文作成
        const order = await createOrder()

        // Payment Intent作成
        await createPaymentIntent(order.id)

        setError(null)
      } catch (err) {
        setError('チェックアウトの初期化に失敗しました')
        console.error(err)
      } finally {
        setIsLoading(false)
      }
    }

    initializeCheckout()
  }, [isAuthenticated, items, router, createOrder, createPaymentIntent, currentOrder])

  if (!isAuthenticated) {
    return null
  }

  if (isLoading) {
    return (
      <Layout>
        <Loading fullScreen text="チェックアウトを準備中..." />
      </Layout>
    )
  }

  if (error || !clientSecret || !currentOrder) {
    return (
      <Layout>
        <Error message={error || 'チェックアウト情報の取得に失敗しました'} />
      </Layout>
    )
  }

  const options = {
    clientSecret,
    appearance: {
      theme: 'stripe' as const,
    },
  }

  const orderItems = currentOrder.order_items || []

  return (
    <Layout>
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold text-gray-800 mb-8">チェックアウト</h1>

        {/* 注文サマリー */}
        <div className="bg-gray-50 rounded-lg p-6 mb-8">
          <h2 className="text-lg font-semibold mb-4">注文内容</h2>
          <div className="space-y-2">
            {orderItems.map((item) => (
              <div key={item.id} className="flex justify-between text-sm">
                <span>
                  {item.product?.name} × {item.quantity}
                </span>
                <span>
                  ¥{(item.price * item.quantity).toLocaleString()}
                </span>
              </div>
            ))}
          </div>
          <div className="border-t mt-4 pt-4 flex justify-between items-center">
            <span className="text-lg font-semibold">合計</span>
            <span className="text-2xl font-bold text-blue-600">
              ¥{currentOrder.total_amount.toLocaleString()}
            </span>
          </div>
        </div>

        {/* Stripe決済フォーム */}
        <Elements stripe={stripePromise} options={options}>
          <CheckoutForm orderId={currentOrder.id} />
        </Elements>
      </div>
    </Layout>
  )
}
