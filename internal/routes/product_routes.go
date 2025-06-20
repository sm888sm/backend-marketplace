package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sm888sm/backend-marketplace/internal/config"
	"github.com/sm888sm/backend-marketplace/internal/controllers"
	"github.com/sm888sm/backend-marketplace/internal/middleware"
)

func SetupProductRoutes(router *gin.Engine, productController *controllers.ProductController, cfg *config.Config) {
	public := router.Group("/api/products")
	{
		public.GET("", productController.GetAllProducts)
		public.GET("/:id", productController.GetProductByID)
		public.GET(":id/images", productController.ListImages)
		public.GET("explore", productController.ExploreProducts)
	}

	merchant := router.Group("/api/merchant/products")
	merchant.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	merchant.Use(middleware.RoleMiddleware("merchant"))
	{
		merchant.POST("", productController.CreateProduct)
		merchant.PUT("/:id", productController.UpdateProduct)
		merchant.DELETE("/:id", productController.DeleteProduct)
		merchant.GET("/my-products", productController.GetMerchantProducts)
		merchant.POST(":id/images", productController.UploadImage)
	}
}
