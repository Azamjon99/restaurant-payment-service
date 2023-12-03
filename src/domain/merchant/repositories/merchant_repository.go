package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/models"
)

type MerchantRepository interface {
	SaveMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}
