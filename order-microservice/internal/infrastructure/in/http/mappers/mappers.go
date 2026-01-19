package mappers

import (
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

func ToDomainCustomer(dtoResp dtos.CustomerResponse) domain.Customer {
	return domain.Customer{
		ID:          dtoResp.ID,
		Firstname:   dtoResp.Firstname,
		Lastname:    dtoResp.Lastname,
		Genda:       domain.Genda(dtoResp.Genda),
		Email:       dtoResp.Email,
		PhoneNumber: dtoResp.PhoneNumber,
		Status:      domain.CustomerStatus(dtoResp.Status),
	}
}

func ToBusinessOrderLine(request dtos.OrderLineRequest) domain.OrderLine {
	return domain.OrderLine{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
}

func ToOrderLineResponse(orderLine domain.OrderLine, product domain.Product) dtos.OrderLineResponse {
	return dtos.OrderLineResponse{
		ID: orderLine.ID,
		ProductResponse: dtos.ProductResponse{
			ID:          product.ID,
			Sku:         product.Sku,
			ProductName: product.ProductName,
			Description: product.Description,
			PriceResponse: dtos.PriceResponse{
				UnitPrice: product.Price.UnitPrice,
				Currency:  string(product.Price.Currency),
			},
		},
		Quantity: orderLine.Quantity,
	}
}
