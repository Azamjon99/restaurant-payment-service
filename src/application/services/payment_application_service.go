package services

import (
	"context"
	"fmt"

	"github.com/Azamjon99/restaurant-payment-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/services"
)

type PaymentApplicationService interface {
	// CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error)
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
	SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) (*pb.SaveTransactionResponse, error)
	UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error)
	GetTransaction(ctx context.Context, txID string) (*models.Transaction, error)
}

type paymentApplicationSvcImpl struct {
	payService services.PaymentService
}

func NewPaymentApplicationService(payService services.PaymentService) PaymentApplicationService {
	return &paymentApplicationSvcImpl{
		payService: payService,
	}
}

// func (p *paymentApplicationSvcImpl) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {
// 	return nil, nil
// }

func (p *paymentApplicationSvcImpl) PayReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error) {

	if req.GetReceiptId() == "" {
		return nil, fmt.Errorf("invalid or missing ReceiptId: %s", req.GetReceiptId())
	}

	receipt, err := p.payService.PayReceipt(ctx, req.GetReceiptId())
	if err != nil {
		return nil, err
	}

	return &pb.GetReceiptResponse{
		Receipt: dtos.GetReceiptResponse(receipt),
	}, nil
}

func (p *paymentApplicationSvcImpl) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {
	return nil, nil
}

func (p *paymentApplicationSvcImpl) SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) (*pb.SaveTransactionResponse, error) {

	err := p.payService.SaveTransaction(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.SaveTransactionResponse{}, nil
}

func (p *paymentApplicationSvcImpl) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {

	if req.Transaction.GetId() == "" {
		return nil, fmt.Errorf("invalid or missing Id: %s", req.Transaction.GetId())
	}

	if req.Transaction.GetCurrency() == "" {
		return nil, fmt.Errorf("invalid or missing Id: %s", req.Transaction.GetCurrency())
	}

	tx, err := p.payService.UpdateTransaction(ctx, req)

	if err != nil {
		return tx, err
	}
	return &pb.UpdateTransactionResponse{}, nil
}

func (p *paymentApplicationSvcImpl) GetTransaction(ctx context.Context, txID string) (*models.Transaction, error) {
	return nil, nil
}
