package container

import (
	"context"
	"fmt"
	"sync"

	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/config"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/logger"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/transport/http"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/pkg/mock_database"
)

var container di.Container
var once sync.Once

func Build(ctx context.Context, wg *sync.WaitGroup) di.Container {
	once.Do(func() {
		builder, _ := di.NewBuilder()
		defs := []di.Def{
			{
				Name: constants.ConfigName,
				Build: func(ctn di.Container) (interface{}, error) {
					return config.New()
				},
			},
			{
				Name: constants.LoggerName,
				Build: func(ctn di.Container) (interface{}, error) {
					cfg := ctn.Get(constants.ConfigName).(*config.Config)

					zapLogger, err := logger.NewLogger(cfg.Logger)
					if err != nil {
						return nil, fmt.Errorf("can't initialize zap logger: %v", err)
					}

					return zapLogger, nil
				},
			},
			{
				Name: constants.MockPaymentRepoName,
				Build: func(ctn di.Container) (interface{}, error) {

					return mockdatabase.NewMockConnection()
				},
			},
			{
				Name: constants.ServerName,
				Build: func(ctn di.Container) (interface{}, error) {
					cfg := ctn.Get(constants.ConfigName).(*config.Config)

					var publicHandlers = []http.Handler{
						ctn.Get(constants.MetaHandlerName).(http.Handler),
						ctn.Get(constants.PaymentHandlerName).(http.Handler),
					}

					return http.New(ctx, wg, cfg.Server, publicHandlers), nil
				},
			},
		}

		defs = append(defs, BuildServices()...)
		defs = append(defs, BuildHandlers()...)

		if err := builder.Add(defs...); err != nil {
			panic(err)
		}

		container = builder.Build()
	})

	return container
}
