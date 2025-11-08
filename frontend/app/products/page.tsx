'use client'

import React, { useEffect, useState } from 'react'
import { Layout } from '@/components/layout'
import { Loading, Error } from '@/components/ui'
import { ProductCard } from '@/components/features/ProductCard'
import { apiClient } from '@/lib/api-client'
import type { Product } from '@/types'

export default function ProductsPage() {
  const [products, setProducts] = useState<Product[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        setIsLoading(true)
        const response = await apiClient.getProducts()
        // APIレスポンスからproductsを取得
        setProducts(response.products || response.data || [])
        setError(null)
      } catch (err) {
        setError('商品の取得に失敗しました')
        console.error(err)
      } finally {
        setIsLoading(false)
      }
    }

    fetchProducts()
  }, [])

  if (isLoading) {
    return (
      <Layout>
        <Loading fullScreen text="商品を読み込み中..." />
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
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-800 mb-2">商品一覧</h1>
        <p className="text-gray-600">
          全{products.length}件の商品
        </p>
      </div>

      {products.length === 0 ? (
        <div className="text-center py-16">
          <p className="text-gray-600 text-lg">商品がありません</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {products.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>
      )}
    </Layout>
  )
}
