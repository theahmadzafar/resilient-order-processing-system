package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/pkg/inventry"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/order-service/pkg/mock_database"
)

func BuildServices() []di.Def {
	return []di.Def{
		{
			Name: constants.OrderServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				orderRepo := ctn.Get(constants.MockOrderRepoName).(*mockdatabase.OrderRepo)
				inventryMicroSvc := ctn.Get(constants.InventryMicroSvcName).(inventry.Client)

				return services.NewOrderService(orderRepo, inventryMicroSvc), nil
			},
		},
	}
}
