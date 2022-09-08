package routes

import (
	"unisun/api/classroom-gateway/src/controllers"

	"github.com/gin-gonic/gin"
)

func HealCheck(g *gin.RouterGroup) {
	healController := controllers.NewControllerHealthCheckHandler()
	g.GET("/healcheck", healController.HealthCheckHandler)
}
