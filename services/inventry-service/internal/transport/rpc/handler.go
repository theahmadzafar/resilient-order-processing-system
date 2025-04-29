package rpc

import (
	"io"

	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/proto/inventry"
)

type Handler struct {
	inventry.InventryServer
	inventrySvc *services.InventryService
	cfg         *Config
}

func (h *Handler) HealthCheck(stream inventry.Inventry_HealthCheckServer) error {
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
