package http

import (
	"product-service/internal/domain/product"
	"product-service/internal/infrastructure/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(dbGorm *gorm.DB) (r *gin.Engine) {
	r = gin.New()
	r.SetTrustedProxies(nil) // Disable trusted proxies

	// Middleware jika diperlukan
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Init repo dan service
	repo := db.NewPostgresRepository(dbGorm)
	service := product.NewService(repo)

	// Pass services and *gin.Engine to init gin handler
	InitGinHandler(r, service)

	return
}
