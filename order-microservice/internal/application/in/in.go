package in

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/domain"
)

type InOrderService interface {
	CreateOrder(ctx context.Context, customerID int64) (domain.Order, error)
	GetOrderByID(ctx context.Context, id int64) (domain.Order, error)
	GetAllOrder(ctx context.Context) ([]domain.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}
type InOrderLineService interface {
	CreateOrderLine(ctx context.Context, orderID int64, orderLine domain.OrderLine) (domain.OrderLine, error)
	GetOrderLineByID(ctx context.Context, id int64) (domain.OrderLine, error)
	GetAllOrderLines(ctx context.Context) ([]domain.OrderLine, error)
	SetOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error)
	IncreaseOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error)
	DecreaseOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error)
	DeleteOrderLine(ctx context.Context, id int64) error
}

// RemoteCustomerService to get remote customer
type RemoteCustomerService interface {
	GetRemoteCustomerByID(ctx context.Context, id int64) (domain.Customer, error)
}

// RemoteProductService to get remote products
type RemoteProductService interface {
	GetRemoteProductByID(ctx context.Context, id int64) (domain.Product, error)
}
