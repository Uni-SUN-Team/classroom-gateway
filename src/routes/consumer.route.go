package routes

import (
	"unisun/api/classroom-gateway/src/controllers"

	"github.com/gin-gonic/gin"
)

func Consumer(r *gin.RouterGroup) {
	r.GET("class-rooms", controllers.GetAll)
	r.GET("class-rooms/slug/:slug", controllers.GetAll)
}
