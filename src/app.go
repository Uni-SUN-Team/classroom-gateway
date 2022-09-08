package src

import (
	"strings"
	"unisun/api/classroom-gateway/src/config/environment"
	"unisun/api/classroom-gateway/src/routes"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	appEnv := environment.ENV.App
	ginEnv := environment.ENV.Gin
	r := gin.Default()
	r.SetTrustedProxies(strings.Split(ginEnv.Configs.TrustedProxies, ","))
	g := r.Group(strings.Join([]string{appEnv.ContextPath, ginEnv.RootPath, ginEnv.Version}, "/"))
	{
		routes.HealCheck(g)
		routes.ClassRoomRoute(g)
	}
	return r
}
