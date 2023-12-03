package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/models"
)

type PaymentCardRepository interface {
	SaveCard(ctx context.Context, card *models.PaymentCard) error
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}
