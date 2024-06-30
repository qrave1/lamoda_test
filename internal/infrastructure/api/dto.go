package api

type ProductsRequest struct {
	ProductCodes []string `json:"product_codes"`
}

type InventoryRequest struct {
	WarehouseID int `json:"warehouse_id"`
}
