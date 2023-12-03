package payment

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/repositories"

	"gorm.io/gorm"
)

type ReceiptRepository interface {
	SaveReceipt(ctx context.Context, receipt *models.Receipt) error
	GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error)
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
}

type receiptRepoImpl struct {
	db *gorm.DB
}

const RECEIPT_TABLE = "payment.receipts"

func NewReceiptRepository(db *gorm.DB) repositories.ReceiptRepository {
	return &receiptRepoImpl{
		db: db,
	}
}

func (r *receiptRepoImpl) SaveReceipt(ctx context.Context, receipt *models.Receipt) error {

	result := r.db.WithContext(ctx).Table(RECEIPT_TABLE).Save(receipt)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *receiptRepoImpl) GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {

	var receipt models.Receipt

	res := r.db.WithContext(ctx).Table(RECEIPT_TABLE).First(&receipt, "id = ?", receiptID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &receipt, nil
}

func (r *receiptRepoImpl) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {

	var receipt models.Receipt

	res := r.db.WithContext(ctx).Table(RECEIPT_TABLE).First(&receipt, "order_id = ?", orderID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &receipt, nil
}
