package dtos

type ProductResponse struct {
	ID            int64         `jsons:"id"`
	Sku           string        `json:"sku"`
	Category      string        `json:"category"`
	ProductName   string        `json:"product_name"`
	Description   string        `json:"description"`
	PriceResponse PriceResponse `json:"price"`
	IsActive      bool          `json:"is_active"`
}

type PriceResponse struct {
	UnitPrice int64  `json:"unit_price"`
	Currency  string `json:"currency"`
}
