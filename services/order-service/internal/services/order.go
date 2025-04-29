package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/entities"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/order-service/pkg/mock_database"
	"github.com/theahmadzafar/resilient-order-processing-system/services/proto/inventry"
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

func (s OrderService) GetOrder(ctx context.Context, orderId uuid.UUID) (*entities.GetOrderResponse, error) {
	res := &entities.GetOrderResponse{}

	for _, item := range s.orderRepo.List {
		if orderId == item.ID {
			res = &entities.GetOrderResponse{
				ID:        item.ID,
				Status:    item.Status,
				PaymentID: item.PaymentID,
				Items:     make([]entities.Item, 0),
			}

			for _, orderItem := range item.Items {
				res.Items = append(res.Items, entities.Item{
					ID:    orderItem.ID,
					Name:  orderItem.Name,
					Count: orderItem.Count,
				})
			}
		}
	}

	if res.ID == uuid.Nil {
		return nil, errors.New("no record found")
	}

	return res, nil
}
