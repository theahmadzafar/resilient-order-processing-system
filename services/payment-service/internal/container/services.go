package container

import (
	"github.com/sarulabs/di"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/constants"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/services"
	mockdatabase "github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/pkg/mock_database"
)

func BuildServices() []di.Def {
	return []di.Def{
		{
			Name: constants.PaymentServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				orderRepo := ctn.Get(constants.MockPaymentRepoName).(*mockdatabase.Repo)

				return services.NewPaymentService(orderRepo), nil
			},
		},
	}
}
