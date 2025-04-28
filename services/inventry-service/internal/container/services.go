package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/services"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/pkg/mock_database"
)

func BuildServices() []di.Def {
	return []di.Def{{
		Name: constants.InventryServiceName,
		Build: func(ctn di.Container) (interface{}, error) {
			inventry := ctn.Get(constants.MockInventryName).(*mockdatabase.Inventry)

			return services.NewInventryService(inventry), nil
		},
	}}
}
