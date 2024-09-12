package inventoryDTO

type CreateInventoryRequest struct {
	ProductID int64  `json:"product_id"`
	Stock     int    `json:"stock"`
	Location  string `json:"location"`
}
