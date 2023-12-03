package merchant

import (
	"context"

	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/repositories"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	SaveMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}

type merchantRepoImpl struct {
	db *gorm.DB
}

const MERCHANT_TABLE = "payment.merchant_settings"

func NewMerchantRepository(db *gorm.DB) repositories.MerchantRepository {
	return &merchantRepoImpl{
		db: db,
	}
}

func (m *merchantRepoImpl) SaveMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {

	result := m.db.WithContext(ctx).Table(MERCHANT_TABLE).Save(setting)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *merchantRepoImpl) UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {

	result := m.db.WithContext(ctx).Table(MERCHANT_TABLE).Save(&setting)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *merchantRepoImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {

	var merchant models.MerchantSetting

	res := m.db.WithContext(ctx).Table(MERCHANT_TABLE).First(&merchant, "entity_id = ?", entityID)

	if res.Error != nil {
		return nil, res.Error
	}

	return nil, nil
}
