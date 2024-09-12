package db

import (
	"product-service/internal/domain/product"

	"gorm.io/gorm"
)

type pgRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) product.Repository {
	return &pgRepository{db}
}

func (pg *pgRepository) Create(inventory *product.Product) (err error) {
	return pg.db.Create(inventory).Error
}

func (pg *pgRepository) FindAll() (inventories []product.Product, err error) {
	err = pg.db.Find(&inventories).Error

	return
}
