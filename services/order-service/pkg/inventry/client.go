package inventry

import (
	"context"
	"crypto/tls"
	"math"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	GetAvailableStocksByID(ctx context.Context, in *GetAvailableStocksByIDIn) (*GetAvailableStocksByIDOut, error)
	BuyStocksByID(ctx context.Context, in *BuyStocksByIDIn) (*BuyStocksByIDOut, error)
}

type Config struct {
	Host     string
	Port     string
	IsSecure bool
}

func NewClient(cfg *Config) (Client, error) {
	var err error

	service := &client{}
	service.api, err = newClient(cfg.Host, cfg.Port, cfg.IsSecure)

	if err != nil {
		return service, err
	}

	return service, nil
}

func newClient(host, port string, isSecure bool) (InventryClient, error) {
	addr := host + ":" + port

	var (
		conn *grpc.ClientConn
		err  error
	)

	if isSecure {
		config := &tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		}

		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(config)),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))
	} else {
		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))
	}

	if err != nil {
		zap.S().Errorf("can not dial %v: %v", addr, err)

		return nil, err
	}

	return NewInventryClient(conn), nil
}

type client struct {
	api InventryClient
}

func (c *client) GetAvailableStocksByID(ctx context.Context,
	in *GetAvailableStocksByIDIn) (*GetAvailableStocksByIDOut, error) {
	return c.api.GetAvailableStocksByID(ctx, in)
}
func (c *client) BuyStocksByID(ctx context.Context,
	in *BuyStocksByIDIn) (*BuyStocksByIDOut, error) {
	return c.api.BuyStocksByID(ctx, in)
}
