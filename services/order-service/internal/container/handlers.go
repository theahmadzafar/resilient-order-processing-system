package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/config"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/transport/http/handlers"
)

func BuildHandlers() []di.Def {
	return []di.Def{
		{
			Name: constants.MetaHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {

				return handlers.NewMetaHandler(), nil
			},
		},
		{
			Name: constants.OrderHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {
				orderSvc := ctn.Get(constants.OrderServiceName).(*services.OrderService)
				conf := ctn.Get(constants.ConfigName).(*config.Config)

				return handlers.NewOrderHandler(orderSvc, conf.Server.Timeout), nil
			},
		},
	}
}
