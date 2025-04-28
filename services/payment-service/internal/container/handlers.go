package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/constants"
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
	}
}
