package usecase

import (
	"context"
	"fmt"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

type RemoteProductServiceImpl struct {
	outSvc out.RemoteProductService
}

func NewRemoteProductServiceImpl(outS out.RemoteProductService) *RemoteProductServiceImpl {
	return &RemoteProductServiceImpl{outSvc: outS}
}

func (o *RemoteProductServiceImpl) GetRemoteProductByID(ctx context.Context, id int64) (domain.Product, error) {
	if err := checkId(id); err != nil {
		return domain.Product{}, err
	}

	bsProduct, err := o.outSvc.GetRemoteProductByID(ctx, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return bsProduct, nil
}
