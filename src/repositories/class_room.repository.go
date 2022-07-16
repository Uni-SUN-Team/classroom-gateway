package repositories

import (
	"encoding/json"
	"log"
	"unisun/api/classroom-gateway/src/constants"
	"unisun/api/classroom-gateway/src/models"
	"unisun/api/classroom-gateway/src/models/original"
	"unisun/api/classroom-gateway/src/ports"

	"github.com/spf13/viper"
)

type Service struct {
	ConsumerService ports.ConsumerService
}

func New(consumerService ports.ConsumerService) *Service {
	return &Service{
		ConsumerService: consumerService,
	}
}

func (svr *Service) FindAll(query string) (*original.ResponseSuccess, error) {
	var request = models.ServiceIncomeRequest{}
	request.Method = constants.GET
	// request.Path = viper.GetString("endpoint.strapi.mapping.class_room.path") + "?populate[SEO][populate][shareImage][populate]=*&populate[thumbnail]=*&populate[advisors][populate]=*&populate[categories]=*&populate[Courses][populate][courses][populate][course][populate]=*" + query
	request.Path = viper.GetString("endpoint.strapi.mapping.class_room.path") + "?populate[SEO][populate][shareImage][populate]=*&populate[thumbnail]=*&populate[advisors]=*&populate[categories]=*&populate[Courses]=*" + query
	request.Body = nil
	var resultMapping = models.ServiceIncomeResponse{}
	if result, err := svr.ConsumerService.GetInformationFormStrapi(request); err != nil {
		log.Panic(err)
		return nil, err
	} else {
		resultMapping = *result
	}
	var classRoom = original.ResponseSuccess{}
	if err := json.Unmarshal([]byte(resultMapping.Payload), &classRoom); err != nil {
		log.Panic(err)
		return nil, err
	}
	return &classRoom, nil
}

func (svr *Service) FindById(id string) (*original.ResponseSingleSuccess, error) {
	var request = models.ServiceIncomeRequest{}
	request.Method = constants.GET
	request.Path = viper.GetString("endpoint.strapi.mapping.class_room.path") + "/" + id + "?populate[SEO][populate][shareImage][populate]=*&populate[thumbnail]=*&populate[advisors][populate]=*&populate[categories]=*&populate[Courses][populate][courses][populate][course][populate]=*"
	request.Body = nil
	var resultMapping = models.ServiceIncomeResponse{}
	if result, err := svr.ConsumerService.GetInformationFormStrapi(request); err != nil {
		log.Panic(err)
		return nil, err
	} else {
		resultMapping = *result
	}
	var classRoom = original.ResponseSingleSuccess{}
	if err := json.Unmarshal([]byte(resultMapping.Payload), &classRoom); err != nil {
		log.Panic(err)
		return nil, err
	}
	return &classRoom, nil
}

func (svr *Service) FindBySlug(slug string) (*original.ResponseSuccess, error) {
	var request = models.ServiceIncomeRequest{}
	request.Method = constants.GET
	request.Path = viper.GetString("endpoint.strapi.mapping.class_room.path") + "?populate[SEO][populate][shareImage][populate]=*&populate[thumbnail]=*&populate[advisors][populate]=*&populate[categories]=*&populate[Courses][populate][courses][populate][course][populate]=*"
	request.Body = nil
	var resultMapping = models.ServiceIncomeResponse{}
	if result, err := svr.ConsumerService.GetInformationFormStrapi(request); err != nil {
		log.Panic(err)
		return nil, err
	} else {
		resultMapping = *result
	}
	var classRoom = original.ResponseSuccess{}
	if err := json.Unmarshal([]byte(resultMapping.Payload), &classRoom); err != nil {
		log.Panic(err)
		return nil, err
	}
	return &classRoom, nil
}
