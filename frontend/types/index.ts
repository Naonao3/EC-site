// ========================================
// 型定義
// ========================================

export interface User {
  id: number
  email: string
  name: string
  created_at: string
  updated_at: string
}

export interface Category {
  id: number
  name: string
  slug: string
  parent_id?: number
  created_at: string
}

export interface Product {
  id: number
  name: string
  description?: string
  price: number
  stock: number  // バックエンドから直接返される在庫数
  category?: string
  image_url?: string
  created_at: string
  updated_at: string
}

export interface Inventory {
  id: number
  product_id: number
  stock_quantity: number
  reserved_quantity: number
  updated_at: string
}

export interface CartItem {
  id: number
  user_id: number
  product_id: number
  quantity: number
  created_at: string
  product?: Product
}

export interface Order {
  id: number
  user_id: number
  order_number: string
  total_amount: number
  status: OrderStatus
  created_at: string
  updated_at: string
  items?: OrderItem[]
  payment?: Payment
}

export type OrderStatus = 'pending' | 'confirmed' | 'shipped' | 'delivered' | 'cancelled'

export interface OrderItem {
  id: number
  order_id: number
  product_id: number
  quantity: number
  unit_price: number
  product?: Product
}

export interface Payment {
  id: number
  order_id: number
  stripe_payment_intent_id?: string
  stripe_payment_method_id?: string
  amount: number
  currency: string
  status: PaymentStatus
  created_at: string
  updated_at: string
  order?: Order
}

export type PaymentStatus = 'pending' | 'succeeded' | 'failed' | 'cancelled'

export interface Review {
  id: number
  user_id: number
  product_id: number
  rating: number
  comment?: string
  created_at: string
  user?: User
  product?: Product
}

// API レスポンス型
export interface ApiResponse<T> {
  data?: T
  error?: string
  message?: string
}

// ページネーション
export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  per_page: number
  total_pages: number
}

// 認証関連
export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  name: string
}

export interface AuthResponse {
  token: string
  user: User
}

// Stripe関連
export interface CreatePaymentIntentRequest {
  order_id: number
}

export interface CreatePaymentIntentResponse {
  client_secret: string
}
