package routes

import (
	"go-admin/handlers"
	"go-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// API 路由组
	api := r.Group("/api")
	{
		// 公开路由 - 不需要认证
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// 受保护路由 - 需要 JWT 认证
		user := api.Group("/user")
		user.Use(middleware.JWTAuth())
		{
			user.GET("/profile", handlers.GetProfile)
		}
	}
}
