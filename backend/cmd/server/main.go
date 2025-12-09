package main

import (
	"log"

	"github.com/Naonao3/EC-site/backend/config"
	"github.com/Naonao3/EC-site/backend/internal/handler"
	"github.com/Naonao3/EC-site/backend/internal/middleware"
	"github.com/Naonao3/EC-site/backend/internal/model"
	"github.com/Naonao3/EC-site/backend/internal/repository"
	"github.com/Naonao3/EC-site/backend/internal/service"
	"github.com/Naonao3/EC-site/backend/pkg/database"
	redisClient "github.com/Naonao3/EC-site/backend/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 設定の読み込み
	cfg := config.Load()

	// データベース接続
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.CloseDB(db)

	// マイグレーション実行（Paymentを追加）
	if err := db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
		&model.CartItem{},
		&model.Payment{}, // NEW
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")

	// Redis接続
	redis, err := redisClient.NewRedisClient(cfg)
	if err != nil {
		log.Println("Warning: Failed to connect to Redis:", err)
	} else {
		defer redisClient.CloseRedis(redis)
	}

	// リポジトリの初期化
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	paymentRepo := repository.NewPaymentRepository(db) // NEW

	// サービスの初期化
	userService := service.NewUserService(userRepo, cfg.JWT.Secret)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	orderService := service.NewOrderService(orderRepo, cartRepo, productRepo)
	paymentService := service.NewPaymentService(paymentRepo, orderRepo, cfg.Stripe.SecretKey) // NEW

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)
	orderHandler := handler.NewOrderHandler(orderService, db)
	paymentHandler := handler.NewPaymentHandler(paymentService, cfg.Stripe.WebhookSecret) // NEW

	// Ginルーターの初期化
	router := gin.Default()

	// ミドルウェアの設定
	router.Use(middleware.CORSMiddleware())

	// ヘルスチェック
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// APIルート
	api := router.Group("/api")
	{
		// 認証不要のルート
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			// 認証が必要なルート
			auth.GET("/me", middleware.AuthMiddleware(cfg.JWT.Secret), userHandler.GetProfile)
		}

		// 商品関連（認証不要）
		products := api.Group("/products")
		{
			products.GET("", productHandler.ListProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.GET("/category/:category", productHandler.GetProductsByCategory)
			products.GET("/search", productHandler.SearchProducts)
		}

		// Stripe Webhook（認証不要）NEW
		api.POST("/webhooks/stripe", paymentHandler.HandleWebhook)  // StripeWebhook → HandleWebhook

		// 認証が必要なルート
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			// ユーザー関連
			users := authenticated.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateUser)
			}

			// カート関連
			cart := authenticated.Group("/cart")
			{
				cart.GET("", cartHandler.GetCart)
				cart.POST("/items", cartHandler.AddToCart)
				cart.PUT("/items/:id", cartHandler.UpdateCartItem)
				cart.DELETE("/items/:id", cartHandler.RemoveFromCart)
				cart.DELETE("", cartHandler.ClearCart)
			}

			// 注文関連
			orders := authenticated.Group("/orders")
			{
				orders.POST("", orderHandler.CreateOrder)
				orders.GET("", orderHandler.GetUserOrders)
				orders.GET("/:id", orderHandler.GetOrderByID)
			}

			// 決済関連（NEW）
			payment := authenticated.Group("/payment")
			{
				payment.POST("/create-intent", paymentHandler.CreatePaymentIntent)
				payment.GET("/order", paymentHandler.GetPaymentByOrderID)
			}

			// 管理者専用ルート
			admin := authenticated.Group("/admin")
			admin.Use(middleware.AdminMiddleware())
			{
				// ユーザー管理
				admin.GET("/users", userHandler.ListUsers)
				admin.GET("/users/:id", userHandler.GetUserByID)
				admin.DELETE("/users/:id", userHandler.DeleteUser)

				// 商品管理
				admin.POST("/products", productHandler.CreateProduct)
				admin.PUT("/products/:id", productHandler.UpdateProduct)
				admin.DELETE("/products/:id", productHandler.DeleteProduct)

				// 注文管理
				admin.GET("/orders", orderHandler.GetAllOrders)
				admin.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)
			}
		}
	}

	// サーバー起動
	log.Printf("Server starting on port %s...", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}