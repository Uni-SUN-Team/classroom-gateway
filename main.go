package main

import (
	"log"
	"unisun/api/classroom-gateway/src"
	"unisun/api/classroom-gateway/src/config"

	"github.com/spf13/viper"
)

func main() {
	envService := config.New("application", "./resource")
	if err := envService.ConfigENV(); err != nil {
		log.Panic(err)
	}
	r := src.App()
	port := viper.GetString("app.port")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
