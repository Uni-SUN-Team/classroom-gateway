package services

import (
	"encoding/json"
	"io/ioutil"
	"unisun/api/classroom-gateway/src/constants"
	"unisun/api/classroom-gateway/src/models"

	"unisun/api/classroom-gateway/src/ports"

	"github.com/spf13/viper"
)

type UtilsService struct {
	UtilsService ports.HTTPService
}

func New(utilsService ports.HTTPService) *UtilsService {
	return &UtilsService{
		UtilsService: utilsService,
	}
}

func (svr *UtilsService) GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (*models.ServiceIncomeResponse, error) {
	var serviceIncomeResponse = models.ServiceIncomeResponse{}
	url := viper.GetString("endpoint.strapi.host") + viper.GetString("endpoint.strapi.path")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	} else {
		err = nil
	}
	response, err := svr.UtilsService.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	} else {
		err = nil
		defer response.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	} else {
		err = nil
	}
	return &serviceIncomeResponse, nil
}
