package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/entities"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/pkg/inventry"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/order-service/pkg/mock_database"
)

type OrderService struct {
	orderRepo        *mockdatabase.OrderRepo
	inventryMicroSvc inventry.Client
}

func NewOrderService(orderRepo *mockdatabase.OrderRepo, inventryMicroSvc inventry.Client) *OrderService {
	return &OrderService{
		orderRepo: orderRepo, inventryMicroSvc: inventryMicroSvc,
	}
}
func (s OrderService) PlaceOrder(ctx context.Context, req entities.OrderRequest) error {
	list := make([]mockdatabase.Item, 0)

	for _, item := range req.List {
		out, err := s.inventryMicroSvc.GetAvailableStocksByID(ctx, &inventry.GetAvailableStocksByIDIn{Id: item.ID.String()})
		if err != nil {
			return err
		}

		if out.Item.Count < item.Count {
			return fmt.Errorf("%s only have %v stocks available ", out.Item.Name, out.Item.Count)
		}

		list = append(list, mockdatabase.Item{ID: item.ID, Name: out.Item.Name, Count: item.Count})
	}

	s.orderRepo.List = append(s.orderRepo.List, mockdatabase.Order{ID: uuid.New(), Status: "pending", Items: list})

	return nil
}
