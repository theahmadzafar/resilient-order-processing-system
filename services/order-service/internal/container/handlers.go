package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/constants"
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
	}
}
