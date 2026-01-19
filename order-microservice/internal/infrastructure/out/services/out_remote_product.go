package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

// RemoteProductServiceImpl Implemnet intreface RemoteProductService
type RemoteProductServiceImpl struct {
	baseUrl string
}

func NewRemoteProductServiceImpl(url string) *RemoteCustomerServiceImpl {
	return &RemoteCustomerServiceImpl{baseUrl: url}
}

func (o *RemoteCustomerServiceImpl) GetRemoteProductByID(ctx context.Context, id int64) (domain.Product, error) {
	remoteApiUrl := fmt.Sprintf(o.baseUrl+"/%d", id)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, remoteApiUrl, nil)
	if err != nil {
		return domain.Product{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return domain.Product{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return domain.Product{}, errors.New("remote product not found")
	}
	if response.StatusCode != http.StatusOK {
		return domain.Product{}, fmt.Errorf("remote product service error: status %d", response.StatusCode)
	}
	// Decoder le remote dto
	var remoteProductResponse dtos.ProductResponse
	if err := json.NewDecoder(response.Body).Decode(&remoteProductResponse); err != nil {
		return domain.Product{}, err
	}

	domainProduct := domain.Product{
		ID:          remoteProductResponse.ID,
		Sku:         remoteProductResponse.Sku,
		ProductName: remoteProductResponse.ProductName,
		Description: remoteProductResponse.Description,
		Price: domain.Price{
			UnitPrice: remoteProductResponse.PriceResponse.UnitPrice,
			Currency:  domain.Currency(remoteProductResponse.PriceResponse.Currency),
		},
		IsActive: remoteProductResponse.IsActive,
	}

	return domainProduct, nil
}
