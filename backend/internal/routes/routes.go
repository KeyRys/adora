package routes

import (
	"backend/internal/delivery/http"
	"backend/internal/delivery/http/middleware"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(router *gin.Engine, db *pgx.Conn, secret string) {

	//----init----
	//product
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := http.NewProductHandler(productUsecase)
	//auth
	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, secret)
	authHandler := http.NewAuthHandler(authUsecase)
	//cart
	cartRepo := repository.NewCartRepository(db)
	cartUsecase := usecase.NewCartUsecase(cartRepo)
	cartHandler := http.NewCartHandler(cartUsecase)
	//profile
	profileRepo := repository.NewProfileRepository(db)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)
	profileHandler := http.NewProfileHandler(profileUsecase)
	//seller
	sellerRepo := repository.NewSellerRepository(db)
	sellerUsecase := usecase.NewSellerUsecase(sellerRepo)
	sellerHandler := http.NewSellerHandler(sellerUsecase)
	//checkout
	checkoutRepo := repository.NewCheckoutRepository(db)
	checkoutUsecase := usecase.NewCheckoutUsecase(checkoutRepo)
	checkoutHandler := http.NewCheckoutHandler(checkoutUsecase)
	//order
	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := http.NewOrderHandler(orderUsecase)
	sellerOrderUsecase := usecase.NewSellerOrderUsecase(orderRepo)
	sellerOrderHandler := http.NewSellerOrderHandler(sellerOrderUsecase)

	//----routes----
	//product
	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})

	//auth
	auth := router.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	//protected routes - seller, cart, profile
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(secret))
	//seller
	protected.POST("/seller/up", sellerHandler.BecomeSeller)
	protected.GET("/seller/rabbits", sellerHandler.GetSellerRabbits)
	protected.POST("/seller/rabbits", sellerHandler.CreateRabbit)
	protected.PUT("/seller/rabbits/:id", sellerHandler.UpdateRabbit)
	protected.DELETE("/seller/rabbits/:id", sellerHandler.DeleteRabbit)
	//profile
	protected.GET("/profile/me", profileHandler.GetMyProfile)
	//cart
	protected.POST("/cart/add", cartHandler.AddToCart)
	protected.GET("/cart", cartHandler.GetCart)
	protected.DELETE("/cart/item/:id", cartHandler.RemoveItem)
	//checkout
	protected.POST("/checkout", checkoutHandler.Checkout)
	//order
	protected.GET("/orders/me", orderHandler.GetBuyerOrders)
	protected.GET("/orders/seller", sellerOrderHandler.GetSellerOrders)
	protected.PUT("/orders/:id", sellerOrderHandler.UpdateOrderStatus)
}
