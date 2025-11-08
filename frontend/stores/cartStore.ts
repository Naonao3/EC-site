import { create } from 'zustand'
import type { CartItem, Product } from '@/types'
import { apiClient } from '@/lib/api-client'

interface CartState {
  items: CartItem[]
  isLoading: boolean
  totalAmount: number
  itemCount: number
  fetchCart: () => Promise<void>
  addItem: (product_id: number, quantity: number) => Promise<void>
  updateItemQuantity: (id: number, quantity: number) => Promise<void>
  removeItem: (id: number) => Promise<void>
  clearCart: () => Promise<void>
  calculateTotals: () => void
}

export const useCartStore = create<CartState>((set, get) => ({
  items: [],
  isLoading: false,
  totalAmount: 0,
  itemCount: 0,

  fetchCart: async () => {
    set({ isLoading: true })
    try {
      const items = await apiClient.getCartItems()
      set({ items, isLoading: false })
      get().calculateTotals()
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  addItem: async (product_id, quantity) => {
    set({ isLoading: true })
    try {
      await apiClient.addToCart(product_id, quantity)
      await get().fetchCart()
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  updateItemQuantity: async (id, quantity) => {
    set({ isLoading: true })
    try {
      await apiClient.updateCartItem(id, quantity)
      await get().fetchCart()
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  removeItem: async (id) => {
    set({ isLoading: true })
    try {
      await apiClient.removeFromCart(id)
      await get().fetchCart()
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  clearCart: async () => {
    set({ isLoading: true })
    try {
      await apiClient.clearCart()
      set({ items: [], isLoading: false })
      get().calculateTotals()
    } catch (error) {
      set({ isLoading: false })
      throw error
    }
  },

  calculateTotals: () => {
    const { items } = get()
    const totalAmount = items.reduce((sum, item) => {
      return sum + (item.product?.price || 0) * item.quantity
    }, 0)
    const itemCount = items.reduce((sum, item) => sum + item.quantity, 0)
    set({ totalAmount, itemCount })
  },
}))
