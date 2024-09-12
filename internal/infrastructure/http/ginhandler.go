package http

import (
	"product-service/internal/domain/product"
	productHandler "product-service/internal/interfaces/productHandler"

	"github.com/gin-gonic/gin"
)

func InitGinHandler(r *gin.Engine, productService product.Service) {
	handler := productHandler.NewHandler(productService)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product", handler.GetAllProduct)
}
