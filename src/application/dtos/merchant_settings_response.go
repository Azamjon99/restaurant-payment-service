package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/merchant/models"
)

func ToMerchantSettingsResponse(setting *models.MerchantSetting) *pb.MerchantSetting {
	return &pb.MerchantSetting{
		EntityId:  setting.EntityID,
		CashboxId: setting.CashboxID,
		Enabled:   setting.Enabled,
		CreatedAt: setting.UpdatedAt.Format(time.RFC3339),
		UpdatedAt: setting.UpdatedAt.Format(time.RFC3339),
	}
}
