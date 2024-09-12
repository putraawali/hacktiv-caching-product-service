package rest

import (
	"fmt"
	"os"
	"product-service/internal/domain/inventory"
	inventoryDTO "product-service/pkg/inventoryDTO"

	"github.com/go-resty/resty/v2"
)

type restRepository struct{}

func NewRestRepository() inventory.Repository {
	return &restRepository{}
}

func (r *restRepository) CreateInventory(req inventoryDTO.CreateInventoryRequest) (err error) {
	client := resty.New()

	client = client.SetDebug(true)

	resp := map[string]interface{}{}

	_, err = client.R().
		SetBody(req).
		SetResult(&resp).
		SetError(&resp).
		Post(fmt.Sprintf("%s/inventory", os.Getenv("INVENTORY_SERVICE_HOST")))

	return
}
