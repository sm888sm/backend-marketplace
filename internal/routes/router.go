package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sm888sm/backend-marketplace/internal/config"
	"github.com/sm888sm/backend-marketplace/internal/controllers"
	"github.com/sm888sm/backend-marketplace/internal/middleware"
	"github.com/sm888sm/backend-marketplace/internal/repositories"
	"github.com/sm888sm/backend-marketplace/internal/services"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, productRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService, categoryService, authService)
	orderController := controllers.NewOrderController(orderService, productService)
	categoryController := controllers.NewCategoryController(categoryService)

	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)

	SetupProductRoutes(router, productController, cfg)

	apiAuth := router.Group("/api")
	apiAuth.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	{
		apiAuth.POST("/orders", orderController.CreateOrder)
		apiAuth.GET("/orders/my-orders", orderController.GetOrdersByCustomer)
		apiAuth.GET("/orders/:id", orderController.GetOrderByID)
		apiAuth.GET("/merchant/orders/:id", orderController.GetOrderByIDMerchant)
		apiAuth.GET("/merchant/products/buyers", orderController.GetBuyersByMerchant)

		apiAuth.POST("/categories", categoryController.CreateCategory)
		apiAuth.PUT("/categories/:id", categoryController.UpdateCategory)
		apiAuth.DELETE("/categories/:id", categoryController.DeleteCategory)

		apiAuth.GET("/users", authController.ListUsers)
		apiAuth.GET("/users/:id", authController.GetUserByID)
	}

	router.GET("/api/categories", categoryController.ListCategories)

	return router
}
