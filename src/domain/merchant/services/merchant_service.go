package services

import (
	"context"

	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/repositories"
)

type MerchantSetting interface {
	CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}

type merchantSvcImpl struct {
	merchantRepo repositories.MerchantRepository
}

func NewMerchantService(merchantRepo repositories.MerchantRepository) MerchantSetting {
	return &merchantSvcImpl{
		merchantRepo: merchantRepo,
	}
}

func (m *merchantSvcImpl) CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {

	// err := m.merchantRepo.SaveMerchantSetting(ctx, settings)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (m *merchantSvcImpl) UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) error {

	// err := m.merchantRepo.UpdateMerchantSetting(ctx, settings)

	// if err != nil {
	// 	return err
	// }

	return nil
}

func (m *merchantSvcImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {

	// merchant, err := m.merchantRepo.GetMerchantSetting(ctx, entityID)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
