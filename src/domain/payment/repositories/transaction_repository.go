package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/models"
)

type TransactionRepository interface {
	SaveTx(ctx context.Context, tx *models.Transaction) error
	UpdateTx(ctx context.Context, tx *models.Transaction) error
	GetTx(ctx context.Context, txID string) (*models.Transaction, error)
}
