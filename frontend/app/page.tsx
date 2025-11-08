'use client'

import Link from 'next/link'
import { Layout } from '@/components/layout'
import { Card } from '@/components/ui'

export default function Home() {
  return (
    <Layout>
      <div className="text-center mb-12">
        <h1 className="text-5xl font-bold text-gray-800 mb-4">
          EC Shop
        </h1>
        <p className="text-xl text-gray-600">
          柴犬画像販売サイト - ポートフォリオプロジェクト
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mt-12">
        <Link href="/products">
          <Card hover>
            <div className="text-center">
              <svg
                className="mx-auto h-12 w-12 text-blue-600 mb-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"
                />
              </svg>
              <h2 className="text-2xl font-semibold mb-2">商品一覧</h2>
              <p className="text-gray-600">全ての商品を見る</p>
            </div>
          </Card>
        </Link>

        <Link href="/cart">
          <Card hover>
            <div className="text-center">
              <svg
                className="mx-auto h-12 w-12 text-blue-600 mb-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                />
              </svg>
              <h2 className="text-2xl font-semibold mb-2">カート</h2>
              <p className="text-gray-600">カート内の商品を確認</p>
            </div>
          </Card>
        </Link>

        <Link href="/orders">
          <Card hover>
            <div className="text-center">
              <svg
                className="mx-auto h-12 w-12 text-blue-600 mb-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                />
              </svg>
              <h2 className="text-2xl font-semibold mb-2">注文履歴</h2>
              <p className="text-gray-600">過去の注文を確認</p>
            </div>
          </Card>
        </Link>
      </div>

      <div className="mt-16 text-center">
        <div className="inline-block bg-blue-50 rounded-lg p-6">
          <h3 className="text-lg font-semibold text-gray-800 mb-2">
            使用技術
          </h3>
          <p className="text-sm text-gray-600">
            TypeScript • Next.js • Go • PostgreSQL • Redis • Stripe
          </p>
        </div>
      </div>
    </Layout>
  )
}
