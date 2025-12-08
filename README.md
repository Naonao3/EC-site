# EC-site

本格的なEコマースサイトのフルスタック実装。Go（バックエンド）とNext.js（フロントエンド）で構築され、Stripe決済統合を含みます。

## 目次

- [機能](#機能)
- [技術スタック](#技術スタック)
- [前提条件](#前提条件)
- [セットアップ](#セットアップ)
- [環境変数の設定](#環境変数の設定)
- [アプリケーションの起動](#アプリケーションの起動)
- [購入フローのテスト](#購入フローのテスト)
- [プロジェクト構成](#プロジェクト構成)
- [API エンドポイント](#apiエンドポイント)

## 機能

### 実装済み機能

- ✅ **ユーザー認証・認可**
  - JWT ベースの認証
  - bcrypt によるパスワードハッシュ化
  - ユーザー登録・ログイン

- ✅ **商品管理**
  - 商品一覧表示
  - 商品詳細表示
  - カテゴリー別フィルタリング

- ✅ **ショッピングカート**
  - カートへの商品追加
  - 数量変更
  - カートからの削除
  - カートクリア

- ✅ **注文管理**
  - 注文作成（トランザクション処理）
  - 在庫管理（注文時に自動減算）
  - 注文履歴表示
  - 注文ステータス管理

- ✅ **決済処理（Stripe）**
  - Stripe Payment Intent による決済
  - テストモード対応
  - 決済完了後の注文確定

- ✅ **管理機能**
  - 全注文の閲覧
  - 注文ステータスの更新

## 技術スタック

### バックエンド

- **言語**: Go 1.24.0
- **フレームワーク**: Gin
- **ORM**: GORM
- **データベース**: PostgreSQL 16
- **キャッシュ**: Redis
- **認証**: JWT (golang-jwt/jwt)
- **決済**: Stripe Go SDK

### フロントエンド

- **フレームワーク**: Next.js 14.2.33
- **言語**: TypeScript
- **状態管理**: Zustand
- **UIライブラリ**: Tailwind CSS
- **決済UI**: Stripe React SDK

## 前提条件

開発環境に以下がインストールされている必要があります：

- Go 1.24.0 以上
- Node.js 18.0 以上
- PostgreSQL 16
- Redis
- Stripeアカウント（テストキー）

## セットアップ

### 1. リポジトリのクローン

```bash
git clone https://github.com/Naonao3/EC-site.git
cd EC-site
```

### 2. データベースのセットアップ

PostgreSQL にデータベースを作成：

```bash
psql -U postgres
CREATE DATABASE ec_site_db;
\q
```

### 3. バックエンドの依存関係インストール

```bash
cd backend
go mod download
```

### 4. フロントエンドの依存関係インストール

```bash
cd ../frontend
npm install
```

## 環境変数の設定

### バックエンド環境変数

`backend/.env` ファイルを作成（または編集）：

```bash
# データベース
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ec_site_db
DB_SSLMODE=disable

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT認証
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Stripe (テストモード)
STRIPE_SECRET_KEY=sk_test_your_stripe_secret_key_here
STRIPE_WEBHOOK_SECRET=whsec_your_webhook_secret_here

# サーバー設定
PORT=8080
GIN_MODE=debug

# CORS設定
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

### フロントエンド環境変数

`frontend/.env.local` ファイルを作成（または編集）：

```bash
# バックエンドAPI
NEXT_PUBLIC_API_URL=http://localhost:8080

# Stripe (公開可能キー)
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_your_stripe_publishable_key_here
```

### Stripe APIキーの取得方法

1. [Stripe Dashboard](https://dashboard.stripe.com) にログイン
2. 画面右上で **テストモード** に切り替え
3. 「開発者」→「APIキー」にアクセス
4. 公開可能キー（`pk_test_...`）とシークレットキー（`sk_test_...`）をコピー
5. 上記の環境変数ファイルに貼り付け

## アプリケーションの起動

### 1. PostgreSQL と Redis の起動

```bash
# PostgreSQL
sudo service postgresql start

# Redis
sudo service redis-server start
```

### 2. バックエンドの起動

```bash
cd backend
go run cmd/server/main.go
```

バックエンドは http://localhost:8080 で起動します。

### 3. フロントエンドの起動

別のターミナルで：

```bash
cd frontend
npm run dev
```

フロントエンドは http://localhost:3000 で起動します。

### 4. サンプルデータの投入（オプション）

商品データがない場合、PostgreSQL で以下を実行：

```sql
INSERT INTO products (name, description, price, stock, category, image_url, created_at, updated_at) VALUES
('ノートパソコン', 'Core i7搭載の高性能ノートパソコン', 120000, 10, 'Electronics', 'https://via.placeholder.com/300x200?text=Laptop', NOW(), NOW()),
('ワイヤレスマウス', '静音設計のワイヤレスマウス', 3500, 50, 'Electronics', 'https://via.placeholder.com/300x200?text=Mouse', NOW(), NOW()),
('キーボード', 'メカニカルキーボード RGB対応', 8500, 30, 'Electronics', 'https://via.placeholder.com/300x200?text=Keyboard', NOW(), NOW()),
('Webカメラ', 'Full HD対応Webカメラ', 6500, 25, 'Electronics', 'https://via.placeholder.com/300x200?text=Webcam', NOW(), NOW()),
('USBハブ', 'USB3.0対応7ポートハブ', 2500, 40, 'Electronics', 'https://via.placeholder.com/300x200?text=USB+Hub', NOW(), NOW());
```

## 購入フローのテスト

### 1. ユーザー登録

http://localhost:3000/register にアクセスしてアカウントを作成

### 2. ログイン

作成したアカウントでログイン

### 3. 商品を選択

商品一覧から任意の商品をカートに追加

### 4. チェックアウト

カートページから「チェックアウトへ進む」をクリック

### 5. Stripe テストカードで決済

以下のテスト情報を使用：

- **カード番号**: `4242 4242 4242 4242`
- **有効期限**: 任意の未来の日付（例: `12/34`）
- **CVC**: 任意の3桁（例: `123`）
- **郵便番号**: 任意（例: `12345`）

### 6. 決済完了

決済が成功すると、決済完了ページにリダイレクトされます。

## プロジェクト構成

```
EC-site/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # エントリーポイント
│   ├── internal/
│   │   ├── handler/                 # HTTPハンドラー
│   │   │   ├── auth_handler.go
│   │   │   ├── cart_handler.go
│   │   │   ├── order_handler.go
│   │   │   ├── payment_handler.go
│   │   │   ├── product_handler.go
│   │   │   └── user_handler.go
│   │   ├── middleware/              # ミドルウェア
│   │   │   └── auth.go
│   │   ├── model/                   # データモデル
│   │   │   ├── cart.go
│   │   │   ├── order.go
│   │   │   ├── payment.go
│   │   │   ├── product.go
│   │   │   └── user.go
│   │   ├── repository/              # データアクセス層
│   │   │   ├── cart_repository.go
│   │   │   ├── order_repository.go
│   │   │   ├── payment_repository.go
│   │   │   ├── product_repository.go
│   │   │   └── user_repository.go
│   │   └── service/                 # ビジネスロジック層
│   │       ├── auth_service.go
│   │       ├── cart_service.go
│   │       ├── order_service.go
│   │       ├── payment_service.go
│   │       ├── product_service.go
│   │       └── user_service.go
│   ├── pkg/
│   │   ├── config/                  # 設定管理
│   │   └── database/                # DB接続
│   ├── .env                         # 環境変数
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── app/                         # App Router ページ
│   │   ├── cart/
│   │   ├── checkout/
│   │   ├── login/
│   │   ├── orders/
│   │   ├── products/
│   │   └── register/
│   ├── components/                  # Reactコンポーネント
│   │   ├── layout/
│   │   ├── stripe/
│   │   └── ui/
│   ├── lib/                         # ユーティリティ
│   │   ├── api-client.ts           # API クライアント
│   │   └── stripe.ts               # Stripe 設定
│   ├── stores/                      # Zustand ストア
│   │   ├── authStore.ts
│   │   ├── cartStore.ts
│   │   └── checkoutStore.ts
│   ├── types/                       # TypeScript型定義
│   │   └── index.ts
│   ├── .env.local                   # 環境変数
│   ├── package.json
│   └── tsconfig.json
└── README.md
```

## API エンドポイント

### 認証

- `POST /api/auth/register` - ユーザー登録
- `POST /api/auth/login` - ログイン
- `GET /api/auth/me` - ログインユーザー情報取得（要認証）

### 商品

- `GET /api/products` - 商品一覧取得
- `GET /api/products/:id` - 商品詳細取得

### カート

- `GET /api/cart` - カート取得（要認証）
- `POST /api/cart/items` - カートに商品追加（要認証）
- `PUT /api/cart/items/:id` - カートアイテム更新（要認証）
- `DELETE /api/cart/items/:id` - カートから商品削除（要認証）
- `DELETE /api/cart` - カートクリア（要認証）

### 注文

- `POST /api/orders` - 注文作成（要認証）
- `GET /api/orders` - 注文一覧取得（要認証）
- `GET /api/orders/:id` - 注文詳細取得（要認証）

### 決済

- `POST /api/payment/create-payment-intent` - Payment Intent作成（要認証）
- `GET /api/payment/order` - 注文の決済情報取得（要認証）
- `POST /api/payment/webhook` - Stripe Webhook

## アーキテクチャ

### バックエンド

クリーンアーキテクチャパターンを採用：

```
Handler（HTTPレイヤー）
    ↓
Service（ビジネスロジック層）
    ↓
Repository（データアクセス層）
    ↓
Database/External Services
```

### 主な設計パターン

- **依存性注入**: サービスとリポジトリの疎結合化
- **トランザクション管理**: 注文作成時の一貫性保証
- **JWT 認証**: ステートレスな認証機構
- **エラーハンドリング**: 統一されたエラーレスポンス

### フロントエンド

- **App Router**: Next.js 14の最新ルーティング
- **状態管理**: Zustandによる軽量な状態管理
- **型安全性**: TypeScriptによる完全な型チェック
- **API クライアント**: 統一されたAPI通信レイヤー

## 開発メモ

### 既知の問題

- Webhook シークレットは未設定（ローカル開発では不要）
- 画像アップロード機能は未実装（プレースホルダー画像使用）

### 今後の改善案

- [ ] 管理画面の実装
- [ ] 商品画像のアップロード機能
- [ ] レビュー・評価機能
- [ ] お気に入り機能
- [ ] クーポン・割引機能
- [ ] メール通知機能
- [ ] 商品検索機能の強化
- [ ] ページネーション最適化

## ライセンス

このプロジェクトはプライベートリポジトリです。

## コントリビューション

現在、外部からのコントリビューションは受け付けていません。

---

**開発者**: Naonao3
**最終更新**: 2025年12月
