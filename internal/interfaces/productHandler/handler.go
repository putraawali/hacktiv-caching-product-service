package productHandler

import (
	"log"
	"net/http"
	"product-service/internal/domain/product"
	"product-service/pkg/productDTO"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service product.Service
}

type Handler interface {
	CreateProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
}

func NewHandler(s product.Service) Handler {
	return &handler{service: s}
}

func (h *handler) CreateProduct(c *gin.Context) {
	var param productDTO.Product

	if err := c.ShouldBindBodyWithJSON(&param); err != nil {
		log.Printf("Failed to bind request JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productData := product.Product{
		Name:  param.Name,
		Brand: param.Brand,
	}

	if err := h.service.CreateProduct(&productData); err != nil {
		log.Printf("Failed create product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"product_id": productData.ID,
	})
}

func (h *handler) GetAllProduct(c *gin.Context) {
	products, err := h.service.GetAllProduct()
	if err != nil {
		log.Printf("Failed get product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
