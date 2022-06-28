package controller

import (
	"log"
	"net/http"
	"strings"
	"unisun/api/classroom-gateway/src/model/original"
	"unisun/api/classroom-gateway/src/repository"
	"unisun/api/classroom-gateway/src/service"
	"unisun/api/classroom-gateway/src/utils"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	new_query := ""
	if query := c.Request.URL.RawQuery; query != "" {
		new_query += "&" + strings.Trim(query, "?")
	}
	httpRequestResult := utils.New()
	consumerResult := service.New(httpRequestResult)
	var classroom = original.ResponseSuccess{}
	if data, err := repository.New(consumerResult).FindAll(new_query); err != nil {
		log.Panic(err)
	} else {
		classroom = *data
	}
	// err := json.Unmarshal([]byte(data.Payload), &response)
	// if err != nil {
	// 	log.Panic("Change byte to json article", err.Error())
	// 	responseError.Error.Status = http.StatusBadRequest
	// 	responseError.Error.Message = err.Error()
	// 	responseError.Error.Name = "Json unmarshal"
	// 	c.JSON(http.StatusBadRequest, responseError)
	// 	return
	// } else {
	// 	err = nil
	// }
	c.JSON(http.StatusOK, classroom)
}

// func ClassRoomBySlugPreview(c *gin.Context) {
// 	slug := c.Param("slug")
// 	var request = model.ServiceIncomeRequest{}
// 	var response = model.ResponseSuccessFull{}
// 	var responseError = model.ResponseFail{}
// 	request.Method = constants.GET
// 	request.Path = os.Getenv(constants.PATH_STRAPI_CLASS_ROOM) + "?populate[SEO][populate][shareImage][populate]=*&populate[thumbnail]=*&populate[advisors]=*&populate[categories]=*&populate[Courses]=*&filters[slug][$eq]=" + slug
// 	request.Body = nil
// 	data := service.GetInformationFormStrapi(request)
// 	err := json.Unmarshal([]byte(data.Payload), &response)
// 	if err != nil {
// 		log.Panic("Change byte to json article", err.Error())
// 		responseError.Error.Status = http.StatusBadRequest
// 		responseError.Error.Message = err.Error()
// 		responseError.Error.Name = "Json unmarshal"
// 		c.JSON(http.StatusBadRequest, responseError)
// 		return
// 	} else {
// 		err = nil
// 	}
// 	c.JSON(http.StatusOK, response)
// }
