import { create } from 'zustand'
import type { Order, Payment } from '@/types'
import { apiClient } from '@/lib/api-client'

interface CheckoutState {
  currentOrder: Order | null
  currentPayment: Payment | null
  isLoading: boolean
  clientSecret: string | null
  createOrder: () => Promise<Order>
  createPaymentIntent: (orderId: number) => Promise<string>
  getPayment: (orderId: number) => Promise<void>
  resetCheckout: () => void
}

export const useCheckoutStore = create<CheckoutState>((set) => ({
  currentOrder: null,
  currentPayment: null,
  isLoading: false,
  clientSecret: null,

  createOrder: async () => {
    set({ isLoading: true })
    try {
      const order = await apiClient.createOrder()
      set({ currentOrder: order, isLoading: false })
      return order
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  createPaymentIntent: async (orderId) => {
    set({ isLoading: true })
    try {
      const response = await apiClient.createPaymentIntent({ order_id: orderId })
      set({ clientSecret: response.client_secret, isLoading: false })
      return response.client_secret
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  getPayment: async (orderId) => {
    set({ isLoading: true })
    try {
      const payment = await apiClient.getPaymentByOrderId(orderId)
      set({ currentPayment: payment, isLoading: false })
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  resetCheckout: () => {
    set({
      currentOrder: null,
      currentPayment: null,
      clientSecret: null,
      isLoading: false,
    })
  },
}))
