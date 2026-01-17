# EC-site - Shiba Inu Image Store

[日本語](#日本語) | [English](#english)

---

## 日本語

### 概要

柴犬画像販売ECサイトのフルスタックアプリケーションです。モダンなWeb技術を活用し、ユーザー認証、商品管理、ショッピングカート、Stripe決済機能を実装したポートフォリオプロジェクトです。

### 技術スタック

#### フロントエンド
- **フレームワーク:** Next.js 14.2 (React 18.3)
- **言語:** TypeScript 5.3
- **スタイリング:** Tailwind CSS 3.4
- **状態管理:** Zustand 4.5
- **HTTPクライアント:** Axios 1.7
- **決済UI:** Stripe React

#### バックエンド
- **言語:** Go 1.24
- **フレームワーク:** Gin 1.11
- **データベース:** PostgreSQL 15 (GORM)
- **キャッシュ:** Redis 7
- **認証:** JWT
- **決済処理:** Stripe

#### インフラ
- **コンテナ:** Docker / Docker Compose
- **デプロイ:** Render
- **データベースホスティング:** Supabase
- **キャッシュホスティング:** Upstash
- **画像ストレージ:** Cloudinary

### 主な機能

#### ユーザー向け機能
- ユーザー登録・ログイン（JWT認証）
- 商品一覧・詳細表示
- カテゴリ別商品フィルタリング
- 商品検索
- ショッピングカート管理
- Stripe決済による注文処理
- 注文履歴確認

#### 管理者向け機能
- 商品のCRUD操作
- ユーザー管理
- 注文管理・ステータス更新

### プロジェクト構成

```
EC-site/
├── frontend/                # Next.js フロントエンド
│   ├── app/                 # App Router ページ
│   ├── components/          # 再利用可能なコンポーネント
│   ├── stores/              # Zustand 状態管理
│   ├── lib/                 # ユーティリティ
│   └── types/               # TypeScript 型定義
│
├── backend/                 # Go バックエンドAPI
│   ├── cmd/server/          # エントリーポイント
│   ├── internal/
│   │   ├── model/           # データモデル
│   │   ├── repository/      # データアクセス層
│   │   ├── service/         # ビジネスロジック層
│   │   ├── handler/         # HTTPハンドラー
│   │   └── middleware/      # ミドルウェア
│   ├── config/              # 設定管理
│   └── pkg/                 # 共通パッケージ
│
├── database/                # データベーススキーマ・マイグレーション
├── docs/                    # ドキュメント（ER図など）
├── docker-compose.yml       # ローカル開発環境
└── render.yaml              # デプロイ設定
```

### セットアップ

#### 必要条件
- Node.js 18以上
- Go 1.24以上
- Docker / Docker Compose
- Stripeアカウント

#### ローカル開発環境の起動

1. リポジトリをクローン
```bash
git clone https://github.com/Naonao3/EC-site.git
cd EC-site
```

2. Docker Composeでデータベースとキャッシュを起動
```bash
docker-compose up -d
```

3. バックエンドの起動
```bash
cd backend
cp .env.production.example .env
# .envファイルを編集して環境変数を設定
go run cmd/server/main.go
```

4. フロントエンドの起動
```bash
cd frontend
npm install
npm run dev
```

5. ブラウザで http://localhost:3000 にアクセス

### 環境変数

#### バックエンド
```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=ec_site
DB_SSLMODE=disable
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h
STRIPE_SECRET_KEY=sk_test_xxx
STRIPE_WEBHOOK_SECRET=whsec_xxx
CORS_ORIGIN=http://localhost:3000
```

#### フロントエンド
```
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_xxx
```

### API エンドポイント

| メソッド | エンドポイント | 説明 | 認証 |
|---------|---------------|------|------|
| POST | /api/auth/register | ユーザー登録 | - |
| POST | /api/auth/login | ログイン | - |
| GET | /api/auth/me | 現在のユーザー情報 | 必要 |
| GET | /api/products | 商品一覧 | - |
| GET | /api/products/:id | 商品詳細 | - |
| GET | /api/products/category/:category | カテゴリ別商品 | - |
| GET | /api/products/search | 商品検索 | - |
| GET | /api/cart | カート取得 | 必要 |
| POST | /api/cart/items | カートに追加 | 必要 |
| PUT | /api/cart/items/:id | カート更新 | 必要 |
| DELETE | /api/cart/items/:id | カートから削除 | 必要 |
| POST | /api/orders | 注文作成 | 必要 |
| GET | /api/orders | 注文履歴 | 必要 |
| POST | /api/payment/create-intent | 決済インテント作成 | 必要 |

### ライセンス

MIT License

---

## English

### Overview

A full-stack e-commerce application for selling Shiba Inu images. This portfolio project implements user authentication, product management, shopping cart, and Stripe payment functionality using modern web technologies.

### Tech Stack

#### Frontend
- **Framework:** Next.js 14.2 (React 18.3)
- **Language:** TypeScript 5.3
- **Styling:** Tailwind CSS 3.4
- **State Management:** Zustand 4.5
- **HTTP Client:** Axios 1.7
- **Payment UI:** Stripe React

#### Backend
- **Language:** Go 1.24
- **Framework:** Gin 1.11
- **Database:** PostgreSQL 15 (GORM)
- **Cache:** Redis 7
- **Authentication:** JWT
- **Payment Processing:** Stripe

#### Infrastructure
- **Containerization:** Docker / Docker Compose
- **Deployment:** Render
- **Database Hosting:** Supabase
- **Cache Hosting:** Upstash
- **Image Storage:** Cloudinary

### Key Features

#### User Features
- User registration and login (JWT authentication)
- Product listing and detail view
- Category-based product filtering
- Product search
- Shopping cart management
- Order processing with Stripe payment
- Order history

#### Admin Features
- Product CRUD operations
- User management
- Order management and status updates

### Project Structure

```
EC-site/
├── frontend/                # Next.js frontend
│   ├── app/                 # App Router pages
│   ├── components/          # Reusable components
│   ├── stores/              # Zustand state management
│   ├── lib/                 # Utilities
│   └── types/               # TypeScript type definitions
│
├── backend/                 # Go backend API
│   ├── cmd/server/          # Entry point
│   ├── internal/
│   │   ├── model/           # Data models
│   │   ├── repository/      # Data access layer
│   │   ├── service/         # Business logic layer
│   │   ├── handler/         # HTTP handlers
│   │   └── middleware/      # Middleware
│   ├── config/              # Configuration
│   └── pkg/                 # Shared packages
│
├── database/                # Database schema and migrations
├── docs/                    # Documentation (ER diagram, etc.)
├── docker-compose.yml       # Local development environment
└── render.yaml              # Deployment configuration
```

### Setup

#### Prerequisites
- Node.js 18 or higher
- Go 1.24 or higher
- Docker / Docker Compose
- Stripe account

#### Starting Local Development Environment

1. Clone the repository
```bash
git clone https://github.com/Naonao3/EC-site.git
cd EC-site
```

2. Start database and cache with Docker Compose
```bash
docker-compose up -d
```

3. Start the backend
```bash
cd backend
cp .env.production.example .env
# Edit .env file to configure environment variables
go run cmd/server/main.go
```

4. Start the frontend
```bash
cd frontend
npm install
npm run dev
```

5. Access http://localhost:3000 in your browser

### Environment Variables

#### Backend
```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=ec_site
DB_SSLMODE=disable
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h
STRIPE_SECRET_KEY=sk_test_xxx
STRIPE_WEBHOOK_SECRET=whsec_xxx
CORS_ORIGIN=http://localhost:3000
```

#### Frontend
```
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_xxx
```

### API Endpoints

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | /api/auth/register | User registration | - |
| POST | /api/auth/login | Login | - |
| GET | /api/auth/me | Current user info | Required |
| GET | /api/products | Product list | - |
| GET | /api/products/:id | Product detail | - |
| GET | /api/products/category/:category | Products by category | - |
| GET | /api/products/search | Product search | - |
| GET | /api/cart | Get cart | Required |
| POST | /api/cart/items | Add to cart | Required |
| PUT | /api/cart/items/:id | Update cart | Required |
| DELETE | /api/cart/items/:id | Remove from cart | Required |
| POST | /api/orders | Create order | Required |
| GET | /api/orders | Order history | Required |
| POST | /api/payment/create-intent | Create payment intent | Required |

### License

MIT License
