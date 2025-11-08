'use client'

import React from 'react'
import Link from 'next/link'
import { Layout } from '@/components/layout'
import { Button } from '@/components/ui'

export default function CheckoutCancelPage() {
  return (
    <Layout>
      <div className="max-w-2xl mx-auto text-center py-16">
        {/* キャンセルアイコン */}
        <div className="mb-8">
          <div className="mx-auto w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center">
            <svg
              className="h-16 w-16 text-gray-600"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </div>
        </div>

        {/* メッセージ */}
        <h1 className="text-4xl font-bold text-gray-800 mb-4">
          決済がキャンセルされました
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          決済処理を中断しました
        </p>

        <div className="space-y-4">
          <p className="text-gray-700">
            カート内の商品はそのまま保存されています。
            <br />
            準備ができたら再度チェックアウトをお試しください。
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center mt-8">
            <Link href="/cart">
              <Button size="lg">カートに戻る</Button>
            </Link>
            <Link href="/products">
              <Button variant="outline" size="lg">
                買い物を続ける
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </Layout>
  )
}
