package ports

import "unisun/api/classroom-gateway/src/model"

type ConsumerService interface {
	GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (*model.ServiceIncomeResponse, error)
}
