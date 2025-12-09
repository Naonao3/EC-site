'use client'

import React, { useEffect } from 'react'
import Link from 'next/link'
import { useAuthStore } from '@/stores/authStore'
import { useCartStore } from '@/stores/cartStore'

export const Header: React.FC = () => {
  const { isAuthenticated, user, logout } = useAuthStore()
  const { itemCount, fetchCart } = useCartStore()

  useEffect(() => {
    if (isAuthenticated) {
      fetchCart().catch(console.error)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated])

  return (
    <header className="bg-white shadow-md">
      <nav className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* ロゴ */}
          <Link href="/" className="flex items-center space-x-2">
            <div className="text-2xl font-bold text-blue-600">EC Shop</div>
          </Link>

          {/* ナビゲーションリンク */}
          <div className="flex items-center space-x-6">
            <Link
              href="/products"
              className="text-gray-700 hover:text-blue-600 transition-colors"
            >
              商品一覧
            </Link>

            {/* カートアイコン */}
            <Link
              href="/cart"
              className="relative text-gray-700 hover:text-blue-600 transition-colors"
            >
              <svg
                className="h-6 w-6"
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
              {itemCount > 0 && (
                <span className="absolute -top-2 -right-2 bg-red-600 text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center">
                  {itemCount}
                </span>
              )}
            </Link>

            {/* 認証関連 */}
            {isAuthenticated ? (
              <div className="flex items-center space-x-4">
                <Link
                  href="/orders"
                  className="text-gray-700 hover:text-blue-600 transition-colors"
                >
                  注文履歴
                </Link>
                <div className="flex items-center space-x-2">
                  <span className="text-sm text-gray-600">{user?.name}</span>
                  <button
                    onClick={logout}
                    className="text-sm text-gray-700 hover:text-red-600 transition-colors"
                  >
                    ログアウト
                  </button>
                </div>
              </div>
            ) : (
              <div className="flex items-center space-x-4">
                <Link
                  href="/login"
                  className="text-gray-700 hover:text-blue-600 transition-colors"
                >
                  ログイン
                </Link>
                <Link
                  href="/register"
                  className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
                >
                  新規登録
                </Link>
              </div>
            )}
          </div>
        </div>
      </nav>
    </header>
  )
}
