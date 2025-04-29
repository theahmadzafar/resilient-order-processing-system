package rpc

import (
	"io"

	"github.com/theahmadzafar/resilient-order-processing-system/services/proto/inventry"
)

type Handler struct {
	inventry.InventryServer
	cfg *Config
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
) *Handler {
	h := &Handler{
		cfg: cnf,
	}

	return h
}
