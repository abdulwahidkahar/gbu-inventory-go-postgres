package inventory

import "time"

type Item struct {
	ID            int64     `json:"id"`
	SKU           string    `json:"sku"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	SupplierName  string    `json:"supplier_name"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
	Stock         int       `json:"stock"`
	ReorderLevel  int       `json:"reorder_level"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Payload struct {
	SKU           string  `json:"sku"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	SupplierName  string  `json:"supplier_name"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
	ReorderLevel  int     `json:"reorder_level"`
	IsActive      bool    `json:"is_active"`
}
