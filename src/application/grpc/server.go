package grpc

import (
	"context"

	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	card "github.com/Azamjon99/restaurant-payment-service/src/application/services"
)

type Server struct {
	pb.RestaurantPaymentServer
	cardApp     card.CardApplicationService
	payApp      card.PaymentApplicationService
	merchantApp card.MechantApplicationService
}

func NewServer(
	cardApp card.CardApplicationService,
	payApp card.PaymentApplicationService,
	merchantApp card.MechantApplicationService,
) *Server {
	return &Server{
		cardApp:     cardApp,
		payApp:      payApp,
		merchantApp: merchantApp,
	}
}

func (s *Server) PayReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error) {
	return s.payApp.PayReceipt(ctx, req)
}

func (s *Server) MerchantSettings(ctx context.Context, req *pb.GetMerchantSettingRequest) (*pb.GetMerchantSettingResponse, error) {
	return s.merchantApp.GetMerchantSetting(ctx, req)
}

func (s *Server) UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) (*pb.UpdateMerchantSettingResponse, error) {
	return s.merchantApp.UpdateMerchantSetting(ctx, req)
}

func (s *Server) SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) (*pb.SaveTransactionResponse, error) {
	return s.payApp.SaveTransaction(ctx, req)
}

func (s *Server) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	return s.payApp.UpdateTransaction(ctx, req)
}
