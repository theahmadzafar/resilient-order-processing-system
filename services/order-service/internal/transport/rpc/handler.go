package rpc

import (
	"io"

	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/pkg/api"
)

type Handler struct {
	api.OrderServer
	orderSvc *services.OrderService
	cfg      *Config
}

func (h *Handler) HealthCheck(stream api.Order_HealthCheckServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		if err = stream.Send(msg); err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}
	}
}

func NewHandler(
	cnf *Config,
	orderSvc *services.OrderService,
) *Handler {
	h := &Handler{
		cfg: cnf, orderSvc: orderSvc,
	}

	return h
}
