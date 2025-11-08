import React from 'react'
import Image from 'next/image'
import { Button } from '@/components/ui'
import type { CartItem as CartItemType } from '@/types'

interface CartItemProps {
  item: CartItemType
  onUpdateQuantity: (id: number, quantity: number) => void
  onRemove: (id: number) => void
  isUpdating: boolean
}

export const CartItem: React.FC<CartItemProps> = ({
  item,
  onUpdateQuantity,
  onRemove,
  isUpdating,
}) => {
  const product = item.product
  if (!product) return null

  const subtotal = product.price * item.quantity

  return (
    <div className="flex gap-4 p-4 border rounded-lg bg-white">
      {/* 商品画像 */}
      <div className="w-24 h-24 flex-shrink-0 bg-gray-100 rounded-lg overflow-hidden relative">
        {product.image_url ? (
          <Image
            src={product.image_url}
            alt={product.name}
            fill
            className="object-cover"
          />
        ) : (
          <div className="flex items-center justify-center h-full text-gray-400 text-xs">
            No Image
          </div>
        )}
      </div>

      {/* 商品情報 */}
      <div className="flex-1 min-w-0">
        <h3 className="font-semibold text-lg mb-1 truncate">{product.name}</h3>
        <p className="text-blue-600 font-bold mb-3">
          ¥{product.price.toLocaleString()}
        </p>

        <div className="flex items-center gap-4">
          {/* 数量変更 */}
          <div className="flex items-center gap-2">
            <label className="text-sm text-gray-600">数量:</label>
            <select
              value={item.quantity}
              onChange={(e) => onUpdateQuantity(item.id, Number(e.target.value))}
              disabled={isUpdating}
              className="px-3 py-1 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-transparent disabled:opacity-50"
            >
              {Array.from({ length: 10 }, (_, i) => i + 1).map((num) => (
                <option key={num} value={num}>
                  {num}
                </option>
              ))}
            </select>
          </div>

          {/* 削除ボタン */}
          <Button
            variant="danger"
            size="sm"
            onClick={() => onRemove(item.id)}
            disabled={isUpdating}
          >
            削除
          </Button>
        </div>
      </div>

      {/* 小計 */}
      <div className="text-right flex-shrink-0">
        <div className="text-sm text-gray-600 mb-1">小計</div>
        <div className="text-xl font-bold text-gray-800">
          ¥{subtotal.toLocaleString()}
        </div>
      </div>
    </div>
  )
}
