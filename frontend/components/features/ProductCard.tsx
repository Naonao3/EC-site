import React from 'react'
import Link from 'next/link'
import Image from 'next/image'
import { Card } from '@/components/ui'
import type { Product } from '@/types'

interface ProductCardProps {
  product: Product
}

export const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  const imageUrl = product.image_url || '/placeholder-product.jpg'
  const stockQuantity = product.stock || 0
  const isInStock = stockQuantity > 0

  return (
    <Link href={`/products/${product.id}`}>
      <Card hover className="h-full">
        <div className="relative aspect-square mb-4 bg-gray-100 rounded-lg overflow-hidden">
          {product.image_url ? (
            <Image
              src={imageUrl}
              alt={product.name}
              fill
              className="object-cover"
              sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-400">
              No Image
            </div>
          )}
          {!isInStock && (
            <div className="absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center">
              <span className="text-white font-bold text-lg">在庫なし</span>
            </div>
          )}
        </div>

        <h3 className="font-semibold text-lg mb-2 line-clamp-2">{product.name}</h3>

        {product.description && (
          <p className="text-gray-600 text-sm mb-3 line-clamp-2">
            {product.description}
          </p>
        )}

        <div className="flex items-center justify-between mt-auto">
          <span className="text-2xl font-bold text-blue-600">
            ¥{product.price.toLocaleString()}
          </span>
          {isInStock && (
            <span className="text-sm text-gray-500">
              在庫: {stockQuantity}
            </span>
          )}
        </div>
      </Card>
    </Link>
  )
}
