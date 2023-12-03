package card

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/card/repositories"

	"gorm.io/gorm"
)

type PaymentCardRepository interface {
	SaveCard(ctx context.Context, card *models.PaymentCard) error
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}

type paymentCardRepoImpl struct {
	db *gorm.DB
}

const CARD_TABLE = "payment.payment_cards"

func NewCardRepository(db *gorm.DB) repositories.PaymentCardRepository {
	return &paymentCardRepoImpl{
		db: db,
	}
}

func (c *paymentCardRepoImpl) SaveCard(ctx context.Context, card *models.PaymentCard) error {

	result := c.db.WithContext(ctx).Table(CARD_TABLE).Save(card)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *paymentCardRepoImpl) DeleteCard(ctx context.Context, cardID string) error {

	var card models.PaymentCard

	result := c.db.WithContext(ctx).Table(CARD_TABLE).Delete(&card, "id = ?", cardID)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *paymentCardRepoImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {

	var card models.PaymentCard

	res := c.db.WithContext(ctx).Table(CARD_TABLE).First(&card, "id = ?", cardID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &card, nil
}
