package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	inventryservice "github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service"
	orderservice "github.com/theahmadzafar/resilient-order-processing-system/services/order-service"
	"github.com/theahmadzafar/resilient-order-processing-system/utils"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("resilient-order-processing-system")

	now := time.Now()
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	go inventryservice.StartInventryService(ctx, wg)
	go orderservice.StartOrderService(ctx, wg)
	// go paymentservice.StartPaymentService(ctx, wg)

	zap.S().Infof("Up and running (%s)", time.Since(now))
	zap.S().Infof("Got %s signal. Shutting down...", <-utils.WaitTermSignal())

	wg.Wait()
	zap.S().Info("Service stopped.")
}
