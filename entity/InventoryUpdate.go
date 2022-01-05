package entity

type InventoryUpdate struct {
	SkuId        string `json:"sku_id,omitempty"`
	WarehousesId string `json:"warehouses_id,omitempty"`
	Limited      bool   `json:"unlimitedQuantity,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
}
