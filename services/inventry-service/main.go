package inventryservice

import (
	"context"
	"sync"

	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/container"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/logger"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/transport/http"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/transport/rpc"
	"github.com/theahmadzafar/resilient-order-processing-system/utils"
	"go.uber.org/zap"
)

func StartInventryService(ctx context.Context, wg *sync.WaitGroup) {

	app := container.Build(ctx, wg)

	_ = app.Get(constants.LoggerName).(*logger.Logger)

	server := app.Get(constants.ServerName).(*http.Server)

	zap.S().Info("Starting http server...")

	go server.Run()

	zap.S().Info("Starting grpc server...")

	rpcHandler := app.Get(constants.RPCName).(*rpc.Handler)

	go rpc.StartUnsecureRPCServer(rpcHandler)

	zap.S().Infof("Got %s signal. Shutting down...", <-utils.WaitTermSignal())

	if err := server.Shutdown(); err != nil {
		zap.S().Errorf("Error stopping server: %s", err)
	}
	zap.S().Info("Service stopped.")

}
