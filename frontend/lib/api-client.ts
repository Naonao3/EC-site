import axios, { AxiosInstance, AxiosError } from 'axios'
import type {
  User,
  Product,
  CartItem,
  Order,
  Payment,
  ApiResponse,
  PaginatedResponse,
  LoginRequest,
  RegisterRequest,
  AuthResponse,
  CreatePaymentIntentRequest,
  CreatePaymentIntentResponse,
} from '@/types'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

class ApiClient {
  private client: AxiosInstance

  constructor() {
    this.client = axios.create({
      baseURL: `${API_URL}/api`,
      headers: {
        'Content-Type': 'application/json',
      },
    })

    // リクエストインターセプター（JWTトークン付与）
    this.client.interceptors.request.use(
      (config) => {
        const token = this.getToken()
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      }
    )

    // レスポンスインターセプター（エラーハンドリング）
    this.client.interceptors.response.use(
      (response) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401) {
          // 認証エラー時はトークン削除
          this.removeToken()
          if (typeof window !== 'undefined') {
            window.location.href = '/login'
          }
        }
        return Promise.reject(error)
      }
    )
  }

  // トークン管理
  private getToken(): string | null {
    if (typeof window === 'undefined') return null
    return localStorage.getItem('auth_token')
  }

  private setToken(token: string): void {
    if (typeof window === 'undefined') return
    localStorage.setItem('auth_token', token)
  }

  private removeToken(): void {
    if (typeof window === 'undefined') return
    localStorage.removeItem('auth_token')
  }

  // ========================================
  // 認証API
  // ========================================

  async login(data: LoginRequest): Promise<AuthResponse> {
    const response = await this.client.post<AuthResponse>('/auth/login', data)
    this.setToken(response.data.token)
    return response.data
  }

  async register(data: RegisterRequest): Promise<AuthResponse> {
    const response = await this.client.post<AuthResponse>('/auth/register', data)
    this.setToken(response.data.token)
    return response.data
  }

  logout(): void {
    this.removeToken()
  }

  async getCurrentUser(): Promise<User> {
    const response = await this.client.get<{ user: User }>('/auth/me')
    return response.data.user
  }

  // ========================================
  // 商品API
  // ========================================

  async getProducts(params?: {
    page?: number
    per_page?: number
    category_id?: number
  }): Promise<any> {
    const response = await this.client.get('/products', { params })
    return response.data
  }

  async getProduct(id: number): Promise<Product> {
    const response = await this.client.get<{ product: Product }>(`/products/${id}`)
    return response.data.product
  }

  // ========================================
  // カートAPI
  // ========================================

  async getCartItems(): Promise<CartItem[]> {
    const response = await this.client.get('/cart')
    // APIレスポンスから正しくitemsを取得
    return response.data.cart?.items || response.data.items || []
  }

  async addToCart(product_id: number, quantity: number): Promise<CartItem> {
    const response = await this.client.post<{ item: CartItem }>('/cart', {
      product_id,
      quantity,
    })
    return response.data.item
  }

  async updateCartItem(id: number, quantity: number): Promise<CartItem> {
    const response = await this.client.put<{ item: CartItem }>(`/cart/${id}`, {
      quantity,
    })
    return response.data.item
  }

  async removeFromCart(id: number): Promise<void> {
    await this.client.delete(`/cart/${id}`)
  }

  async clearCart(): Promise<void> {
    await this.client.delete('/cart')
  }

  // ========================================
  // 注文API
  // ========================================

  async createOrder(): Promise<Order> {
    const response = await this.client.post<{ order: Order }>('/orders')
    return response.data.order
  }

  async getOrders(): Promise<Order[]> {
    const response = await this.client.get<{ orders: Order[] }>('/orders')
    return response.data.orders
  }

  async getOrder(id: number): Promise<Order> {
    const response = await this.client.get<{ order: Order }>(`/orders/${id}`)
    return response.data.order
  }

  // ========================================
  // 決済API（Stripe）
  // ========================================

  async createPaymentIntent(
    data: CreatePaymentIntentRequest
  ): Promise<CreatePaymentIntentResponse> {
    const response = await this.client.post<CreatePaymentIntentResponse>(
      '/payment/create-intent',
      data
    )
    return response.data
  }

  async getPaymentByOrderId(order_id: number): Promise<Payment> {
    const response = await this.client.get<{ payment: Payment }>(
      `/payment/order/${order_id}`
    )
    return response.data.payment
  }
}

export const apiClient = new ApiClient()
