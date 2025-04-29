package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/config"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/transport/http/handlers"
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
			Name: constants.PaymentHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {
				orderSvc := ctn.Get(constants.PaymentServiceName).(*services.PaymentService)
				conf := ctn.Get(constants.ConfigName).(*config.Config)

				return handlers.NewPaymentHandler(orderSvc, conf.Server.Timeout), nil
			},
		},
	}
}
