package services

import (
	"context"

	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/entities"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/pkg/mock_database"
)

type PaymentService struct {
	orderRepo *mockdatabase.Repo
}

func NewPaymentService(orderRepo *mockdatabase.Repo) *PaymentService {
	return &PaymentService{
		orderRepo: orderRepo,
	}
}
func (s PaymentService) Pay(ctx context.Context, req entities.PaymentRequest) error {
	s.orderRepo.List = append(s.orderRepo.List, mockdatabase.Invoice{
		ID:      req.ID,
		OrderID: req.OrderID,
	})

	// todo event is added to the distributed mechanism so the order service can dispatch the order

	return nil
}
