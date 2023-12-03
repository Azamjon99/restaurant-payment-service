package services

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/services"
)

type CardApplicationService interface {
	SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}

type cardApplicationSvcImpl struct {
	cardService services.PaymentCardService
}

func NewCardApplicationService(cardService services.PaymentCardService) CardApplicationService {
	return &cardApplicationSvcImpl{
		cardService: cardService,
	}
}

func (c *cardApplicationSvcImpl) SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error) {
	return nil, nil
}

func (c *cardApplicationSvcImpl) DeleteCard(ctx context.Context, cardID string) error {
	return nil
}
func (c *cardApplicationSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	return nil, nil
}
