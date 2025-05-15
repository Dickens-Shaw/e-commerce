package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title 电商系统 API
// @version 1.0
// @description 电商系统后端服务 API 文档
// @host localhost:8080
// @BasePath /
func main() {
	// 初始化 Viper 配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Config file not found: %v, using default config", err)
	}

	// 初始化 Zap 日志
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// 初始化 Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())

	// 日志中间件
	r.Use(func(c *gin.Context) {
		zap.L().Info("request", zap.String("path", c.Request.URL.Path))
		c.Next()
	})

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// TODO: 注册各业务模块路由

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	logger.Info("服务启动", zap.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Fatal("启动失败", zap.Error(err))
	}
}