package main

import (
	"golang/trial-class/part2/config"
	"golang/trial-class/part2/controller"

	_ "golang/trial-class/part2/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Trial Class Mini Ecommerce
// @description Dokumentasi REST API
// @version 0.1
// @host localhost:8000

func main() {
	s := gin.Default()
	config.DBConnect()

	s.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})

	s.GET("/products", controller.HandlerGetProduct)

	s.POST("/orders", controller.HandlerPostOrder)

	s.GET("/orders", controller.HandlerGetOrders)

	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Run("localhost:8000")
}
