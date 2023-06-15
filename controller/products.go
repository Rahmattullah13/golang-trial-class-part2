package controller

import (
	"golang/trial-class/part2/config"
	"golang/trial-class/part2/entity"
	"golang/trial-class/part2/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductId    int    `json:"product_id"`
}

// @Summary Get Product
// @Schemes Product
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} entity.Product
// @Router /products [get]
func HandlerGetProduct(ctx *gin.Context) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// @Summary Post Orders
// @Schemes Orders
// @Description Create new orders
// @Tags Orders
// @Produce json
// @Success 200 {array} entity.Order
// @Router /orders [post]
func HandlerPostOrder(ctx *gin.Context) {
	var orderBody OrderRequest
	if err := ctx.ShouldBindJSON(&orderBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Terjadi kesalahan parsing data",
		})
		return
	}

	var product entity.Product
	if result := config.DB.Where("ID = ?", orderBody.ProductId).First(&product); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Produk tidak ditemukan",
		})
		return
	}

	newOrder := entity.Order{
		BuyerEmail:   orderBody.BuyerEmail,
		BuyerAddress: orderBody.BuyerAddress,
		ProductId:    int(product.ID),
		OrderDate:    time.Now(),
	}

	if err := config.DB.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Terjadi kesalahan saat membuat order",
		})
		return
	}

	// trigger mailer service of send mail
	service.SendMail(newOrder.BuyerEmail, newOrder.BuyerAddress, product.Name)

	ctx.JSON(http.StatusOK, "Berhasil membuat order!")
}

// @Summary Get Orders
// @Schemes Orders
// @Description Get list of all orders
// @Tags Orders
// @Produce json
// @Success 200 {array} entity.Order
// @Router /orders [get]
func HandlerGetOrders(ctx *gin.Context) {
	var orders []entity.Order
	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
