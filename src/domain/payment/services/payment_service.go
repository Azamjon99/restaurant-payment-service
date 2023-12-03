package services

import (
	"context"

	pb "github.com/Azamjon99/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/models"
	"github.com/Azamjon99/restaurant-payment-service/src/domain/payment/repositories"
)

type PaymentService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayReceipt(ctx context.Context, receiptID string) (*models.Receipt, error)
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
	SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) error
	UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error)
	GetTransaction(ctx context.Context, txID string) (*models.Transaction, error)
}

type paymentSvcImpl struct {
	receiptRepo repositories.ReceiptRepository
	txRepo      repositories.TransactionRepository
}

func NewPaymentCardService(receiptRepo repositories.ReceiptRepository, txRepo repositories.TransactionRepository) PaymentService {
	return &paymentSvcImpl{
		receiptRepo: receiptRepo,
		txRepo:      txRepo,
	}
}

func (p *paymentSvcImpl) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {

	// err := p.receiptRepo.SaveReceipt(ctx, receipt)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
func (p *paymentSvcImpl) PayReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {

	receipt, err := p.receiptRepo.GetReceipt(ctx, receiptID)

	if err != nil {
		return nil, err
	}
	return receipt, nil
}
func (p *paymentSvcImpl) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {

	// err := p.receiptRepo.GetReceiptByOrder(ctx, orderID)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
func (p *paymentSvcImpl) SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) error {

	// tx := models.Transaction{
	// 	ID: req.Ge,
	// }

	// err := p.txRepo.SaveTx(ctx, req)

	// if err != nil {
	// 	return err
	// }

	return nil
}
func (p *paymentSvcImpl) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {

	tx := models.Transaction{
		ID:       req.Transaction.GetId(),
		Currency: req.Transaction.GetCurrency(),
	}

	err := p.txRepo.UpdateTx(ctx, &tx)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateTransactionResponse{}, nil
}
func (p *paymentSvcImpl) GetTransaction(ctx context.Context, txID string) (*models.Transaction, error) {

	// tx, err := p.txRepo.GetTx(ctx, txID)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
