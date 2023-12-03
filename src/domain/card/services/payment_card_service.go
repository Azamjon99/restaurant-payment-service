package services

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/repositories"
)

type PaymentCardService interface {
	SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}

type paymentCardSvcImpl struct {
	paymentCardRepo repositories.PaymentCardRepository
}

func NewCardService(paymentCardRepo repositories.PaymentCardRepository) PaymentCardService {
	return &paymentCardSvcImpl{
		paymentCardRepo: paymentCardRepo,
	}
}

func (p *paymentCardSvcImpl) SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error) {

	// err := p.paymentCardRepo.SaveCard(ctx, card)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (p *paymentCardSvcImpl) DeleteCard(ctx context.Context, cardID string) error {

	err := p.paymentCardRepo.DeleteCard(ctx, cardID)

	if err != nil {
		return err
	}

	return nil
}

func (p *paymentCardSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {

	// card, err := p.paymentCardRepo.GetCard(ctx, cardID)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
