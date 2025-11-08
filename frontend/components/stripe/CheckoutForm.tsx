'use client'

import React, { useState } from 'react'
import { useRouter } from 'next/navigation'
import {
  useStripe,
  useElements,
  PaymentElement,
} from '@stripe/react-stripe-js'
import { Button } from '@/components/ui'

interface CheckoutFormProps {
  orderId: number
}

export const CheckoutForm: React.FC<CheckoutFormProps> = ({ orderId }) => {
  const router = useRouter()
  const stripe = useStripe()
  const elements = useElements()

  const [isProcessing, setIsProcessing] = useState(false)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!stripe || !elements) {
      return
    }

    setIsProcessing(true)
    setErrorMessage(null)

    try {
      const { error } = await stripe.confirmPayment({
        elements,
        confirmParams: {
          return_url: `${window.location.origin}/checkout/success?order_id=${orderId}`,
        },
      })

      if (error) {
        setErrorMessage(error.message || '決済処理に失敗しました')
      }
    } catch (err) {
      setErrorMessage('予期しないエラーが発生しました')
      console.error(err)
    } finally {
      setIsProcessing(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="bg-white p-6 rounded-lg border">
        <h2 className="text-lg font-semibold mb-4">カード情報入力</h2>
        <PaymentElement />
      </div>

      {errorMessage && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-4">
          <p className="text-red-800 text-sm">{errorMessage}</p>
        </div>
      )}

      <div className="space-y-3">
        <Button
          type="submit"
          size="lg"
          className="w-full"
          disabled={!stripe || isProcessing}
          isLoading={isProcessing}
        >
          {isProcessing ? '処理中...' : '決済を確定する'}
        </Button>

        <Button
          type="button"
          variant="outline"
          className="w-full"
          onClick={() => router.push('/cart')}
          disabled={isProcessing}
        >
          カートに戻る
        </Button>
      </div>

      <div className="text-center text-sm text-gray-500">
        <p>
          テストモード：テスト用カード番号 4242 4242 4242 4242
        </p>
      </div>
    </form>
  )
}
