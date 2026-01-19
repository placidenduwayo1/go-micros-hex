package dtos

import (
	"time"
)

type ProductResponse struct {
	ID            int64         `jsons:"id"`
	Sku           string        `json:"sku"`
	Category      string        `json:"category"`
	ProductName   string        `json:"product_name"`
	Description   string        `json:"description"`
	PriceResponse PriceResponse `json:"price"`
	CreatedAt     time.Time     `json:"ceated_at"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
	IsActive      bool          `json:"is_active"`
}

type PriceResponse struct {
	UnitPrice int64  `json:"unit_price"`
	Currency  string `json:"currency"`
}

type StockResponse struct {
	ID              int64           `json:"id"`
	ProductId       int64           `json:"product_id"`
	Quantity        int64           `json:"stock_quantity"`
	UpdatedAt       time.Time       `json:"updated_at"`
	ProductResponse ProductResponse `json:"product"`
}

type Response struct {
	Status    Status `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `jsnon:"created_at"`
}
type Status string

const (
	Success Status = "SUCCESS"
	Fail    Status = "FAILED"
)

func NewResponse(s Status, m string) *Response {
	return &Response{
		Status:    s,
		Message:   m,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
