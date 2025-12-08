'use client'

import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { Layout } from '@/components/layout'
import { Loading, Error, Card } from '@/components/ui'
import { useAuthStore } from '@/stores/authStore'
import { apiClient } from '@/lib/api-client'
import type { Order } from '@/types'

export default function OrdersPage() {
  const router = useRouter()
  const { isAuthenticated } = useAuthStore()

  const [orders, setOrders] = useState<Order[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login')
      return
    }

    const fetchOrders = async () => {
      try {
        setIsLoading(true)
        const data = await apiClient.getOrders()
        setOrders(data)
        setError(null)
      } catch (err) {
        setError('注文履歴の取得に失敗しました')
        console.error(err)
      } finally {
        setIsLoading(false)
      }
    }

    fetchOrders()
  }, [isAuthenticated, router])

  if (!isAuthenticated) {
    return null
  }

  if (isLoading) {
    return (
      <Layout>
        <Loading fullScreen text="注文履歴を読み込み中..." />
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

  const getStatusBadge = (status: string) => {
    const statusConfig: Record<string, { label: string; className: string }> = {
      pending: { label: '処理中', className: 'bg-yellow-100 text-yellow-800' },
      confirmed: { label: '確定', className: 'bg-green-100 text-green-800' },
      shipped: { label: '発送済み', className: 'bg-blue-100 text-blue-800' },
      delivered: { label: '配達完了', className: 'bg-gray-100 text-gray-800' },
      cancelled: { label: 'キャンセル', className: 'bg-red-100 text-red-800' },
    }

    const config = statusConfig[status] || { label: status, className: 'bg-gray-100 text-gray-800' }

    return (
      <span className={`px-3 py-1 rounded-full text-sm font-medium ${config.className}`}>
        {config.label}
      </span>
    )
  }

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('ja-JP', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    })
  }

  return (
    <Layout>
      <div className="max-w-4xl mx-auto">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-800 mb-2">注文履歴</h1>
          <p className="text-gray-600">
            {orders.length > 0 ? `${orders.length}件の注文` : '注文履歴がありません'}
          </p>
        </div>

        {orders.length === 0 ? (
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
                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
              />
            </svg>
            <p className="text-gray-600 text-lg mb-6">まだ注文がありません</p>
            <Link
              href="/products"
              className="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors"
            >
              商品を見る
            </Link>
          </div>
        ) : (
          <div className="space-y-4">
            {orders.map((order) => (
              <Card key={order.id} className="hover:shadow-lg transition-shadow">
                <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-4">
                  <div>
                    <div className="flex items-center gap-3 mb-2">
                      <h3 className="font-semibold text-lg">
                        注文番号: #{order.id}
                      </h3>
                      {getStatusBadge(order.status)}
                    </div>
                    <p className="text-sm text-gray-600">
                      {formatDate(order.created_at)}
                    </p>
                  </div>
                  <div className="mt-4 sm:mt-0 text-right">
                    <div className="text-sm text-gray-600 mb-1">合計金額</div>
                    <div className="text-2xl font-bold text-blue-600">
                      ¥{order.total_amount.toLocaleString()}
                    </div>
                  </div>
                </div>

                {/* 注文商品 */}
                {order.order_items && order.order_items.length > 0 && (
                  <div className="border-t pt-4">
                    <h4 className="text-sm font-medium text-gray-700 mb-3">注文内容</h4>
                    <div className="space-y-2">
                      {order.order_items.map((item) => (
                        <div
                          key={item.id}
                          className="flex justify-between items-center text-sm"
                        >
                          <span className="text-gray-700">
                            {item.product?.name || '商品名不明'} × {item.quantity}
                          </span>
                          <span className="text-gray-600">
                            ¥{(item.price * item.quantity).toLocaleString()}
                          </span>
                        </div>
                      ))}
                    </div>
                  </div>
                )}

                {/* 決済情報 */}
                {order.payment && (
                  <div className="border-t mt-4 pt-4">
                    <div className="flex items-center justify-between text-sm">
                      <span className="text-gray-600">決済状況</span>
                      <span className="font-medium">
                        {order.payment.status === 'succeeded'
                          ? '決済完了'
                          : order.payment.status === 'pending'
                          ? '決済待ち'
                          : '決済失敗'}
                      </span>
                    </div>
                  </div>
                )}
              </Card>
            ))}
          </div>
        )}
      </div>
    </Layout>
  )
}
