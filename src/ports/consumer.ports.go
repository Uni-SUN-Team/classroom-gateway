package ports

import "unisun/api/classroom-gateway/src/models"

type ConsumerService interface {
	GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (*models.ServiceIncomeResponse, error)
}
