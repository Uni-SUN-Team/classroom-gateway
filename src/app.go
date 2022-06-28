package src

import (
	"unisun/api/classroom-gateway/src/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func App() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(viper.GetString("app.context_path") + viper.GetString("app.root_path") + "/v1")
	{
		route.Consumer(g)
	}
	return r
}
