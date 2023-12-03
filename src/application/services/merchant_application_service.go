package services

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Azamjon99/restaurant-payment-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/services"
)

type MechantApplicationService interface {
	CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) (*pb.UpdateMerchantSettingResponse, error)
	GetMerchantSetting(ctx context.Context, req *pb.GetMerchantSettingRequest) (*pb.GetMerchantSettingResponse, error)
}

type merchantApplicationSvcImpl struct {
	merchantService services.MerchantSetting
}

func NewMechantApplicationService(merchantService services.MerchantSetting) MechantApplicationService {
	return &merchantApplicationSvcImpl{
		merchantService: merchantService,
	}
}

func (s *merchantApplicationSvcImpl) CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {
	return nil, nil
}

func (s *merchantApplicationSvcImpl) UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) (*pb.UpdateMerchantSettingResponse, error) {
	if req.GetCashboxId() == "" {
		return nil, fmt.Errorf("invalid or missing GetCashboxId: %s", req.GetCashboxId())
	}

	if req.GetEntityId() == "" {
		return nil, fmt.Errorf("invalid or missing GetEntityId: %s", req.GetEntityId())
	}

	varType := reflect.TypeOf(req.GetEnabled())

	if varType.Kind() == reflect.Bool {
		return nil, fmt.Errorf("invalid or missing GetEnabled: %t", req.GetEnabled())
	}

	err := s.merchantService.UpdateMerchantSetting(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMerchantSettingResponse{}, nil
}

func (s *merchantApplicationSvcImpl) GetMerchantSetting(ctx context.Context, req *pb.GetMerchantSettingRequest) (*pb.GetMerchantSettingResponse, error) {

	if req.GetEntityId() == "" {
		return nil, fmt.Errorf("invalid or missing GetEntityId: %s", req.GetEntityId())
	}

	settings, err := s.merchantService.GetMerchantSetting(ctx, req.GetEntityId())
	if err != nil {
		return nil, err
	}

	return &pb.GetMerchantSettingResponse{
		Setting: dtos.ToMerchantSettingsResponse(settings),
	}, nil
}
