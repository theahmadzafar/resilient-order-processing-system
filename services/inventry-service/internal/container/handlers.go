package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/config"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/transport/http/handlers"
)

func BuildHandlers() []di.Def {
	return []di.Def{
		{
			Name: constants.MetaHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {

				return handlers.NewMetaHandler(), nil
			},
		}, {
			Name: constants.InventryHandlerName,
			Build: func(ctn di.Container) (interface{}, error) {
				inventrySvc := ctn.Get(constants.InventryServiceName).(*services.InventryService)
				conf := ctn.Get(constants.ConfigName).(*config.Config)

				return handlers.NewInventryHandler(inventrySvc, conf.Server.Timeout), nil
			},
		},
	}
}
