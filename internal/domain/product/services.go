package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"product-service/internal/domain/inventory"
	"product-service/internal/infrastructure/rest"
	"product-service/pkg/inventoryDTO"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type service struct {
	repo     Repository
	restRepo inventory.Repository
	cache    sync.Map
	redis    *redis.Client
}

const productKey = "products"

type Service interface {
	CreateProduct(product *Product) (err error)
	GetAllProduct() (products []Product, err error)
}

func NewService(r Repository, redis *redis.Client) Service {
	return &service{
		repo:     r,
		restRepo: rest.NewRestRepository(),
		redis:    redis,
	}
}

func (s *service) CreateProduct(product *Product) (err error) {
	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.Brand == "" {
		return errors.New("product brand is required")
	}

	if err = s.repo.Create(product); err != nil {
		return
	}

	s.restRepo.CreateInventory(inventoryDTO.CreateInventoryRequest{
		ProductID: int64(product.ID),
		Stock:     100,
		Location:  "Warehouse A",
	})

	fmt.Println("Delete data products from Redis")
	// Delete cache from redis
	if err = s.redis.Del(context.Background(), productKey).Err(); err != nil {
		fmt.Println(err.Error())
		err = nil
	}

	// Delete cache from sync.Map
	// s.cache.Delete("products")

	return
}

func (s *service) GetAllProduct() (products []Product, err error) {
	ctx := context.Background()

	// Cached using redis
	cachedProducts, err := s.redis.Get(ctx, productKey).Result()
	if err == nil { // Cache hit
		fmt.Println("Cache hit: Data diambil dari redis")
		if err = json.Unmarshal([]byte(cachedProducts), &products); err != nil {
			return
		}

		return
	}

	// Cached sync.Map redis
	// if chachedProduct, found := s.cache.Load(productKey); found {
	// 	fmt.Println("Data diakses dari cache")
	// 	return chachedProduct.([]Product), nil
	// }

	fmt.Println("Tidak ditemukan data dari cache")
	if products, err = s.repo.FindAll(); err != nil {
		return
	}

	productData, err := json.Marshal(products)
	if err != nil {
		return
	}

	fmt.Println("Menyimpan data ke redis")
	if err = s.redis.Set(ctx, productKey, productData, time.Second*20).Err(); err != nil {
		fmt.Println(err.Error())
		err = nil
	}

	// Store data to cache using sync.Map
	// s.cache.Store(productKey, products)

	return
}
