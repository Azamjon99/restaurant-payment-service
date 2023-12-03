package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/models"
)

func GetReceiptResponse(receipt *models.Receipt) *pb.Receipt {
	return &pb.Receipt{
		Id:        receipt.ID,
		OrderId:   receipt.OrderID,
		CardId:    receipt.CardID,
		Amount:    int32(receipt.Amount),
		CreatedAt: receipt.CreatedAt.Format(time.RFC3339),
	}
}
