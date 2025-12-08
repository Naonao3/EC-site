'use client'

import React, { useEffect, useState, Suspense } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import Link from 'next/link'
import { Layout } from '@/components/layout'
import { Button, Loading } from '@/components/ui'
import { useCartStore } from '@/stores/cartStore'
import { useCheckoutStore } from '@/stores/checkoutStore'

function SuccessContent() {
  const router = useRouter()
  const searchParams = useSearchParams()
  const orderId = searchParams.get('order_id')

  const { clearCart } = useCartStore()
  const { resetCheckout } = useCheckoutStore()
  const [isClearing, setIsClearing] = useState(true)

  useEffect(() => {
    // カートとチェックアウト状態をクリア
    const clear = async () => {
      try {
        await clearCart()
        resetCheckout()
      } catch (err) {
        console.error('カートのクリアに失敗:', err)
      } finally {
        setIsClearing(false)
      }
    }

    clear()
  }, [clearCart, resetCheckout])

  if (isClearing) {
    return <Loading fullScreen text="注文を処理中..." />
  }

  return (
    <Layout>
      <div className="max-w-2xl mx-auto text-center py-16">
        {/* 成功アイコン */}
        <div className="mb-8">
          <div className="mx-auto w-24 h-24 bg-green-100 rounded-full flex items-center justify-center">
            <svg
              className="h-16 w-16 text-green-600"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M5 13l4 4L19 7"
              />
            </svg>
          </div>
        </div>

        {/* メッセージ */}
        <h1 className="text-4xl font-bold text-gray-800 mb-4">
          ご注文ありがとうございます！
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          決済が正常に完了しました
        </p>

        {orderId && (
          <div className="bg-blue-50 rounded-lg p-6 mb-8">
            <p className="text-sm text-gray-600 mb-2">注文番号</p>
            <p className="text-2xl font-bold text-blue-600">#{orderId}</p>
          </div>
        )}

        <div className="space-y-4">
          <p className="text-gray-700">
            注文確認メールを送信しました。
            <br />
            注文履歴からも詳細を確認できます。
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center mt-8">
            <Link href="/orders">
              <Button size="lg">注文履歴を見る</Button>
            </Link>
            <Link href="/products">
              <Button variant="outline" size="lg">
                買い物を続ける
              </Button>
            </Link>
          </div>
        </div>

        {/* テストモード注意 */}
        <div className="mt-12 p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
          <p className="text-sm text-yellow-800">
            ⚠️ テストモード：実際の決済は行われていません
          </p>
        </div>
      </div>
    </Layout>
  )
}

export default function CheckoutSuccessPage() {
  return (
    <Suspense fallback={<Loading fullScreen />}>
      <SuccessContent />
    </Suspense>
  )
}
