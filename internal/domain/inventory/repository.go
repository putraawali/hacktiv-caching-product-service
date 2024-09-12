package inventory

import inventoryDTO "product-service/pkg/inventoryDTO"

type Repository interface {
	CreateInventory(req inventoryDTO.CreateInventoryRequest) (err error)
}
