package main

import (
	"log"

	"go-admin/config"
	"go-admin/database"
	"go-admin/models"
	"go-admin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.Init()

	// 连接数据库
	database.Init()

	// 自动迁移表结构
	models.Migrate()

	// 创建 Gin 引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务
	log.Printf("Server starting on port %s...", config.AppConfig.ServerPort)
	if err := r.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
