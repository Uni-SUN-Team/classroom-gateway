package route

import (
	"unisun/api/classroom-gateway/src/controller"

	"github.com/gin-gonic/gin"
)

func Consumer(r *gin.RouterGroup) {
	r.GET("class-rooms", controller.GetAll)
	r.GET("class-rooms/slug/:slug", controller.GetAll)
}
