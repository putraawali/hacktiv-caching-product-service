package product

import (
	"errors"
	"fmt"
	"product-service/internal/domain/inventory"
	"product-service/internal/infrastructure/rest"
	"product-service/pkg/inventoryDTO"
	"sync"
)

type service struct {
	repo     Repository
	restRepo inventory.Repository
	cache    sync.Map
}

type Service interface {
	CreateProduct(product *Product) (err error)
	GetAllProduct() (products []Product, err error)
}

func NewService(r Repository) Service {
	return &service{
		repo:     r,
		restRepo: rest.NewRestRepository(),
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

	err = s.restRepo.CreateInventory(inventoryDTO.CreateInventoryRequest{
		ProductID: int64(product.ID),
		Stock:     100,
		Location:  "Warehouse A",
	})

	return
}

func (s *service) GetAllProduct() (products []Product, err error) {
	if chachedProduct, found := s.cache.Load("products"); found {
		fmt.Println("Data diakses dari cache")
		return chachedProduct.([]Product), nil
	}

	fmt.Println("Tidak ditemukan data dari cache")
	if products, err = s.repo.FindAll(); err != nil {
		return
	}

	s.cache.Store("products", products)

	return
}
