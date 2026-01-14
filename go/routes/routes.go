package routes

import (
	"go-admin/handlers"
	"go-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// API route group
	api := r.Group("/api")
	{
		// Public routes - no authentication required
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// Protected routes - require JWT authentication
		user := api.Group("/user")
		user.Use(middleware.JWTAuth())
		{
			user.GET("/profile", handlers.GetProfile)
		}
	}
}
