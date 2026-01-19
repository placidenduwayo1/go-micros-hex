package out

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/domain"
)

type OutOrderService interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrderByID(ctx context.Context, id int64) (domain.Order, error)
	GetAllOrder(ctx context.Context) ([]domain.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}
type OutOrderLineService interface {
	CreateOrderLine(ctx context.Context, orderLine domain.OrderLine) (domain.OrderLine, error)
	GetOrderLineByID(ctx context.Context, id int64) (domain.OrderLine, error)
	GetAllOrderLines(ctx context.Context) ([]domain.OrderLine, error)
	UpdateOrderLine(ctx context.Context, id int64, orderLine domain.OrderLine) (domain.OrderLine, error)
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
