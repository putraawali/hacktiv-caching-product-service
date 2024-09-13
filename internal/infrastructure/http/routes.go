package http

import (
	"log"
	"product-service/config"
	"product-service/internal/cache"
	"product-service/internal/domain/product"
	"product-service/internal/infrastructure/db"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func SetupRoute(store persistence.CacheStore) (r *gin.Engine) {
	r = gin.New()
	r.SetTrustedProxies(nil) // Disable trusted proxies

	// Middleware jika diperlukan
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	dbGorm, err := config.InitDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	redis := cache.InitRedis()

	// Init repo dan service
	repo := db.NewPostgresRepository(dbGorm)
	service := product.NewService(repo, redis)

	// Pass services and *gin.Engine to init gin handler
	InitGinHandler(r, service, store)

	return
}
