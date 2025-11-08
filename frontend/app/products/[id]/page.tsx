'use client'

import React, { useEffect, useState } from 'react'
import { useParams, useRouter } from 'next/navigation'
import Image from 'next/image'
import { Layout } from '@/components/layout'
import { Button, Loading, Error } from '@/components/ui'
import { apiClient } from '@/lib/api-client'
import { useCartStore } from '@/stores/cartStore'
import { useAuthStore } from '@/stores/authStore'
import type { Product } from '@/types'

export default function ProductDetailPage() {
  const params = useParams()
  const router = useRouter()
  const productId = Number(params.id)

  const [product, setProduct] = useState<Product | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [quantity, setQuantity] = useState(1)
  const [isAddingToCart, setIsAddingToCart] = useState(false)

  const { addItem } = useCartStore()
  const { isAuthenticated } = useAuthStore()

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        setIsLoading(true)
        const data = await apiClient.getProduct(productId)
        setProduct(data)
        setError(null)
      } catch (err) {
        setError('商品の取得に失敗しました')
        console.error(err)
      } finally {
        setIsLoading(false)
      }
    }

    fetchProduct()
  }, [productId])

  const handleAddToCart = async () => {
    if (!isAuthenticated) {
      router.push('/login')
      return
    }

    try {
      setIsAddingToCart(true)
      await addItem(productId, quantity)
      alert('カートに追加しました')
    } catch (err) {
      alert('カートへの追加に失敗しました')
      console.error(err)
    } finally {
      setIsAddingToCart(false)
    }
  }

  if (isLoading) {
    return (
      <Layout>
        <Loading fullScreen text="商品を読み込み中..." />
      </Layout>
    )
  }

  if (error || !product) {
    return (
      <Layout>
        <Error message={error || '商品が見つかりません'} />
      </Layout>
    )
  }

  const stockQuantity = product.inventory?.stock_quantity || 0
  const isInStock = stockQuantity > 0

  return (
    <Layout>
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
        {/* 商品画像 */}
        <div className="aspect-square relative bg-gray-100 rounded-lg overflow-hidden">
          {product.image_url ? (
            <Image
              src={product.image_url}
              alt={product.name}
              fill
              className="object-cover"
              priority
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-400 text-xl">
              No Image
            </div>
          )}
        </div>

        {/* 商品詳細 */}
        <div>
          <h1 className="text-4xl font-bold text-gray-800 mb-4">
            {product.name}
          </h1>

          <div className="text-4xl font-bold text-blue-600 mb-6">
            ¥{product.price.toLocaleString()}
          </div>

          <div className="mb-6">
            {isInStock ? (
              <span className="inline-block bg-green-100 text-green-800 px-4 py-2 rounded-full font-medium">
                在庫あり（{stockQuantity}個）
              </span>
            ) : (
              <span className="inline-block bg-red-100 text-red-800 px-4 py-2 rounded-full font-medium">
                在庫なし
              </span>
            )}
          </div>

          {product.description && (
            <div className="mb-8">
              <h2 className="text-xl font-semibold mb-3">商品説明</h2>
              <p className="text-gray-700 leading-relaxed whitespace-pre-line">
                {product.description}
              </p>
            </div>
          )}

          {/* カート追加 */}
          {isInStock && (
            <div className="border-t pt-6">
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  数量
                </label>
                <select
                  value={quantity}
                  onChange={(e) => setQuantity(Number(e.target.value))}
                  className="w-32 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  {Array.from({ length: Math.min(stockQuantity, 10) }, (_, i) => i + 1).map(
                    (num) => (
                      <option key={num} value={num}>
                        {num}
                      </option>
                    )
                  )}
                </select>
              </div>

              <Button
                onClick={handleAddToCart}
                isLoading={isAddingToCart}
                size="lg"
                className="w-full"
              >
                カートに追加
              </Button>
            </div>
          )}

          <div className="mt-6">
            <Button
              variant="outline"
              onClick={() => router.back()}
              className="w-full"
            >
              商品一覧に戻る
            </Button>
          </div>
        </div>
      </div>
    </Layout>
  )
}
