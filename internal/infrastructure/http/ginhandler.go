package http

import (
	"product-service/internal/domain/product"
	productHandler "product-service/internal/interfaces/productHandler"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func InitGinHandler(r *gin.Engine, productService product.Service, store persistence.CacheStore) {
	handler := productHandler.NewHandler(productService)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product", handler.GetAllProduct)
	// r.GET("/product", cache.CachePage(store, time.Second*20, handler.GetAllProduct))
}
