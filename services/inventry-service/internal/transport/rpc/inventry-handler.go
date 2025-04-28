package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) GetAvailableStocksByID(ctx context.Context, in *api.GetAvailableStocksByIDIn) (
	*api.GetAvailableStocksByIDOut, error) {
	ID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	item, err := h.inventrySvc.GetByID(ID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &api.GetAvailableStocksByIDOut{Item: &api.Item{Id: item.ID.String(), Name: item.Name, Count: item.Count}}, nil
}
func (h Handler) BuyStocksByID(ctx context.Context, in *api.BuyStocksByIDIn) (*api.BuyStocksByIDOut, error) {
	ID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	err = h.inventrySvc.BuyByID(ID, in.Count)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &api.BuyStocksByIDOut{Status: "success"}, nil
}
