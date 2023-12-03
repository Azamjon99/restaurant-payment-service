package payment

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/repositories"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTx(ctx context.Context, tx *models.Transaction) error
	UpdateTx(ctx context.Context, tx *models.Transaction) error
	GetTx(ctx context.Context, txID string) (*models.Transaction, error)
}

type txRepoImpl struct {
	db *gorm.DB
}

const TRANSACTION_TABLE = "payment.transactions"

func NewTransactionRepository(db *gorm.DB) repositories.TransactionRepository {
	return &txRepoImpl{
		db: db,
	}
}

func (t *txRepoImpl) SaveTx(ctx context.Context, tx *models.Transaction) error {

	result := t.db.WithContext(ctx).Table(TRANSACTION_TABLE).Save(tx)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *txRepoImpl) UpdateTx(ctx context.Context, tx *models.Transaction) error {

	result := t.db.WithContext(ctx).Table(TRANSACTION_TABLE).Save(&tx)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *txRepoImpl) GetTx(ctx context.Context, txID string) (*models.Transaction, error) {

	var tx models.Transaction

	res := t.db.WithContext(ctx).Table(TRANSACTION_TABLE).First(&tx, "id = ?", txID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &tx, nil
}
