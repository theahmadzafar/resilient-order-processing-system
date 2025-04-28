package rpc

import (
	"io"

	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/pkg/api"
)

type Handler struct {
	api.InventryServer
	inventrySvc *services.InventryService
	cfg         *Config
}

func (h *Handler) HealthCheck(stream api.Inventry_HealthCheckServer) error {
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
	inventrySvc *services.InventryService,
) *Handler {
	h := &Handler{
		cfg:         cnf,
		inventrySvc: inventrySvc,
	}

	return h
}
