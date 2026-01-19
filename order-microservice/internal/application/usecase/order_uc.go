package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

// OrderUseCase implement OrderService
type OrderUseCase struct {
	outOrderSvc out.OutOrderService
	remote      out.RemoteCustomerService
}

// NewOrderNewOrderUseCaseServiceImpl DI by contructor
func NewOrderUseCase(outOrderSvc out.OutOrderService, remote out.RemoteCustomerService) *OrderUseCase {
	return &OrderUseCase{outOrderSvc: outOrderSvc, remote: remote}
}

// CreateOrder implement OrderService
func (o *OrderUseCase) CreateOrder(ctx context.Context, customerID int64) (domain.Order, error) {
	if err := checkId(customerID); err != nil {
		return domain.Order{}, err
	}

	//call remote adapter to check remote customer
	remoteCustomer, err := o.remote.GetRemoteCustomerByID(ctx, customerID)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w: %v", errOccurred, err)
	}

	if remoteCustomer.Status != domain.Active {
		return domain.Order{}, errors.New("error: remote customer status not allowed")
	}
	var order = domain.Order{
		CustomerID: customerID,
		CreatedAt:  time.Now(),
		Status:     domain.Created,
	}
	//call output service to save order
	savedOrder, err := o.outOrderSvc.CreateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return savedOrder, nil

}

// GetOrderByID implement OrderService
func (o *OrderUseCase) GetOrderByID(ctx context.Context, id int64) (domain.Order, error) {
	if err := checkId(id); err != nil {
		return domain.Order{}, err
	}

	savedOrder, err := o.outOrderSvc.GetOrderByID(ctx, id)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return savedOrder, nil
}

// GetAllOrder implement OrderService
func (o *OrderUseCase) GetAllOrder(ctx context.Context) ([]domain.Order, error) {
	orders, err := o.outOrderSvc.GetAllOrder(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return orders, nil
}

// DeleteOrder implement OrderService
func (o *OrderUseCase) DeleteOrder(ctx context.Context, id int64) error {
	if err := checkId(id); err != nil {
		return err
	}
	order, err := o.GetOrderByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w:%v", errNotFound, err)
	}
	if order.Status != domain.Created {
		return errors.New("order can no longer be deleted")
	}
	if err := o.outOrderSvc.DeleteOrder(ctx, id); err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}

	return nil
}
